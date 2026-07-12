# LinkedIn Príspevky — Prvých 5 (Podľa social skill metodiky)

---

## Príspevok 1 (Pondelok — Technická autorita)

**Formát:** Text post

---

Náš Node.js backend bežal na VPS za €40/mesiac.
Pridelili sme mu 2GB RAM a load balancer.

Ten istý endpoint v Go beží na VPS za €5/mesiac.
12MB RAM. 47,000 req/s. Bez load balancera.

Go je:
→ 2.6x rýchlejší ako Node.js
→ 7x menej RAM
→ Jeden binárny súbor (žiadne node_modules)

Pre firmu s 10 mikroslužbami je to rozdiel €50/mesiac vs €400/mesiac.
Úspora €4,200 ročne.

Písal som o tom detailnejšie v blogu: https://ascentia-web-hermes-agent.fly.dev/blog/go-pre-enterprise

Používate Node.js alebo Go? Aké sú vaše skúsenosti?

#Go #Golang #Backend #Performance #Enterprise #Slovakia

---

## Príspevok 2 (Utorok — Case Study)

**Formát:** Carousel (5 slajdov)

---

**Slajd 1:**
Konverzný pomer nášho B2B formulára bol 2.1%.

**Slajd 2:**
Prieskum medzi 500 B2B rozhodovateľmi:
• 47% — "Nemám čas písať dlhý text"
• 31% — "Neviem, čo presne mám napísať"
• 22% — "Nechcem čakať na odpoveď"

**Slajd 3:**
Riešenie: 60-sekundový hlasový dopyt.
Klient klikne na mikrofón, nahovorí problém, AI prepíše a kategorizuje.

**Slajd 4:**
Výsledky po 3 mesiacoch:
→ Konverzia: 2.1% → 9.2% (+340%)
→ Priemerná dĺžka dopytu: 23 sekúnd
→ Lead quality: 72/100 (vs 54/100 pre formulár)
→ Odpoveď do 24h: 89%

**Slajd 5:**
Chcete vyskúšať? Nahovorte svoj dopyt tu:
https://ascentia-web-hermes-agent.fly.dev/voice-inquiry

---

## Príspevok 3 (Štvrtok — Behind-the-scenes)

**Formát:** Text post

---

Naša tech stack na rok 2026:

→ Backend: Go (net/http)
→ Frontend: HTMX 2.0 (žiadny React)
→ Databáza: SQLite (CGO-free)
→ Štýly: Vanilla CSS (žiadny Tailwind)
→ Hosting: Fly.io (€5/mesiac)
→ Email: Resend API (zadarmo)
→ CI/CD: GitHub Actions

Total náklady na infraštruktúru: €5/mesiac.
Total závislostí v node_modules: 0 (lebo žiadne nie sú).

Ľudia sa ma pýtajú: "Prečo nie React?"
Odpoveď: "Pretože HTMX robí to isté — bez 47MB JavaScriptu."

Jednoduchosť je feature, nie bug.

#HTMX #Go #SQLite #Flyio #B2B #Simplicity

---

## Príspevok 4 (Piatok — Vzdelávanie)

**Formát:** LinkedIn dokument (PDF carousel — 6 slajdov)

---

**Názov:** 5 vecí, ktoré by mal vedieť každý CTO o AI v roku 2026

**Slajd 1:** "AI nie je chatbot. AI je backend."
**Slajd 2:** "AI Lead Scoring — automatické triedenie dopytov podľa rozpočtu a urgency"
**Slajd 3:** "Voice-to-CRM — klienti nahovoria, AI prepíše, CRM prijme"
**Slajd 4:** "GDPR — AI môže byť compliant, ak hostujete v EÚ"
**Slajd 5:** "Náklady — AI agent na mieru: €2,000. Voice-to-CRM SaaS: €200/mesiac"
**Slajd 6:** "Rezervujte si bezplatnú 30-min konzultáciu: https://ascentia-web-hermes-agent.fly.dev/consultation"

---

## Príspevok 5 (Pondelok — Promo / Nová služba)

**Formát:** Text post

---

Pridali sme na web cenník s transparentnými cenami.

Žiadne "Kontaktujte nás pre cenovú ponuku".

Tier 1: AI Systém — €2,000
• AI agent na mieru + CRM integrácia + lead scoring

Tier 2: Go Backend — €8,000 (najpopulárnejší)
• Go backend + HTMX + Voice-to-CRM + CI/CD

Tier 3: R&D — €15,000+
• Hardvér-softvér integrácia + AI fine-tuning + prototyp

Voice-to-CRM SaaS: €200-500/mesiac

Prečo transparentné ceny?
→ Šetříme čas — vy aj my
→ Budujeme dôveru — žiadne skryté poplatky
→ Filtrujeme leady — iba vážne záujemci

Pozrite si cenník: https://ascentia-web-hermes-agent.fly.dev/pricing

Máte otázky? Napíšte: ascentia@agentmail.to

#Pricing #Transparency #B2B #Go #AI #Slovakia
