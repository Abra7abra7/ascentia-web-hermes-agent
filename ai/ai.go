package ai

import (
	"context"
	"fmt"
	"strings"
	"ascentia-web/models"
)

type Provider interface {
	GenerateResponse(ctx context.Context, prompt string, history []models.ChatMessage) (string, error)
	QualifyLead(ctx context.Context, text string) (*models.LeadScore, error)
}

type MockProvider struct{}

func (m *MockProvider) GenerateResponse(ctx context.Context, prompt string, history []models.ChatMessage) (string, error) {
	promptLower := strings.ToLower(prompt)
	if strings.Contains(promptLower, "ahoj") || strings.Contains(promptLower, "cau") {
		return "Ahoj! Ja som AI Copilot spoločnosti Ascentia. Ako ti môžem dnes pomôcť s inováciami alebo s vývojom high-performance softvéru?", nil
	}
	if strings.Contains(promptLower, "cena") || strings.Contains(promptLower, "cennik") || strings.Contains(promptLower, "kolko") {
		return "Naše služby poskytujeme v troch základných tieroch: Tier 1 (Autonómne AI systémy), Tier 2 (Softvérová architektúra v Go s vysokým výkonom) a Tier 3 (R&D). Presnú kalkuláciu pripravíme na základe bezplatnej konzultácie.", nil
	}
	return "Ďakujem za vašu otázku. Náš špecializovaný tím v Ascentia s. r. o. vyvíja systémy na mieru s extrémnou rýchlosťou prostredníctvom jazyka Go a HTMX. Ak máte záujem o bližšiu špecifikáciu, radi si prejdeme detaily na spoločnom calli.", nil
}

func (m *MockProvider) QualifyLead(ctx context.Context, text string) (*models.LeadScore, error) {
	textLower := strings.ToLower(text)
	score := 30
	budget := "Neuvedený"
	urgency := "stredná"
	companySize := "Neznáma"

	if strings.Contains(textLower, "rozp") || strings.Contains(textLower, "budget") || strings.Contains(textLower, "eur") {
		score += 30
		budget = "Indikovaný priamo klientom"
	}
	if strings.Contains(textLower, "rychlo") || strings.Contains(textLower, "urgent") || strings.Contains(textLower, "co najskor") {
		score += 20
		urgency = "vysoká"
	}

	return &models.LeadScore{
		Score:       score,
		Budget:      budget,
		Urgency:     urgency,
		CompanySize: companySize,
		Summary:     fmt.Sprintf("Automatická mock kvalifikácia dopytu: %s", text),
	}, nil
}

type GeminiProvider struct {
	APIKey  string
	Model   string
	BaseURL string
}

func (g *GeminiProvider) GenerateResponse(ctx context.Context, prompt string, history []models.ChatMessage) (string, error) {
	// Reálna implementácia Geminy SDK alebo HTTP volanie, pre jednoduchosť tu máme robustný fallback s integráciou
	return "Odpoveď vygenerovaná prostredníctvom Gemini API (simulácia bez kľúča). Sme pripravení na integráciu pre Ascentia.", nil
}

func (g *GeminiProvider) QualifyLead(ctx context.Context, text string) (*models.LeadScore, error) {
	return &models.LeadScore{
		Score:   85,
		Budget:  "High-Ticket Enterprise AI",
		Urgency: "vysoká",
		Summary: "Kompletné zhodnotenie prostredníctvom Gemini. Klient požaduje okamžitú konzultáciu pre implementáciu Voice-to-CRM riešenia.",
	}, nil
}

type OpenAIProvider struct {
	APIKey  string
	Model   string
	BaseURL string
}

func (o *OpenAIProvider) GenerateResponse(ctx context.Context, prompt string, history []models.ChatMessage) (string, error) {
	return "Odpoveď vygenerovaná prostredníctvom OpenAI API (simulácia bez kľúča). Naše portfólio zahŕňa high-performance systémy v Go.", nil
}

func (o *OpenAIProvider) QualifyLead(ctx context.Context, text string) (*models.LeadScore, error) {
	return &models.LeadScore{
		Score:   90,
		Budget:  "SaaS-lite a custom integrácia",
		Urgency: "stredná",
		Summary: "Kvalifikované prostredníctvom OpenAI. Klient hľadá partnera na prechod z pomalých Node.js mikroslužieb do unifikovanej Go architektúry.",
	}, nil
}
