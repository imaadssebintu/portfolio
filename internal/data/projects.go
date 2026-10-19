package data

import "html/template"

type Project struct {
	Slug        string
	Name        string
	Tagline     string
	Description string
	LongDesc    string
	URL         string
	URLLabel    string
	Tech        []string
	Features    []string
	Status      string
}

func Projects() []Project {
	return []Project{
		{
			Slug:        "flexduuka",
			Name:        "Flexduuka",
			Tagline:     "Business inventory & operations platform",
			Description: "An inventory management platform built to help businesses manage products, sales, suppliers, expenses, POS operations, payments, production, and analytics in one place.",
			LongDesc: `Flexduuka is a business management platform designed to help businesses run day-to-day operations from a single system. It brings inventory, point of sale, suppliers, expenses, payments, production (bill of materials), analytics, and multi-branch management together so owners get a clear picture of their business without juggling spreadsheets or disconnected tools.

Built with a Go backend, PostgreSQL, Redis, and Docker, Flexduuka is designed for real deployment on VPS and cloud infrastructure, with mobile-friendly access for teams on the go.`,
			URL:      "https://flexduuka.com/",
			URLLabel: "flexduuka.com",
			Tech:     []string{"Go", "Go Fiber", "PostgreSQL", "Redis", "Docker", "Cloudflare"},
			Features: []string{
				"Dashboard & analytics",
				"Products & inventory management",
				"Point of Sale (POS)",
				"Suppliers & expenses",
				"Payments & mobile money",
				"Production / Bill of Materials",
				"Multi-user & multi-branch",
				"Reports & sales insights",
			},
			Status: "Live",
		},
		{
			Slug:        "voyawalk",
			Name:        "VoyaWalk",
			Tagline:     "Travel safer with verified companions",
			Description: "A travel safety application that connects travelers with verified companions, helping people explore more safely while walking in unfamiliar places.",
			LongDesc: `VoyaWalk is designed to connect travelers with verified companions or guides for safer walking journeys. The product spans mobile applications for iOS and Android backed by Go services, PostgreSQL, and cloud deployment.

Core building blocks include user profiles, verification flows, maps and location features, and backend APIs that keep companion matching reliable.`,
			URL:      "https://voyawalk.com/",
			URLLabel: "voyawalk.com",
			Tech:     []string{"React Native", "Expo", "iOS", "Android", "Go", "PostgreSQL", "Docker", "OpenStreetMap"},
			Features: []string{
				"Verified companion profiles",
				"iOS & Android apps",
				"Maps & location features",
				"Travel safety focus",
				"Backend APIs & cloud deployment",
			},
			Status: "Live",
		},
		{
			Slug:        "city-rides",
			Name:        "City Rides",
			Tagline:     "Vehicle marketplace for Uganda",
			Description: "A vehicle marketplace and car-request application built for the Ugandan market, letting users browse cars, filter by type and specs, and request vehicles.",
			LongDesc: `City Rides is a vehicle marketplace application focused on Uganda. Users can browse car listings, search and filter by make and body type, review new and used vehicle information, and make car requests with direct contact options via WhatsApp, phone, or email.

The app is published on the Apple App Store under developer imaad ssebintu, with cityridesug.com as the associated website.`,
			URL:      "https://cityridesug.com",
			URLLabel: "cityridesug.com",
			Tech:     []string{"iOS", "Swift", "App Store", "REST APIs"},
			Features: []string{
				"Car listings & search",
				"Make and body-type filters",
				"New & used vehicle info",
				"Car request flow",
				"WhatsApp / phone / email contact",
			},
			Status: "Live on the App Store",
		},
		{
			Slug:        "flexschool",
			Name:        "Flexschool",
			Tagline:     "School management system",
			Description: "A school management platform that brings administration, student records, attendance, results, grading, finance, and permissions into one system.",
			LongDesc: `Flexschool is a school management system designed for real school operations. It supports multiple roles — admin, head teacher, class teacher, subject teacher, and bursar — each with appropriate permissions.

Modules cover student management, attendance, academic results and grading, finance, administration, and email notifications. Development goals include standalone deployments with embedded databases and templates, and Windows executable builds for schools that prefer simple installs.`,
			Tech: []string{"Go", "Go Fiber", "PostgreSQL", "SQLite", "Docker", "HTML"},
			Features: []string{
				"Student management",
				"Attendance tracking",
				"Results & grading",
				"Finance & bursar tools",
				"Role-based permissions",
				"Email notifications",
			},
			Status: "In development",
		},
		{
			Slug:        "motor-hub",
			Name:        "Motor Hub",
			Tagline:     "Vehicle finance & management",
			Description: "A vehicle management and finance-tracking system that replaces spreadsheet-based tracking with a mobile-friendly application.",
			LongDesc: `Motor Hub makes it easier to record and manage vehicle-related financial information without relying on spreadsheets. The data model deliberately separates car IDs from number plates (for example CAR001 — Subaru) so records stay stable even when vehicle details change.

It combines vehicle records, financial tracking, an API-backed backend, and a mobile-friendly interface.`,
			Tech: []string{"SQLite", "API backend", "Mobile-friendly UI"},
			Features: []string{
				"Vehicle records",
				"Financial tracking",
				"Stable car IDs separate from plates",
				"Mobile-friendly interface",
			},
			Status: "In development",
		},
		{
			Slug:        "techflex",
			Name:        "Techflex",
			Tagline:     "Transforming Ideas into Digital Reality",
			Description: "A technology business focused on building websites, mobile applications, business systems, and custom digital products for real-world problems.",
			LongDesc: `Techflex is Imaad's technology brand — a business focused on turning ideas into practical digital products and software solutions. Work spans website development, web applications, mobile applications (iOS and Android), business software, and custom digital solutions.

Techflex is the umbrella behind products like Flexduuka, VoyaWalk, City Rides, Flexschool, and Motor Hub.`,
			Tech: []string{"Web", "Mobile", "Business systems", "Custom software"},
			Features: []string{
				"Website development",
				"Web & mobile applications",
				"Business software",
				"Custom digital solutions",
			},
			Status: "Active brand",
		},
	}
}

func ProjectBySlug(slug string) (Project, bool) {
	for _, p := range Projects() {
		if p.Slug == slug {
			return p, true
		}
	}
	return Project{}, false
}

var Skills = map[string][]string{
	"Backend":      {"Go", "Go Fiber", "REST APIs", "Backend development"},
	"Frontend":     {"HTML", "Responsive interfaces", "Business dashboards", "Web applications"},
	"Mobile":       {"React Native", "Expo", "iOS", "Android"},
	"Databases":    {"PostgreSQL", "SQLite", "Redis"},
	"DevOps":       {"Docker", "Docker Compose", "Linux", "VPS", "Cloudflare", "Cloud tunnels"},
	"Integrations": {"Mobile money", "Payment APIs", "WhatsApp API", "Webhooks"},
}

type Service struct {
	Title string
	Desc  string
	Icon  template.HTML
}

var Services = []Service{
	{
		Title: "Web Applications",
		Desc:  "Custom web apps and business dashboards built for real workflows.",
		Icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><rect x="2.5" y="3.5" width="19" height="14" rx="2"/><path d="M8 20.5h8"/><path d="M12 17.5v3"/><path d="M6.5 7.5h5"/><path d="M6.5 10.5h8"/><path d="M6.5 13.5h3"/></svg>`,
	},
	{
		Title: "Mobile Apps",
		Desc:  "iOS and Android applications with React Native and Expo.",
		Icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><rect x="7" y="2.5" width="10" height="19" rx="2.5"/><path d="M11 18.5h2"/><path d="M10 5h4"/></svg>`,
	},
	{
		Title: "Business Systems",
		Desc:  "Inventory, POS, school, and vehicle management platforms.",
		Icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M3.5 9.5 12 4l8.5 5.5"/><path d="M5 11v8.5h14V11"/><path d="M9.5 19.5v-5h5v5"/><path d="M3.5 9.5h17"/></svg>`,
	},
	{
		Title: "API Development",
		Desc:  "REST APIs, integrations, webhooks, and payment collection flows.",
		Icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M8.5 8.5 5 12l3.5 3.5"/><path d="M15.5 8.5 19 12l-3.5 3.5"/><path d="M13.5 5.5 10.5 18.5"/></svg>`,
	},
	{
		Title: "Payment Integration",
		Desc:  "Mobile money and payment workflows for African markets.",
		Icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><rect x="2.5" y="5.5" width="19" height="13" rx="2.5"/><path d="M2.5 10h19"/><path d="M6.5 14.5h3.5"/><path d="M14 14.5h3.5"/></svg>`,
	},
	{
		Title: "Cloud & DevOps",
		Desc:  "Docker deployments, VPS setup, Cloudflare, and server support.",
		Icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M7.5 17.5h10a4 4 0 0 0 .4-8 5.5 5.5 0 0 0-10.5-1.2A3.8 3.8 0 0 0 7.5 17.5z"/><path d="M12 12.5v5"/><path d="M9.5 15 12 17.5 14.5 15"/></svg>`,
	},
}

var Experience = []struct{ Role, Org, Period, Detail string }{
	{"Self-Employed IT Specialist", "Techflex", "2021 – Present", "Website and software development, business technology solutions, PC and software support, technical troubleshooting."},
	{"Digital & Service Support", "Lugave Auto / URA Services", "2022 – Present", "TIN services, driver's licence services, vehicle transfer and URA-related services, digital support."},
	{"ICT Intern", "Primacy ICT Solution / KIU", "Internship", "Approximately two months of hands-on ICT experience."},
}

var Education = []struct{ Title, Org, Year string }{
	{"B.Sc. Information Technology", "Kampala International University", "2024"},
	{"Huawei Big Data Certification", "Huawei", "November 2022"},
	{"UACE", "Kyaddondo Secondary School", "2019"},
	{"UCE", "Kyaddondo Secondary School", "2017"},
}
