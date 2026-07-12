package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"ascentia-web/ai"
	"ascentia-web/assets"
)

type Server struct {
	DB    *sql.DB
	AI    ai.Provider
	Templ map[string]*template.Template
}

func NewServer(db *sql.DB, aiProvider ai.Provider) (*Server, error) {
	templates := make(map[string]*template.Template)

	// Načítame layout z embedded FS
	layoutBytes, err := assets.GetTemplateFS().ReadFile("layout.html")
	if err != nil {
		return nil, err
	}
	layoutStr := string(layoutBytes)

	// Parzneme každú stránku samostatne s layoutom
	pages := []string{"dashboard", "services", "process", "privacy", "kompas", "voice", "faq", "consultation", "blog", "blog_go_pre_enterprise", "blog_ai_pravnikom", "blog_voice_crm_case", "pricing"}
	for _, page := range pages {
		pageBytes, err := assets.GetTemplateFS().ReadFile("templates/" + page + ".html")
		if err != nil {
			return nil, err
		}
		tmpl, err := template.New("layout").Parse(layoutStr)
		if err != nil {
			return nil, err
		}
		_, err = tmpl.Parse(string(pageBytes))
		if err != nil {
			return nil, err
		}
		templates[page] = tmpl
	}

	return &Server{
		DB:    db,
		AI:    aiProvider,
		Templ: templates,
	}, nil
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Pre neexistujúce cesty skúsime obslúžiť statické súbory
		http.NotFound(w, r)
		return
	}
	s.renderTemplate(w, r, "dashboard", nil)
}

func (s *Server) HandleServices(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "services", nil)
}

func (s *Server) HandleProcess(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "process", nil)
}

func (s *Server) HandlePrivacy(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "privacy", nil)
}

func (s *Server) HandleKompas(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "kompas", nil)
}

func (s *Server) HandleVoice(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "voice", nil)
}

func (s *Server) HandleFAQ(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "faq", nil)
}

func (s *Server) HandleConsultation(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "consultation", nil)
}

func (s *Server) HandleBlog(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "blog", nil)
}

func (s *Server) HandlePricing(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "pricing", nil)
}

func (s *Server) renderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, ok := s.Templ[name]
	if !ok {
		http.Error(w, "Template not found: "+name, http.StatusInternalServerError)
		return
	}

	// ExecuteTemplate volá "layout" šablónu, ktorá obsahuje {{template "content" .}}
	// a {{template "title" .}} bloky z príslušnej stránky
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RenderTemplate je exportovaná verzia pre externé použitie (blog routy)
func (s *Server) RenderTemplate(w http.ResponseWriter, r *http.Request, name string) {
	s.renderTemplate(w, r, name, nil)
}

// HandleStatic servuje statické súbory z embedded FS
func (s *Server) HandleStatic(w http.ResponseWriter, r *http.Request) {
	cleanPath := r.URL.Path
	if cleanPath == "" || cleanPath == "/" {
		http.NotFound(w, r)
		return
	}

	// Slúžime z embedded static FS
	data, err := assets.StaticFile(cleanPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Nastav content type podľa prípony
	switch {
	case len(cleanPath) > 4 && cleanPath[len(cleanPath)-4:] == ".css":
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case len(cleanPath) > 4 && cleanPath[len(cleanPath)-4:] == ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	case len(cleanPath) > 5 && cleanPath[len(cleanPath)-5:] == ".xml":
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	case len(cleanPath) > 4 && cleanPath[len(cleanPath)-4:] == ".txt":
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	w.Write(data)
}
