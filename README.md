# 🚀 ASCENTIA s. r. o. - Enterprise B2B Web

Vysoko-výkonná, bleskovo rýchla a energeticky nenáročná korporátna B2B webová prezentácia pre spoločnosť **ASCENTIA s. r. o.** (IČO: 51858959).

Tento projekt je navrhnutý s ohľadom na maximálnu maržu (prevádzkové náklady 0 €), bleskové načítanie s 100% skóre v Google Lighthouse a optimálnu štruktúru pre sémantické indexovanie AI crawlerov (GEO/SEO 2026).

---

## 🛠️ Technologický Stack

Táto aplikácia kompletne odmieta ťažké a nestabilné javascriptové frameworky a vsádza na puristickú, bezpečnú a bleskovú architektúru:

- **Go 1.26+ (Golang):** Čisté kognitívne jadro postavené na štandardnej knižnici (`net/http`) bez externých závislostí.
- **HTMX 2.x:** Dynamické asynchrónne interakcie a doručovanie HTML fragmentov zo servera (Server-driven UI) bez prenačítavania stránok.
- **SQLite (CGO-Free):** Štruktúrovaná a bezúdržbová databáza využívajúca driver `modernc.org/sqlite` s nulovou závislosťou od C knižníc.
- **Pure CSS Variable & Glassmorphism:** Žiadny Tailwind, čistý moderný kód s parallax mikro-interakciami na mieru šitými pre éru v roku 2026.

---

## 📂 Architektúra Projektu

```
/ascentia-web
├── AGENTS.md                     # Hlavný kontextový plán (AI-Agent optimalizácia, monetizácia)
├── go.mod / go.sum                # Čistá správa balíkov a modulov v Go (Golang)
├── .env                            # Komerčná a technická konfigurácia prostredia
├── /cmd/server/main.go            # Vstupný bod, štandardný net/http multiplexer a limitery
├── /db/sqlite.go                  # Databázové pripojenie pre modernc.org/sqlite (bez CGO)
├── /db/sqlite_test.go             # TDD testy pre štruktúru tabuliek
├── /models/models.go              # Go štruktúry pre chat, dopyty a lead scoring
├── /ai/ai.go                      # Rozhrania a adaptéry pre Gemini, OpenAI a Mock providerov
├── /ai/ai_test.go                 # Testy pre overenie kognitívnej logiky a skórovania dopytov
├── /handlers/                     # Hlavná logika routovania, contact submission a spracovania leadov
│   ├── handlers.go
│   ├── contact.go
│   ├── streaming.go               # Výkonné streamovanie odpovedí AI po slovách cez Server-Sent Events (SSE)
│   └── handlers_test.go           # HTTP integračné testy routovania
├── /templates/                    # Prémiové, vysoko-konverzné HTML šablóny s JSON-LD sémantickými dátami
│   ├── layout.html                # Hlavná kostra webu (falcon logo s linden listom + privacy cookie lišta)
│   ├── dashboard.html             # Domov. Overená transparentnosť s ORSR dátami pre B2B, CTA formulár
│   ├── services.html              # Rozdelenie ponuky na Tier 1, Tier 2, Tier 3 + interaktívny ROI graf
│   ├── process.html               # 4-kroková B2B časová os spolupráce (Analýza až Nasadenie)
│   ├── kompas.html                # Interaktívny kognitívny sprievodca s bleskovým SSE streamovaním dopytov
│   ├── voice.html                 # Voice-to-CRM rozhranie na vyhodnocovanie hlasových dopytov (max. 30s)
│   └── privacy.html               # Právne a GDPR bezchybné vyhlásenie s odkazom na tvoj email
└── /static/css/style.css          # Luxusný tmavý vizuál (Deep Blue #0A192F + Cyber Cyan #00F0FF)
```

---

## 🚀 Prevádzka a Spustenie

Projekt je pripravený na produkčné nasadenie ako jediný skompilovaný super-rýchly binárny súbor.

### Miestny vývoj:
1. Nakonfiguruj kľúče v súbore `.env`.
2. Spusti server pomocou príkazu:
   ```bash
   go run cmd/server/main.go
   ```
3. Otvor prehliadač na adrese: `http://localhost:8080`.

### TDD Testy:
Spustenie kompletnej sady testov na overenie funkčnosti databázy a smerovania pre splnenie prísnych kvalitatívnych brán:
```bash
go test ./... -v
```

### Výroba optimalizovanej produkčnej binárky:
```bash
go build -ldflags="-s -w" -o server cmd/server/main.go
```
