package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"ascentia-web/ai"
)

type Server struct {
	DB    *sql.DB
	AI    ai.Provider
	Templ map[string]*template.Template
}

func NewServer(db *sql.DB, aiProvider ai.Provider) (*Server, error) {
	templates := make(map[string]*template.Template)

	// Načítame layout ako základ pre všetky stránky
	layoutPath := filepath.Join("templates", "layout.html")

	// Parzneme každú stránku samostatne s layoutom, aby sa predišlo kolízii define blokov
	pages := []string{"dashboard", "services", "process", "privacy", "kompas", "voice", "faq"}
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
