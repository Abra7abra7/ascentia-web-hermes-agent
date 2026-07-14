package models

import "time"

type ChatMessage struct {
	ID        int64     `json:"id"`
	SessionID string    `json:"session_id"`
	Role      string    `json:"role"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type ContactInquiry struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Company   string    `json:"company"`
	Message   string    `json:"message"`
	VoicePath string    `json:"voice_path"`
	CreatedAt time.Time `json:"created_at"`
}

type LeadScore struct {
	ID          int64     `json:"id"`
	InquiryID   int64     `json:"inquiry_id"`
	Score       int       `json:"score"`   // 1-100 score
	Budget      string    `json:"budget"`  // Extracted budget
	Urgency     string    `json:"urgency"` // low, medium, high
	CompanySize string    `json:"company_size"`
	Summary     string    `json:"summary"`
	CreatedAt   time.Time `json:"created_at"`
}
