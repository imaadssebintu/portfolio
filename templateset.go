package main

import (
	"html/template"
	"io"
	"io/fs"
)

type templateSet struct {
	tmpl *template.Template
}

func newTemplateSet(files ...string) *templateSet {
	sub, err := fs.Sub(templatesFS, "templates")
	if err != nil {
		panic(err)
	}
	names := make([]string, len(files))
	for i, f := range files {
		names[i] = trimPrefix(f, "templates/")
	}
	t := template.Must(template.New("").ParseFS(sub, names...))
	return &templateSet{tmpl: t}
}

func (s *templateSet) Execute(w io.Writer, data any) error {
	return s.tmpl.ExecuteTemplate(w, "layout", data)
}

func trimPrefix(s, prefix string) string {
	if len(s) >= len(prefix) && s[:len(prefix)] == prefix {
		return s[len(prefix):]
	}
	return s
}
