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
├── .agents/skills/htmx/             # HTMX development guidelines (npx skills)
├── .agents/skills/[47 skills]/     # Marketing skills (coreyhaines31/marketingskills)
├── .agents/product-marketing.md    # Product context pre marketing skills
├── cmd/server/main.go              # Vstupný bod — HTTP server, routy, env config
├── db/sqlite.go                     # Databázová vrstva — InitDB, schémy tabuliek
├── db/sqlite_test.go                # TDD testy pre databázu
├── models/models.go                 # Go štruktúry: ChatMessage, ContactInquiry, LeadScore
├── ai/ai.go                         # Pluggable AI Provider interface (Mock, Gemini, OpenAI)
├── ai/ai_test.go                    # TDD testy pre AI provider
├── handlers/
│   ├── handlers.go                  # HTTP handlers, template rendering, routing
│   ├── handlers_test.go             # TDD testy pre HTTP routes
│   ├── contact.go                   # B2B contact form (HTMX POST, AI lead scoring, confirmation email)
│   ├── streaming.go                  # SSE streaming pre AI Kompas chat
│   ├── voice.go                     # Voice upload (multipart audio, AI scoring, notification email)
│   ├── email.go                     # Resend REST API — sendEmail, sendLeadNotification, sendVoiceLeadNotification, sendClientConfirmation
│   └── followup.go                  # Automated 24h follow-up scheduler (hourly check, AI analysis email)
├── templates/
│   ├── layout.html                  # HTML kostra, JSON-LD, nav, footer, GA4, OG/Twitter meta tagy
│   ├── dashboard.html               # Domov (`/`) — hero, ORSR info, B2B formulár
│   ├── services.html                # Služby (`/services`) — Tier 1/2/3, ROI SVG graf
│   ├── process.html                 # Proces (`/process`) — 4-kroková timeline
│   ├── kompas.html                  # AI Kompas (`/kompas`) — chat so SSE streamovaním
│   ├── voice.html                   # Hlasový dopyt (`/voice-inquiry`) — MediaRecorder API, 60s limit, transcript
│   ├── consultation.html            # Konzultácia (`/consultation`) — Cal.com inline embed, Service JSON-LD
│   ├── blog.html                    # Blog listing (`/blog`) — 3 články
│   ├── blog_go_pre_enterprise.html  # Blog: Go vs Node.js vs Python + BlogPosting JSON-LD
│   ├── blog_ai_pravnikom.html       # Blog: AI pre právnikov + BlogPosting JSON-LD
│   ├── blog_voice_crm_case.html     # Blog: Voice-to-CRM case study + BlogPosting JSON-LD
│   ├── pricing.html                 # Cenník (`/pricing`) — Tier 1/2/3 + SaaS predplatné
│   ├── privacy.html                 # GDPR (`/privacy`) — plná EU compliance
│   └── faq.html                     # FAQ (`/faq`) — LLM-friendly Q&A + FAQPage JSON-LD
├── static/
│   ├── css/style.css                # Prémiový CSS (glassmorphism, micro-interakcie, blog, pricing, consultation)
│   ├── img/og-image.svg             # Open Graph image (1200x630, ASCENTIA branding)
│   ├── robots.txt                   # AI crawler povolenia (GPTBot, PerplexityBot, ClaudeBot)
│   └── sitemap.xml                  # Sitemap (13 stránok s prioritami)
└── docs/
    ├── marketing-strategy.md         # Kompletná 9-sekčná marketingová stratégia
    ├── linkedin-posts.md             # 5 ready-to-post LinkedIn príspevkov
    ├── cold-email-templates.md       # 3 cold email šablóny (právnici, CTO, realitné kancelárie)
    └── linkedin-outreach-list.md     # 34 cieľových firiem (právo, RE, tech, fintech, zdravotníctvo)
```

---

## 🔗 Routing map

| URL | Handler | Template | Popis |
|-----|---------|----------|-------|
| `GET /` | `HandleIndex` | `dashboard` | Domov — hero, kredibilita, B2B formulár |
| `GET /services` | `HandleServices` | `services` | 3 tier služby + ROI graf |
| `GET /pricing` | `HandlePricing` | `pricing` | Cenník — Tier 1/2/3 + SaaS predplatné |
| `GET /process` | `HandleProcess` | `process` | 4-krokový proces spolupráce |
| `GET /kompas` | `HandleKompas` | `kompas` | AI chat so SSE streamovaním |
| `GET /voice-inquiry` | `HandleVoice` | `voice` | Hlasový dopyt (MediaRecorder API, 60s, transcript) |
| `GET /consultation` | `HandleConsultation` | `consultation` | Cal.com inline embed bookovanie |
| `GET /blog` | `HandleBlog` | `blog` | Blog listing (3 články) |
| `GET /blog/go-pre-enterprise` | `RenderTemplate` | `blog_go_pre_enterprise` | Blog: Go vs Node.js vs Python |
| `GET /blog/ai-automatizacia-pravnikom` | `RenderTemplate` | `blog_ai_pravnikom` | Blog: AI pre právnické kancelárie |
| `GET /blog/voice-to-crm-case-study` | `RenderTemplate` | `blog_voice_crm_case` | Blog: Voice-to-CRM +340% |
| `GET /privacy` | `HandlePrivacy` | `privacy` | GDPR dokument |
| `GET /faq` | `HandleFAQ` | `faq` | Časté otázky + FAQPage JSON-LD |
| `POST /api/contact` | `HandleContactSubmit` | — | B2B formulár (HTMX, AI lead scoring, confirmation email) |
| `POST /api/voice-upload` | `HandleVoiceUpload` | — | Audio upload (multipart, AI scoring, notification email) |
| `GET /api/stream` | `HandleStreamingAI` | — | SSE streamovanie AI odpovedí |
| `GET /static/*` | `http.FileServer` | — | Statické súbory (CSS, img, robots.txt, sitemap.xml) |

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
    follow_up_sent INTEGER DEFAULT 0,  -- 0=nedošlo, 1=odoslaný follow-up
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
2. **`sitemap.xml`** — 13 stránok s prioritami
3. **JSON-LD schémy** na každej stránke:
   - `Organization` — IČO, adresa, email
   - `ProfessionalService` — priceRange, openingHours
   - `WebSite` — inLanguage, description
   - `FAQPage` — 7 LLM-friendly Q&A pre AI citácie
   - `BlogPosting` — headline, datePublished, author, keywords (3 blog články)
   - `Service` — bezplatná konzultácia, price=0
4. **LLM Citation Hooks** — FAQ obsahuje presné definície ako *"Kto na Slovensku vyvíja AI backendy v Go?"*
5. **Open Graph + Twitter Card** — og:image (SVG 1200x630), og:site_name, twitter:card
6. **Google Analytics 4** — `anonymize_ip: true` (GDPR compliant), placeholder `G-XXXXXXX`

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

### HTMX
- **htmx** (`.agents/skills/htmx/`) — HTMX development guidelines pre AI agentov. Best practices pre `hx-get`, `hx-post`, `hx-target`, `hx-swap`, `hx-trigger` atribúty a server-driven UI patterns.

### Marketing skills (47 zo coreyhaines31/marketingskills)
Nainštalované v `.agents/skills/`:
- **Social media:** `social`, `community-marketing`, `video`
- **Content:** `content-strategy`, `copywriting`, `copy-editing`, `image`
- **Email:** `cold-email`, `emails`, `sms`
- **SEO/GEO:** `ai-seo`, `seo-audit`, `programmatic-seo`, `schema`
- **Ads:** `ads`, `ad-creative`, `analytics`
- **Strategy:** `marketing-plan`, `marketing-ideas`, `marketing-psychology`, `marketing-loops`, `marketing-council`
- **Sales:** `prospecting`, `sales-enablement`, `pricing`, `offers`, `referrals`
- **CRO:** `cro`, `ab-testing`, `popups`, `signup`, `paywalls`
- **PR:** `public-relations`, `co-marketing`, `competitor-profiling`, `competitors`
- **Customer:** `customer-research`, `onboarding`, `churn-prevention`, `revops`
- **Other:** `directory-submissions`, `free-tools`, `launch`, `lead-magnets`, `product-marketing`, `site-architecture`

### Product marketing context
- **`.agents/product-marketing.md`** — ICP, produkty, value proposition, competitive advantage, brand voice, tech stack. Používajú marketing skills ako vstupný kontext.

---

## 📧 Email systém (Resend REST API)

### Architektúra
```
Go backend → HTTP POST api.resend.com/emails → from: ASCENTIA Web <ascentia@marianstancik.dev> → to: ascentia@agentmail.to / klient
```

### Handler: `handlers/email.go`
- `loadEmailConfig()` — načíta `SMTP_PASS` z env (Fly.io secret), `from` = `ascientia@marianstancik.dev`
- `sendEmail(config, to, subject, body)` — všeobecné odoslanie cez Resend REST API (JSON payload cez `encoding/json.Marshal`)
- `sendLeadNotification(name, email, company, message)` — notifikácia na `ascentia@agentmail.to`
- `sendVoiceLeadNotification(name, email, company, voicePath)` — hlasový dopyt notifikácia
- `sendClientConfirmation(name, email, source)` — potvrdenie klientovi na jeho email

### Email flow (automatizovaný)

| Čas | Email | Príjemca | Obsah |
|-----|-------|----------|-------|
| **0h** | Potvrdenie prijatia dopytu | Klient (ich email) | "Ďakujeme za váš dopyt... architekt vás kontaktuje do 24h" |
| **0h** | Notifikácia o dopyte | `ascentia@agentmail.to` | Meno, email, spoločnosť, správa/audio cesta |
| **24h** | AI analýza dopytu (follow-up) | Klient (ich email) | Lead score, rozpočet, urgencia, návrh riešenia, CTA na AI Kompas |

### Handler: `handlers/followup.go` — Automated 24h follow-up scheduler
- `StartFollowUpScheduler()` — spustí sa 5 min po štarte servera, kontroluje každú hodinu
- `RunFollowUpCheck()` — SQL query: `WHERE created_at < datetime('now', '-24 hours') AND follow_up_sent = 0`
- `generateInquiryAnalysis()` — vytvorí AI analýzu (lead score, rozpočet, priorita, návrh riešenia)
- `sendFollowUpEmail()` — odošle follow-up email, urgency "high" → `[URGENT]` prefix v subjecte
- Po odoslaní sa `follow_up_sent` nastaví na 1 (aby sa neodosielal duplicitne)

### Bezpečnosť
- **Žiadne hardcoded API kľúče** — `SMTP_PASS` (Resend API key) výhradne cez `fly secrets set SMTP_PASS=[REDACTED]`
- **`fly.toml`** neobsahuje žiadne secrets — iba `PORT`, `DB_PATH`, `AI_PROVIDER` v `[env]`

---

## 📅 Cal.com bookovací systém

- **Stránka:** `/consultation` (`templates/consultation.html`)
- **Embed:** Oficiálny Cal.com inline embed (nie iframe) — `embed.js`, `#my-cal-inline-30min`
- **Event:** `ascentia/30min` na `app.cal.com`
- **Layout:** Month view, slots view na malých obrazovkách
- **JSON-LD:** `Service` schema (price=0, EUR, InStock)
- **Konfigurácia:** Cal.com účet `ascentia` — nastavené cez web UI na cal.com

---

## 📝 Blog sekcia

3 SEO-optimalizované články s `BlogPosting` JSON-LD schémou:

1. **`/blog/go-pre-enterprise`** — "Prečo Go (Golang) pre enterprise backend v roku 2026?"
   - Benchmarky: Go 47k req/s vs Node.js 18k vs Python 8k
   - Náklady: €5/mesiac (Go) vs €400/mesiac (Node.js)
   - CTA: Konzultácia + Služby

2. **`/blog/ai-automatizacia-pravnikom`** — "AI automatizácia pre právnické kancelárie"
   - Voice-to-CRM pre advokátov, GDPR súlad
   - ROI: 1,200-2,400% mesačne
   - CTA: Demo + Hlasový dopyt

3. **`/blog/voice-to-crm-case-study`** — "Voice-to-CRM: +340% konverzia"
   - Konverzia 2.1% → 9.2%, lead quality 72/100
   - Technická architektúra (Go + HTMX + SQLite)
   - CTA: Konzultácia + Hlasový dopyt

---

## 💰 Cenník (`/pricing`)

| Tier | Názov | Cena | Obsah |
|------|-------|------|-------|
| 1 | AI Systém | €2,000/projekt | AI agent + CRM integrácia + lead scoring |
| 2 | Go Backend | €8,000/projekt | Go + HTMX + Voice-to-CRM + CI/CD (najpopulárnejší) |
| 3 | R&D | €15,000+/projekt | Hardvér-softvér + AI fine-tuning + prototyp |

**SaaS predplatné:**
- Jednotlivec: €200/mesiac (1 endpoint, 100 dopytov/mesiac)
- Tím: €500/mesiac (5 endpointov, 500 dopytov/mesiac, analytics)

---

## 📊 Marketingová stratégia

Dokumentácia v `docs/`:

### `docs/marketing-strategy.md` — 9 sekcií:
1. **ICP** — 3 typy klientov (Tier 2 Go Backend €8k, Voice-to-CRM SaaS €200-500/mes, Tier 1 AI €2k)
2. **Content pillars** — 5 pilierov (technická 30%, AI 25%, BTS 20%, vzdelávanie 15%, promo 10%)
3. **LinkedIn stratégia** — 4 posty/týždeň (Pondelok/Utorok/Štvrtok/Piatok), 8 carousel nápadov
4. **Cold email kampaň** — 4-touch sekvencia (Deň 0/3/7/14), 34 cieľových firiem
5. **AI-SEO / GEO 2026** — kľúčové slová, blog 2x/mesiac, programmatic SEO
6. **CRO** — exit-intent popup, sticky CTA, A/B test
7. **PR** — StartupGrind, Hospodárske noviny, tech podcasty
8. **Marketing ops stack** — €5/mesiac (GA4, Resend, AgentMail, Cal.com, Fly.io)
9. **90-dňový plán** — KPIs: Mesiac 3 → 3,000 návštevníkov, 50 leadov, €25,000 revenue

### `docs/linkedin-posts.md` — 5 ready-to-post príspevkov
### `docs/cold-email-templates.md` — 3 šablóny (právnici, CTO, realitné kancelárie)
### `docs/linkedin-outreach-list.md` — 34 firiem v 5 segmentoch (právo, RE, tech, fintech, zdravotníctvo)

### Cron job: LinkedIn post reminder
- **Schedule:** Pondelok, Utorok, Štvrtok, Piatok o 8:00
- **Deliver:** Telegram (origin chat)
- **Typ:** LLM-driven — generuje nový príspevok podľa content kalendára

---

## 🔧 Externé integrácie

| Služba | Účel | Konfigurácia |
|--------|------|---------------|
| **Fly.io** | Hosting (Frankfurt, EÚ) | App: `ascentia-web-hermes-agent`, region `fra`, port 8080 |
| **Resend API** | Email notifikácie | Domain: `marianstancik.dev` (verified), API key cez `fly secrets set SMTP_PASS` |
| **AgentMail** | Inbox pre Ascentiu | `ascentia@agentmail.to` — prijíma lead notifikácie |
| **Cal.com** | Bookovanie konzultácií | Event: `ascentia/30min` na `app.cal.com` |
| **GitHub** | Repo + CI/CD | `git@github.com:Abra7abra7/ascentia-web-hermes-agent.git` |
| **Google Analytics 4** | Web analytics | Placeholder `G-XXXXXXX` — nahradiť reálnym ID |

### Fly.io volume
- **Name:** `ascentia_data` (1GB)
- **Mount:** `/data`
- **Obsah:** SQLite databáza (`/data/ascentia.db`), hlasové nahrávky (`/data/voice_uploads/`)

### Env premenné (Fly.io)
| Premenná | Hodnota | Popis |
|----------|--------|------|
| `PORT` | `8080` | HTTP port |
| `DB_PATH` | `/data/ascentia.db` | SQLite cesta |
| `AI_PROVIDER` | `mock` | AI provider (mock/gemini/openai/openrouter) |
| `SMTP_PASS` | [REDACTED] | Resend API key (cez `fly secrets set`) |
