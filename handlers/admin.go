package handlers

import (
	"database/sql"
	"net/http"
	"os"
)

// HandleAdmin zobrazí admin dashboard so všetkými dopytmi a lead score
func (s *Server) HandleAdmin(w http.ResponseWriter, r *http.Request) {
	// Auth kontrola
	password := r.Header.Get("X-Admin-Pass")
	if password == "" {
		password = r.URL.Query().Get("key")
	}
	adminPass := os.Getenv("ADMIN_PASS")
	if adminPass == "" {
		adminPass = "ascentia-admin-2026"
	}
	if password != adminPass {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Načítaj všetky dopyty s lead score
	rows, err := s.DB.Query(`
		SELECT ci.id, ci.name, ci.email, ci.company, ci.message, ci.voice_path, ci.follow_up_sent, ci.created_at,
		       COALESCE(ls.score, 0), COALESCE(ls.budget, ''), COALESCE(ls.urgency, ''), COALESCE(ls.summary, '')
		FROM contact_inquiries ci
		LEFT JOIN lead_scores ls ON ls.inquiry_id = ci.id
		ORDER BY ci.created_at DESC
		LIMIT 100
	`)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type AdminInquiry struct {
		ID         int
		Name       string
		Email      string
		Company    string
		Message    string
		VoicePath  string
		FollowUp   int
		CreatedAt  string
		Score      int
		Budget     string
		Urgency    string
		Summary    string
	}

	var inquiries []AdminInquiry
	for rows.Next() {
		var i AdminInquiry
		var voicePath sql.NullString
		err := rows.Scan(&i.ID, &i.Name, &i.Email, &i.Company, &i.Message, &voicePath, &i.FollowUp, &i.CreatedAt, &i.Score, &i.Budget, &i.Urgency, &i.Summary)
		if err != nil {
			continue
		}
		i.VoicePath = voicePath.String
		inquiries = append(inquiries, i)
	}

	// Štatistiky
	var totalLeads, totalFollowUp, highUrgency int
	s.DB.QueryRow("SELECT COUNT(*) FROM contact_inquiries").Scan(&totalLeads)
	s.DB.QueryRow("SELECT COUNT(*) FROM contact_inquiries WHERE follow_up_sent = 1").Scan(&totalFollowUp)
	s.DB.QueryRow("SELECT COUNT(*) FROM lead_scores WHERE urgency = 'high'").Scan(&highUrgency)

	data := struct {
		Inquiries     []AdminInquiry
		TotalLeads    int
		TotalFollowUp int
		HighUrgency   int
	}{
		Inquiries:     inquiries,
		TotalLeads:    totalLeads,
		TotalFollowUp: totalFollowUp,
		HighUrgency:   highUrgency,
	}

	s.renderTemplate(w, r, "admin", data)
}
