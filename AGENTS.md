# AGENTS.md — ASCENTIA s. r. o. — High-Performance B2B Web Platform

Kompletný technický a biznisový kontext pre AI agentov a vývojárov pracujúcich na ekosystéme Ascentia.

---

## ✦ Biznisová misia

ASCENTIA s. r. o. (IČO: 51858959, Bratislava) je prémiová B2B technologická spoločnosť. Web nie je len prezentácia — je to vysoko konverzný nástroj na generovanie leadov a portál na demonštráciu podnikových AI riešení.

### Monetizačné piliere
1. **Predplatné AI automatizácií (SaaS-lite)** — AI agenti prispôsobení pre slovenských a európskych firemných klientov.
2. **High-Ticket vývoj na mieru** — Prémiové softvérové inžinierstvo v Go (Golang) pre enterprise klientov.
3. **Voice-to-CRM integrácia** — Licencovanie hlasového dopytu pre právnické, lekárske a realitné kancelárie.

**Kontaktný email:** `ascentia@agentmail.to`

---

## 🛠️ Technologický stack

| Vrstva | Technológia | Verzia |
|--------|------------|--------|
| Backend | Go (`net/http` standard mux) | 1.22+ |
| Frontend interakcie | HTMX | 2.0.0 |
| Databáza | SQLite (CGO-free `modernc.org/sqlite`) | 1.29.1 |
| Štýly | Vanilla CSS (CSS variables, glassmorphism) | — |
| Deployment | Fly.io (Frankfurt región) + Docker | — |
| CI/CD | GitHub Actions (auto-deploy na push to `main`) | — |

---

## 📂 Adresárová štruktúra

```
/ascentia-web
├── AGENTS.md                        # Tento súbor — kontext pre AI agentov
├── README.md                        # Verejná dokumentácia
├── Dockerfile                       # Multi-stage Go build (golang:1.22-alpine → alpine)
├── fly.toml                         # Fly.io konfigurácia (región fra, persistent volume)
├── go.mod / go.sum                  # Go modul definitions
├── .github/workflows/fly-deploy.yml # GitHub Actions auto-deploy
├── .agents/skills/htmx/SKILL.md     # HTMX development guidelines (npx skills)
├── cmd/server/main.go               # Vstupný bod — HTTP server, routy, env config
├── db/sqlite.go                     # Databázová vrstva — InitDB, schémy tabuliek
├── db/sqlite_test.go                # TDD testy pre databázu
├── models/models.go                 # Go štruktúry: ChatMessage, ContactInquiry, LeadScore
├── ai/ai.go                         # Pluggable AI Provider interface (Mock, Gemini, OpenAI)
├── ai/ai_test.go                    # TDD testy pre AI provider
├── handlers/
│   ├── handlers.go                  # HTTP handlers, template rendering, routing
│   ├── handlers_test.go             # TDD testy pre HTTP routes
│   ├── contact.go                   # B2B contact form handler (HTMX POST, AI lead scoring)
│   ├── streaming.go                 # SSE streaming pre AI Kompas chat
│   └── voice.go                     # Voice upload handler (multipart audio, MediaRecorder API)
├── templates/
│   ├── layout.html                  # HTML kostra, JSON-LD schémy, nav, footer, cookie bar
│   ├── dashboard.html               # Domov (`/`) — hero, ORSR info, B2B formulár
│   ├── services.html                # Služby (`/services`) — Tier 1/2/3, ROI SVG graf
│   ├── process.html                 # Proces (`/process`) — 4-kroková timeline
│   ├── kompas.html                  # AI Kompas (`/kompas`) — chat so SSE streamovaním
│   ├── voice.html                   # Hlasový dopyt (`/voice-inquiry`) — MediaRecorder API
│   ├── privacy.html                 # GDPR (`/privacy`) — plná EU compliance
│   └── faq.html                     # FAQ (`/faq`) — LLM-friendly Q&A + FAQPage JSON-LD
└── static/
    ├── css/style.css                # Prémiový CSS (glassmorphism, micro-interakcie)
    ├── robots.txt                   # AI crawler povolenia (GPTBot, PerplexityBot, ClaudeBot)
    └── sitemap.xml                  # Sitemap pre vyhľadávače
```

---

## 🔗 Routing map

| URL | Handler | Template | Popis |
|-----|---------|----------|-------|
| `GET /` | `HandleIndex` | `dashboard` | Domov — hero, kredibilita, B2B formulár |
| `GET /services` | `HandleServices` | `services` | 3 tier služby + ROI graf |
| `GET /process` | `HandleProcess` | `process` | 4-krokový proces spolupráce |
| `GET /kompas` | `HandleKompas` | `kompas` | AI chat so SSE streamovaním |
| `GET /voice-inquiry` | `HandleVoice` | `voice` | Hlasový dopyt (MediaRecorder API) |
| `GET /privacy` | `HandlePrivacy` | `privacy` | GDPR dokument |
| `GET /faq` | `HandleFAQ` | `faq` | Časté otázky + FAQPage JSON-LD |
| `POST /api/contact` | `HandleContactSubmit` | — | B2B formulár (HTMX, AI lead scoring) |
| `POST /api/voice-upload` | `HandleVoiceUpload` | — | Audio upload (multipart, AI scoring) |
| `GET /api/stream` | `HandleStreamingAI` | — | SSE streamovanie AI odpovedí |
| `GET /static/*` | `http.FileServer` | — | Statické súbory (CSS, robots.txt, sitemap.xml) |

---

## 🧠 AI Engine (`ai.Provider` interface)

Backend abstrahuje LLM poskytovateľov cez univerzálne rozhranie:

```go
type Provider interface {
    GenerateResponse(ctx context.Context, prompt string, history []models.ChatMessage) (string, error)
    QualifyLead(ctx context.Context, text string) (*models.LeadScore, error)
}
```

Implementovaní provideri:
- **MockProvider** — lokálne testovanie bez API nákladov (default)
- **GeminiProvider** — Google Gemini API (multimodálny prepis hlasu)
- **OpenAIProvider** — OpenAI / DeepSeek / OpenRouter (cez `AI_BASE_URL`)

Konfigurácia cez env premenné:
- `AI_PROVIDER` — `"mock"` | `"gemini"` | `"openai"` | `"openrouter"`
- `AI_MODEL` — názov modelu (napr. `gemini-2.5-flash`, `gpt-4o`)
- `AI_API_KEY` — API kľúč
- `AI_BASE_URL` — vlastná proxy alebo lokálny endpoint

---

## 🗄️ Databázová schéma (SQLite)

```sql
CREATE TABLE chat_messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    session_id TEXT NOT NULL,
    role TEXT NOT NULL,        -- "user" | "assistant"
    message TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE contact_inquiries (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    company TEXT DEFAULT '',
    message TEXT NOT NULL,
    voice_path TEXT DEFAULT '',  -- cesta k audio súboru
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE lead_scores (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    inquiry_id INTEGER,
    score INTEGER DEFAULT 0,     -- 1-100
    budget TEXT DEFAULT '',      -- extrahovaný rozpočet
    urgency TEXT DEFAULT '',     -- low/medium/high
    company_size TEXT DEFAULT '',
    summary TEXT DEFAULT '',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(inquiry_id) REFERENCES contact_inquiries(id) ON DELETE CASCADE
);
```

---

## 🤖 GEO 2026 (Generative Engine Optimization)

Web je optimalizovaný pre AI vyhľadávače (ChatGPT, Perplexity, Gemini):

1. **`robots.txt`** — explicitne povoľuje GPTBot, OAI-SearchBot, Google-Extended, PerplexityBot, ClaudeBot, CCBot, Applebot
2. **`sitemap.xml`** — 7 stránok s prioritami
3. **JSON-LD schémy** na každej stránke:
   - `Organization` — IČO, adresa, email
   - `ProfessionalService` — priceRange, openingHours
   - `WebSite` — inLanguage, description
   - `FAQPage` — 7 LLM-friendly Q&A pre AI citácie
4. **LLM Citation Hooks** — FAQ obsahuje presné definície ako *"Kto na Slovensku vyvíja AI backendy v Go?"*

---

## ⚖️ GDPR Compliance

- Plný súlad s GDPR (EÚ 2016/679) a slovenským zákonom č. 18/2018 Z. z.
- Hlasové nahrávky: vymazané po 30 dňoch
- Dopyty: uchovávané 12 mesiacov
- Chat správy: 6 mesiacov
- Údaje na šifrovaných EÚ serveroch (Fly.io Frankfurt)
- Kontakt pre GDPR: `ascentia@agentmail.to`

---

## 🎨 Dizajnový systém

| Token | Farba | Použitie |
|-------|-------|----------|
| Ascent Deep Blue | `#0A192F` | Dominantné pozadie |
| Cyber Cyan | `#00F0FF` | Akcenty, CTA, linky |
| Neural Grey | `#F4F6F9` | Primárny text |
| Code Charcoal | `#1E2229` | Panely, ohraničenia |

Štýly: glassmorphism (`backdrop-filter: blur`), micro-interakcie (button ripple, nav underline), smooth scroll, SVG path animations, responsive mobile-first.

---

## 🚀 Deployment

### Lokálne
```bash
go run cmd/server/main.go  # → http://localhost:8080
```

### Produkcia (Fly.io)
```bash
fly deploy  # alebo automaticky cez GitHub Actions pri git push
```

### Docker build
```bash
docker build -t ascentia-web .
docker run -p 8080:8080 -v ascentia_data:/data ascentia-web
```

### CI/CD
GitHub Actions (`.github/workflows/fly-deploy.yml`) automaticky nasadí pri push na `main`. Vyžaduje `FLY_API_TOKEN` v GitHub Secrets.

---

## 🧪 TDD

```bash
go test ./... -v
```

Testy pokrývajú:
- `TestInitDB` — vytvorenie databázových tabuliek
- `TestMockProvider` — AI provider GenerateResponse a QualifyLead
- `TestMainRoutes` — HTTP routing pre všetky stránky

---

## 📦 Nainštalované skills

- **htmx** (`.agents/skills/htmx/SKILL.md`) — HTMX development guidelines pre AI agentov. Obsahuje best practices pre `hx-get`, `hx-post`, `hx-target`, `hx-swap`, `hx-trigger` atribúty a server-driven UI patterns.
