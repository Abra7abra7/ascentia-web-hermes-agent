package api

import (
	"log"
	"net/http"
	"os"
	"ascentia-web/ai"
	"ascentia-web/db"
	"ascentia-web/handlers"
)

var server *handlers.Server
var initErr error

func init() {
	// AI Provider
	providerType := os.Getenv("AI_PROVIDER")
	if providerType == "" {
		providerType = "mock"
	}

	var aiProvider ai.Provider
	switch providerType {
	case "gemini":
		aiProvider = &ai.GeminiProvider{
			APIKey:  os.Getenv("AI_API_KEY"),
			Model:   os.Getenv("AI_MODEL"),
			BaseURL: os.Getenv("AI_BASE_URL"),
		}
	case "openai", "openrouter":
		aiProvider = &ai.OpenAIProvider{
			APIKey:  os.Getenv("AI_API_KEY"),
			Model:   os.Getenv("AI_MODEL"),
			BaseURL: os.Getenv("AI_BASE_URL"),
		}
	default:
		aiProvider = &ai.MockProvider{}
	}

	// DB — na Verele používame /tmp pre read-write
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "/tmp/ascentia.db"
	}

	sqliteDB, err := db.InitDB(dbPath)
	if err != nil {
		initErr = err
		log.Printf("[API] DB init error: %v", err)
		return
	}

	server, err = handlers.NewServer(sqliteDB, aiProvider)
	if err != nil {
		initErr = err
		log.Printf("[API] Server init error: %v", err)
	}
}

// Handler je Vercel serverless entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	if initErr != nil {
		http.Error(w, "Server initialization failed: "+initErr.Error(), http.StatusInternalServerError)
		return
	}

	mux := http.NewServeMux()

	// Statické súbory z embedded FS
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		// Odstránime /static/ prefix
		path := r.URL.Path
		if len(path) > 8 && path[:8] == "/static/" {
			r.URL.Path = path[8:]
		}
		server.HandleStatic(w, r)
	})

	// Stránky
	mux.HandleFunc("/", server.HandleIndex)
	mux.HandleFunc("/services", server.HandleServices)
	mux.HandleFunc("/pricing", server.HandlePricing)
	mux.HandleFunc("/process", server.HandleProcess)
	mux.HandleFunc("/privacy", server.HandlePrivacy)
	mux.HandleFunc("/kompas", server.HandleKompas)
	mux.HandleFunc("/voice-inquiry", server.HandleVoice)
	mux.HandleFunc("/faq", server.HandleFAQ)
	mux.HandleFunc("/consultation", server.HandleConsultation)
	mux.HandleFunc("/blog", server.HandleBlog)
	mux.HandleFunc("/blog/go-pre-enterprise", func(w http.ResponseWriter, r *http.Request) {
		server.RenderTemplate(w, r, "blog_go_pre_enterprise")
	})
	mux.HandleFunc("/blog/ai-automatizacia-pravnikom", func(w http.ResponseWriter, r *http.Request) {
		server.RenderTemplate(w, r, "blog_ai_pravnikom")
	})
	mux.HandleFunc("/blog/voice-to-crm-case-study", func(w http.ResponseWriter, r *http.Request) {
		server.RenderTemplate(w, r, "blog_voice_crm_case")
	})

	// POST API
	mux.HandleFunc("/api/contact", server.HandleContactSubmit)
	mux.HandleFunc("/api/stream", server.HandleStreamingAI)
	mux.HandleFunc("/api/voice-upload", server.HandleVoiceUpload)

	mux.ServeHTTP(w, r)
}
