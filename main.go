package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imaadssebintu/portfolio/internal/data"
)

//go:embed templates
var templatesFS embed.FS

//go:embed static
var staticFS embed.FS

var pages = map[string]string{
	"home":    "templates/home.html",
	"project": "templates/project.html",
	"404":     "templates/404.html",
}

var pageTemplates map[string]*templateSet

func buildTemplates() map[string]*templateSet {
	sets := make(map[string]*templateSet)
	for name, file := range pages {
		sets[name] = newTemplateSet("templates/layout.html", file)
	}
	return sets
}

type contactEntry struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	At      string `json:"at"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	contactsFile := os.Getenv("CONTACTS_FILE")
	if contactsFile == "" {
		contactsFile = "data/contacts.jsonl"
	}

	pageTemplates = buildTemplates()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))
	r.Use(securityHeaders)

	staticHandler := staticFileServer()
	r.Handle("/static/*", http.StripPrefix("/static/", staticHandler))

	r.Get("/", handleHome(contactsFile))
	r.Get("/projects/{slug}", handleProject)
	r.Post("/contact", handleContact(contactsFile))
	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})
	r.NotFound(handle404)

	log.Printf("portfolio listening on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func staticFileServer() http.Handler {
	sub, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(sub))
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		next.ServeHTTP(w, r)
	})
}

func render(w http.ResponseWriter, set *templateSet, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := set.Execute(w, data); err != nil {
		log.Printf("render error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func handleHome(string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			handle404(w, r)
			return
		}
		render(w, pageTemplates["home"], map[string]any{
			"Title":      "",
			"MetaDesc":   "",
			"Year":       time.Now().Year(),
			"Projects":   data.Projects(),
			"Services":   data.Services,
			"Skills":     data.Skills,
			"Experience": data.Experience,
			"Education":  data.Education,
		})
	}
}

func handleProject(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	project, ok := data.ProjectBySlug(slug)
	if !ok {
		handle404(w, r)
		return
	}

	projects := data.Projects()
	var prev, next *data.Project
	for i := range projects {
		if projects[i].Slug == slug {
			if i > 0 {
				prev = &projects[i-1]
			}
			if i < len(projects)-1 {
				next = &projects[i+1]
			}
			break
		}
	}

	paragraphs := splitParagraphs(project.LongDesc)

	render(w, pageTemplates["project"], map[string]any{
		"Title":      project.Name,
		"MetaDesc":   project.Description,
		"Year":       time.Now().Year(),
		"Project":    project,
		"Paragraphs": paragraphs,
		"Prev":       prev,
		"Next":       next,
	})
}

func handleContact(contactsFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form", http.StatusBadRequest)
			return
		}
		entry := contactEntry{
			Name:    strings.TrimSpace(r.FormValue("name")),
			Email:   strings.TrimSpace(r.FormValue("email")),
			Message: strings.TrimSpace(r.FormValue("message")),
			At:      time.Now().UTC().Format(time.RFC3339),
		}
		if entry.Name == "" || entry.Email == "" || entry.Message == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}
		if len(entry.Message) > 5000 {
			http.Error(w, "Message too long", http.StatusBadRequest)
			return
		}
		if err := appendContact(contactsFile, entry); err != nil {
			log.Printf("save contact: %v", err)
			http.Error(w, "Could not save message", http.StatusInternalServerError)
			return
		}

		if strings.Contains(r.Header.Get("Accept"), "application/json") || r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
			return
		}
		http.Redirect(w, r, "/#contact", http.StatusSeeOther)
	}
}

func appendContact(path string, entry contactEntry) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(entry)
}

func handle404(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	render(w, pageTemplates["404"], map[string]any{
		"Title":      "Not Found",
		"Year":       time.Now().Year(),
		"Projects":   data.Projects(),
		"Services":   data.Services,
		"Skills":     data.Skills,
		"Experience": data.Experience,
		"Education":  data.Education,
	})
}

func splitParagraphs(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	parts := strings.Split(s, "\n\n")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
