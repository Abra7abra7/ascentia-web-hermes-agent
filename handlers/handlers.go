package handlers

import (
	"ascentia-web/ai"
	"database/sql"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Server struct {
	DB    *sql.DB
	AI    ai.Provider
	Templ map[string]*template.Template
	GA4ID string
}

// PageData prenáša spoločné premenné do všetkých šablón
type PageData struct {
	GA4ID string
}

func NewServer(db *sql.DB, aiProvider ai.Provider) (*Server, error) {
	templates := make(map[string]*template.Template)

	// Načítame layout ako základ pre všetky stránky
	layoutPath := filepath.Join("templates", "layout.html")

	// Parzneme každú stránku samostatne s layoutom, aby sa predišlo kolízii define blokov
	pages := []string{"dashboard", "services", "process", "privacy", "kompas", "voice", "faq", "consultation", "blog", "blog_go_pre_enterprise", "blog_ai_pravnikom", "blog_voice_crm_case", "pricing", "admin"}
	for _, page := range pages {
		pagePath := filepath.Join("templates", page+".html")
		tmpl, err := template.New("layout").ParseFiles(layoutPath, pagePath)
		if err != nil {
			return nil, err
		}
		templates[page] = tmpl
	}

	return &Server{
		DB:    db,
		AI:    aiProvider,
		Templ: templates,
		GA4ID: getGA4ID(),
	}, nil
}

// getGA4ID načíta GA4 ID z env alebo použije placeholder
func getGA4ID() string {
	return os.Getenv("GA4_ID")
}

// pageData vráti spoločné PageData pre všetky šablóny
func (s *Server) pageData() PageData {
	return PageData{GA4ID: s.GA4ID}
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	s.renderTemplate(w, r, "dashboard", s.pageData())
}

func (s *Server) HandleServices(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "services", s.pageData())
}

func (s *Server) HandleProcess(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "process", s.pageData())
}

func (s *Server) HandlePrivacy(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "privacy", s.pageData())
}

func (s *Server) HandleKompas(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "kompas", s.pageData())
}

func (s *Server) HandleVoice(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "voice", s.pageData())
}

func (s *Server) HandleFAQ(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "faq", s.pageData())
}

func (s *Server) HandleConsultation(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "consultation", s.pageData())
}

func (s *Server) HandleBlog(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "blog", s.pageData())
}

func (s *Server) HandlePricing(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, r, "pricing", s.pageData())
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

// HandleStatic servuje statické súbory (CSS, JS, obrázky)
func (s *Server) HandleStatic(w http.ResponseWriter, r *http.Request) {
	// Bezpečnostná kontrola: zamedzenie directory traversal
	cleanPath := filepath.Clean(r.URL.Path)
	if cleanPath == "" || cleanPath == "/" {
		http.NotFound(w, r)
		return
	}

	// Servujeme z static adresára
	staticDir := "static"
	filePath := filepath.Join(staticDir, cleanPath)

	// Overíme, či súbor existuje
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, filePath)
}
