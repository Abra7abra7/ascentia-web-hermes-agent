package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"ascentia-web/ai"
)

type Server struct {
	DB   *sql.DB
	AI   ai.Provider
	Templ *template.Template
}

func NewServer(db *sql.DB, aiProvider ai.Provider) (*Server, error) {
	// Parse all HTML files from templates directory
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		return nil, err
	}

	return &Server{
		DB:    db,
		AI:    aiProvider,
		Templ: tmpl,
	}, nil
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
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

func (s *Server) renderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// Pre HTMX požiadavky môžeme vracať len čiastočný fragment,
	// ale pre plné načítania vrátime kompletný layout.
	// Tu pre jednoduchosť a spoľahlivosť rendrujeme celú šablónu, ktorá zahŕňa layout interným mechanizmom
	err := s.Templ.ExecuteTemplate(w, name+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
