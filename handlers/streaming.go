package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"ascentia-web/models"
)

func (s *Server) HandleStreamingAI(w http.ResponseWriter, r *http.Request) {
	// Set headers for Server-Sent Events (SSE)
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	prompt := r.URL.Query().Get("prompt")
	sessionID := r.URL.Query().Get("session_id")
	if sessionID == "" {
		sessionID = "default"
	}

	if prompt == "" {
		fmt.Fprintf(w, "data: %s\n\n", "Prosím, zadajte vašu otázku pre AI Kompas.")
		flusher.Flush()
		return
	}

	// Uložíme dotaz usera do db
	_, _ = s.DB.Exec("INSERT INTO chat_messages (session_id, role, message) VALUES (?, ?, ?)", sessionID, "user", prompt)

	// Získame históriu správ pre kontext
	rows, err := s.DB.Query("SELECT role, message FROM chat_messages WHERE session_id = ? ORDER BY created_at ASC LIMIT 10", sessionID)
	var history []models.ChatMessage
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var msg models.ChatMessage
			if err := rows.Scan(&msg.Role, &msg.Message); err == nil {
				history = append(history, msg)
			}
		}
	}

	// Získanie odpovede z AI Providers
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	reply, err := s.AI.GenerateResponse(ctx, prompt, history)
	if err != nil {
		reply = "Ospravedlňujeme sa, ale spojenie s naším AI kognitívnym jadrom bolo prerušené. Skúste to prosím znova."
	}

	// Simulujeme reálne streamovanie po slovách na frontende pre vynikajúci používateľský UX zážitok
	words := []string{}
	// Rozdelíme odpoveď po slovách, aby sme predviedli streamovanie cez SSE
	for _, word := range []string{"Ascentia", "AI-Engine:", "Konfigurujeme", "Váš", "softvér", "na", "mieru.", "Riešenie", "stojí", "na", "čistej", "funkčnosti,", "rýchlosti", "Go", "a", "dynamickom", "HTMX.", "Vaša", "otázka", "bola", "spracovaná."} {
		words = append(words, word)
	}

	// Použijeme aj skutočný vygenerovaný text od AI providera
	realWords := []string{}
	// Pre zjednodušenie budeme po slovách streamovať celú vygenerovanú odpoveď
	for _, word := range rURLQuerySplit(reply) {
		realWords = append(realWords, word)
	}

	if len(realWords) > 0 {
		words = realWords
	}

	// Streamujeme slová
	for _, word := range words {
		select {
		case <-r.Context().Done():
			return
		default:
			fmt.Fprintf(w, "data: %s \n\n", word)
			flusher.Flush()
			time.Sleep(100 * time.Millisecond) // micro-interakcia pre simuláciu premýšľania
		}
	}

	// Uložíme odpoveď systému do db
	_, _ = s.DB.Exec("INSERT INTO chat_messages (session_id, role, message) VALUES (?, ?, ?)", sessionID, "assistant", reply)

	fmt.Fprintf(w, "data: [DONE]\n\n")
	flusher.Flush()
}

// Pomocná domáca funkcia na rozbitie stringu po slovách bez ext dependencies
func rURLQuerySplit(s string) []string {
	var words []string
	current := ""
	for _, char := range s {
		if char == ' ' || char == '\n' || char == '\t' {
			if current != "" {
				words = append(words, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}
	if current != "" {
		words = append(words, current)
	}
	return words
}
