# ASCENTIA — Setup sociálnych sietí a Google služieb

> Poradie je dôležité — niektoré účty závisia od iných.

---

## FÁZA 1: Google (priorita 1 — bez toho nefunguje nič)

### 1. Google účet (osobný → firemný)
- **URL:** https://accounts.google.com/signup
- **Email:** `ascentia.sro@gmail.com` (alebo použi existujúci)
- **Čo treba:** Telefón na overenie
- **Čas:** 5 minút

### 2. Google Analytics 4
- **URL:** https://analytics.google.com
- **Kroky:**
  1. Prihlás sa s Google účtom
  2. Admin → Create Account → "ASCENTIA s. r. o."
  3. Create Property → "Ascentia Web"
  4. Data Stream → Web → `https://ascentia-web-hermes-agent.fly.dev`
  5. Skopíruj Measurement ID (format: `G-XXXXXXXXXX`)
  6. Nahraď `G-XXXXXXX` v `templates/layout.html` reálnym ID
- **Čas:** 10 minút

### 3. Google Search Console
- **URL:** https://search.google.com/search-console
- **Kroky:**
  1. Pridaj property → `https://ascentia-web-hermes-agent.fly.dev`
  2. Overenie: HTML tag (pridaj do `<head>` v layout.html)
  3. Submit sitemap: `https://ascentia-web-hermes-agent.fly.dev/static/sitemap.xml`
- **Čas:** 10 minút

### 4. Google Business Profile (firemný profil)
- **URL:** https://www.google.com/business
- **Kroky:**
  1. Pridaj firmu: "ASCENTIA s. r. o."
  2. IČO: 51858959
  3. Adresa: Bratislava
  4. Kategória: "Software Company"
  5. Overenie: Google pošle PIN kód poštou (5-14 dní)
- **Čas:** 15 minút + čakanie na PIN

### 5. YouTube kanál
- **URL:** https://www.youtube.com/create_channel
- **Kroky:**
  1. Prihlás sa s Google účtom
  2. Vytvor kanál: "ASCENTIA s. r. o."
  3. Pridaj logo, banner, popis
- **Čas:** 15 minút

---

## FÁZA 2: LinkedIn (priorita 2 — B2B platforma)

### 6. Osobný LinkedIn profil (Marian Stancik)
- **URL:** https://www.linkedin.com/signup
- **Čo treba:** Foto, pozícia, o firme
- **Headline:** "CEO & Software Architect at ASCENTIA s. r. o. | Go + AI | Voice-to-CRM"
- **Čas:** 20 minút

### 7. LinkedIn Company Page
- **URL:** https://www.linkedin.com/company/setup/new/
- **Kroky:**
  1. Názov: "ASCENTIA s. r. o."
  2. URL: `linkedin.com/company/ascentia-sro`
  3. Priemysel: "Information Technology & Services"
  4. Veľkosť: 1-10 zamestnancov
  5. Typ: "Privately Held"
  6. Logo: ASCENTIA SVG logo
  7. Popis: "Prémiový vývoj high-performance softvéru v Go, autonómne AI systémy a Voice-to-CRM integrácie pre enterprise klientov."
  8. Web: https://ascentia-web-hermes-agent.fly.dev
- **Čas:** 20 minút

---

## FÁZA 3: Facebook + Instagram (priorita 3)

### 8. Facebook Business Page
- **URL:** https://www.facebook.com/pages/create
- **Kroky:**
  1. Typ: "B2B / Software Company"
  2. Názov: "ASCENTIA s. r. o."
  3. Kategória: "Software Company"
  4. Popis: rovnaký ako LinkedIn
  5. Web: https://ascentia-web-hermes-agent.fly.dev
  6. Profilovka + cover photo
- **Čas:** 15 minút

### 9. Instagram Business account
- **URL:** Otvor Instagram app → Settings → Switch to Business
- **Kroky:**
  1. Pripoj Facebook Page (z kroku 8)
  2. Názov: @ascentia.sro
  3. Kategória: "Software Company"
  4. Bio: "Enterprise AI & Go Backend | Voice-to-CRM | Bratislava 🇸🇰"
  5. Link: https://ascentia-web-hermes-agent.fly.dev
- **Čas:** 10 minút

### 10. Meta Business Suite
- **URL:** https://business.facebook.com
- **Kroky:**
  1. Vytvor Business Manager: "ASCENTIA s. r. o."
  2. Pridaj Facebook Page a Instagram account
  3. Nastav Pixel (pre budúce ads)
- **Čas:** 15 minút

---

## FÁZA 4: X/Twitter (priorita 4)

### 11. X/Twitter account
- **URL:** https://twitter.com/i/flow/signup
- **Kroky:**
  1. Názov: "ASCENTIA"
  2. Handle: @ascentia_sro
  3. Bio: "Enterprise AI & High-Performance Go Backend | Voice-to-CRM +340% konverzia | Bratislava 🇸🇰"
  4. Web: https://ascentia-web-hermes-agent.fly.dev
  5. Profilovka: ASCENTIA logo
- **Čas:** 10 minút

---

## Zhrnutie — čo treba urobiť

| # | Platforma | URL | Priorita | Čas |
|---|-----------|-----|----------|-----|
| 1 | Google účet | accounts.google.com | 🔴 Kritické | 5 min |
| 2 | Google Analytics 4 | analytics.google.com | 🔴 Kritické | 10 min |
| 3 | Google Search Console | search.google.com/search-console | 🔴 Kritické | 10 min |
| 4 | Google Business Profile | google.com/business | 🟡 Dôležité | 15 min + PIN |
| 5 | YouTube kanál | youtube.com/create_channel | 🟢 Neskôr | 15 min |
| 6 | LinkedIn osobný | linkedin.com/signup | 🔴 Kritické | 20 min |
| 7 | LinkedIn Company Page | linkedin.com/company/setup/new | 🔴 Kritické | 20 min |
| 8 | Facebook Page | facebook.com/pages/create | 🟡 Dôležité | 15 min |
| 9 | Instagram Business | Instagram app | 🟡 Dôležité | 10 min |
| 10 | Meta Business Suite | business.facebook.com | 🟢 Neskôr | 15 min |
| 11 | X/Twitter | twitter.com | 🟢 Neskôr | 10 min |

**Total čas:** ~2.5 hodiny
**Kritické (urob ako prvé):** Google účet + Analytics + Search Console + LinkedIn

---

## Po založení účtov — čo urobím ja:

1. **Nahradím `G-XXXXXXX`** v layout.html reálnym GA4 ID
2. **Pridám Search Console overovací tag** do layout.html
3. **Nastavím Meta Pixel** ak budete robiť Facebook/Instagram ads
4. **Aktualizujem OG meta tagy** s správnymi social profile linkami
5. **Pridám social linky do footer** (LinkedIn, Facebook, Instagram, Twitter, YouTube)
6. **Nastavím LinkedIn Company Page** s obsahom (popis, logo, cover)
7. **Publikujem prvých 5 LinkedIn postov** (z docs/linkedin-posts.md)
