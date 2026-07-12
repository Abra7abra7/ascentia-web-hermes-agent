# ASCENTIA s. r. o. — Brand Manual

> Generované s použitím brandkit skill (leonxlnx/taste-skill)

---

## 1. Brand Strategy

| Parameter | Value |
|-----------|-------|
| **Category** | Enterprise AI & High-Performance Go Backend Development |
| **Audience** | CTO/CEO of B2B technology companies (20-200 employees), Slovakia/Czechia/EU |
| **Personality** | Precise, authoritative, technical, premium, direct, no-fluff |
| **Core Metaphor** | ASCENT — upward momentum, architectural ascent, climbing to technical excellence |
| **Emotional Promise** | "Transformujeme kód v bleskové riešenia." |
| **Cultural Position** | Slovak engineering excellence, EU-based, GDPR-first |
| **Trust Level** | High — enterprise grade, transparent pricing, TDD methodology |
| **What to Avoid** | Buzzwords, generic AI imagery, corporate fluff, stock photos, purple-blue gradients |

---

## 2. Logo System

### Logo Concept: "A" + Upward Chevron

**Method:** Monogram + Meaning (Method 1) + Negative Space (Method 4)

Písmeno "A" z ASCENTIA je kombinované s:
- **Upward chevron** — symbolizuje stúpanie, ascent, momentum
- **Negative space** — medzi nohami "A" je vytvorená šípka hore
- **Geometric construction** — postavené na základe kruhu a diagonálnych rezov

### Logo Variants

| Variant | Usage | Background |
|---------|-------|------------|
| **Full** (logo + wordmark) | Header, footer, email | Deep Blue |
| **Symbol only** (A mark) | Favicon, app icon, avatar | Deep Blue / Cyan |
| **Wordmark only** (ASCENTIA text) | Documents, signatures | Neural Grey |
| **Reversed** (white on cyan) | CTAs, badges, stamps | Cyber Cyan |
| **Monochrome** (single color) | Embossing, stamps | Neural Grey |

### Logo Construction

```
    ◆         ← apex point (chevron tip)
   / \
  /   \       ← negative space creates upward arrow
 /  ▲  \      ← inner triangle = ascent path
/___|___\
```

**Grid:** 100×100 viewBox, golden ratio proportions
**Stroke:** 4px (primary), 3px (secondary)
**Fill:** rgba(0, 240, 255, 0.1) — subtle cyan glow

### Logo Clear Space

Minimum clear space = výška písmena "A" vo wordmarku na všetkých stranách.

### Logo Minimum Size

| Application | Min size |
|-------------|----------|
| Print | 25mm šírka |
| Web (desktop) | 120px šírka |
| Web (mobile) | 80px šírka |
| Favicon | 16×16px |
| Social avatar | 400×400px |

---

## 3. Color System

### Primary Palette

| Token | Hex | RGB | Usage |
|-------|-----|-----|-------|
| **Ascent Deep Blue** | `#0A192F` | 10, 25, 47 | Dominantné pozadie, hlavná farba brandu |
| **Cyber Cyan** | `#00F0FF` | 0, 240, 255 | Akcenty, CTA, linky, glow efekty |
| **Neural Grey** | `#F4F6F9` | 244, 246, 249 | Primárny text, svetlé panely |
| **Code Charcoal** | `#1E2229` | 30, 34, 41 | Panely, ohraničenia, kód |

### Secondary Palette

| Token | Hex | Usage |
|-------|-----|-------|
| **Glass White** | `rgba(255,255,255,0.05)` | Glassmorphism panely |
| **Cyan Glow** | `rgba(0,240,255,0.1)` | Subtle akcenty, hovers |
| **Cyan Border** | `rgba(0,240,255,0.3)` | Ohraničenia |
| **Text Muted** | `#8892b0` | Sekundárny text, metadáta |
| **Dark Overlay** | `rgba(10,25,47,0.6)` | Modálne okná, iframes |

### Color Rules

1. **Dominant:** Deep Blue (#0A192F) — 70% plochy
2. **Accent:** Cyber Cyan (#00F0FF) — 15% plochy (iba akcenty, nikdy veľké plochy)
3. **Text:** Neural Grey (#F4F6F9) — 10% plochy
4. **Panels:** Code Charcoal / Glass White — 5% plochy
5. **Nikdy:** Žiadne ďalšie farby. Žiadne purple, green, orange, red (okrem error states)

### Gradient

```
background: linear-gradient(135deg, #0A192F 0%, #0d2238 100%);
accent: linear-gradient(90deg, #00F0FF 0%, rgba(0,240,255,0.2) 100%);
```

---

## 4. Typography

### Primary Font: System UI / Sans-serif

Používame systémové písmo pre maximálny výkon (žiadne externé fonty = rýchlejší web).

| Usage | Font | Size | Weight | Color |
|-------|------|------|--------|-------|
| H1 (Hero) | system-ui, -apple-system, sans-serif | 3.5rem (56px) | 800 | Neural Grey |
| H2 (Section) | system-ui, sans-serif | 2rem (32px) | 700 | Neural Grey |
| H3 (Card) | system-ui, sans-serif | 1.5rem (24px) | 600 | Cyber Cyan |
| Body | system-ui, sans-serif | 1rem (16px) | 400 | Neural Grey |
| Small / Meta | system-ui, sans-serif | 0.875rem (14px) | 400 | Text Muted |
| Code | 'Monaco', 'Courier New', monospace | 0.9rem (14px) | 400 | Cyber Cyan |

### Typography Rules

1. **Line height:** 1.6 pre body, 1.2 pre headings
2. **Letter spacing:** 2px pre H1 (premium feel), 4px pre logo wordmark
3. **Max line length:** 65ch pre čitateľnosť
4. **Nikdy nepoužívať:** Comic Sans, Times New Roman, Papyrus, Comic fonts

---

## 5. Voice & Tone

### Brand Voice

| Attribute | Do | Don't |
|-----------|----|----|
| **Direct** | "Go je 2.6x rýchlejší" | "Naša inovatívna AI riešenia ponúkajú..." |
| **Technical** | "47,000 req/s, 12MB RAM" | "Vysoký výkon a škálovateľnosť" |
| **Data-driven** | "Konverzia stúpla z 2.1% na 9.2%" | "Výrazne zlepšili sme konverziu" |
| **Premium** | Krátke vety, veľa white space | Husté bloky textu |
| **Honest** | "€5/mesiac namiesto €400" | "Nákladovo efektívne riešenie" |

### Taglines

| Type | Tagline |
|------|---------|
| **Primary** | "Transformujeme kód v bleskové riešenia." |
| **Short** | "Build better." |
| **Technical** | "47,000 req/s. €5/mesiac. Žiadny React." |
| **Voice-to-CRM** | "Nahovorte. AI prepíše. CRM prijme." |
| **GDPR** | "Plne GDPR compliant. EÚ hostované." |

### Language

- **Slovensky** pre slovenský trh (web, blog, LinkedIn)
- **English** pre medzinárodný trh (technická dokumentácia, GitHub)
- **Kód a technická terminológia** vždy v angličtine (Go, HTMX, SQLite, Voice-to-CRM)

---

## 6. Visual System

### Glassmorphism

```css
.glass-panel {
    background: rgba(255, 255, 255, 0.05);
    backdrop-filter: blur(10px);
    border: 1px solid rgba(0, 240, 255, 0.1);
    border-radius: 16px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
}
```

### Glow Effects

```css
.glow-text {
    text-shadow: 0 0 20px rgba(0, 240, 255, 0.5);
}
.glow-box {
    box-shadow: 0 0 30px rgba(0, 240, 255, 0.15);
}
```

### Buttons

| Type | Background | Text | Border | Glow |
|------|-----------|------|--------|------|
| **Primary** | `#0A192F` | `#00F0FF` | `1px solid #00F0FF` | hover: 0 0 20px rgba(0,240,255,0.3) |
| **Accent** | `#00F0FF` | `#0A192F` | none | hover: 0 0 30px rgba(0,240,255,0.5) |
| **Ghost** | transparent | `#F4F6F9` | `1px solid rgba(255,255,255,0.1)` | none |

### Grid System

- **Desktop:** CSS Grid, 12 columns, 30px gutter
- **Tablet:** 8 columns, 20px gutter
- **Mobile:** 4 columns, 16px gutter
- **Container:** max-width 1200px, centered

### Spacing Scale

| Token | Value | Usage |
|-------|-------|-------|
| `xs` | 4px | Inline spacing |
| `sm` | 8px | Small gaps |
| `md` | 16px | Default gaps |
| `lg` | 24px | Section gaps |
| `xl` | 32px | Large sections |
| `2xl` | 48px | Page sections |
| `3xl` | 80px | Hero spacing |

---

## 7. Iconography

### Icon Style

- **Style:** Line icons, 2px stroke
- **Color:** Cyber Cyan (#00F0FF) pre default, Neural Grey pre inactive
- **Size:** 24px (default), 16px (inline), 48px (feature)
- **Source:** Vlastné SVG (žiadne externé ikon knižnice)

### Custom Icons (SVG)

| Icon | Symbol | Usage |
|------|--------|-------|
| AI Brain | Hexagon + dot | AI služby |
| Go Gopher | Stylized gopher | Go backend |
| Microphone | Circle + stem | Voice-to-CRM |
| Shield | Shield + check | GDPR |
| Terminal | Rectangle + cursor | Code/TDD |
| Chevron Up | Upward arrow | Ascent logo element |

---

## 8. Applications

### Website

| Element | Spec |
|---------|------|
| Header | Sticky, glassmorphism, logo left, nav right |
| Hero | Full-width, gradient bg, H1 + CTA |
| Cards | Glass panels, 16px radius, cyan hover glow |
| Footer | Deep Blue, 3 columns, social links |
| Cookie bar | Bottom fixed, glassmorphism |

### Email

| Element | Spec |
|---------|------|
| From | `ASCENTIA Web <ascentia@marianstancik.dev>` |
| Subject prefix | `[ASCENTIA]` pre notifikácie |
| Body | Plain text (no HTML templates) |
| Signature | "S pozdravom, Tim ASCENTIA s. r. o." + URL |

### Social Media

| Platform | Avatar | Cover | Bio |
|----------|--------|-------|-----|
| LinkedIn | Logo symbol (400×400) | Dark navy + tagline | "Enterprise AI & Go Backend | Voice-to-CRM | Bratislava" |
| Facebook | Logo symbol (400×400) | Dark navy + tagline | Same as LinkedIn |
| Instagram | Logo symbol (400×400) | N/A | "Enterprise AI & Go Backend | Voice-to-CRM | Bratislava 🇸🇰" |
| X/Twitter | Logo symbol (400×400) | Dark navy + tagline | "Enterprise AI & Go Backend | Voice-to-CRM | Bratislava" |
| YouTube | Logo symbol (800×800) | Dark navy banner | "ASCENTIA s. r. o. — AI & Go Backend" |

### Business Card

```
┌─────────────────────────┐
│                         │
│   ◆ ASCENTIA            │  ← Logo + wordmark, cyan accent
│                         │
│   ─────────────────     │  ← Thin cyan rule
│                         │
│   Marian Stancik        │  ← Name, Neural Grey
│   CEO & Architect       │  ← Title, Text Muted
│                         │
│   ascentia@agentmail.to │  ← Email, cyan
│   +421 900 000 000      │  ← Phone, muted
│                         │
│         ascentia-web-   │  ← URL, small
│       hermes-agent.fly   │
│            .dev          │
└─────────────────────────┘
Background: #0A192F (Deep Blue)
Card: 85×55mm
```

---

## 9. Anti-Generic Rules

### Nikdy nerobiť:

- ❌ Generic startup gradienty (purple → blue)
- ❌ Stock foto ľudí v okuliaroch hľadiacich do obrazovky
- ❌ Buzzwordy: "inovatívny", "next-gen", "disruptívny"
- ❌ Robot ikony (okrem vlastných)
- ❌ Purple-blue AI glow
- ❌ Full-screen video backgrounds
- ❌ Generic lightning bolts
- ❌ Random emoji v professional kontexte
- ❌ Korporátne PowerPoint slide designy
- ❌ Full-page hero images s textom cez ne

### Vždy robiť:

- ✅ Dark navy pozadie s cyan akcentmi
- ✅ Veľa negative space
- ✅ Tvrdé dáta a čísla
- ✅ Krátke, rezané vety
- ✅ Monospace pre kód a technické detaily
- ✅ Glassmorphism pre panely
- ✅ Vlastné SVG ikony
- ✅ Minimalistický prístup
