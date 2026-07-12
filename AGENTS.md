# AGENTS.md - Kontext projektu Ascentia, architektúra, B2B monetizácia a AI-GEO Engine

Tento dokument slúži ako hlavný prehľad systému, architektonický plán a komerčná stratégia pre AI agentov a vývojárov pracujúcich na ekosystéme Ascentia.

---

## ✦ Prehľad projektu a biznisová misia

Ascentia nie je len korporátna webstránka; je to vysoko konverzný, prémiový B2B nástroj na generovanie leadov (dopytov) a portál na demonštráciu podnikových AI riešení pre spoločnosť ASCENTIA s. r. o.

Našou hlavnou biznisovou misiou je transformovať špičkové technológie (Go, HTMX, multimodálna AI) do B2B konzultácií s vysokou maržou, vývoja softvéru na mieru a škálovateľných AI riešení pre korporátnu klientelu.

### 💰 Monetizačné piliere a hodnota pre klienta
1. Predplatné AI automatizácií (SaaS-lite): Predaj predpripravených AI agentov (inšpirovaných sekciami /kompas a `/voice-inquiry`) prispôsobených pre slovenských a európskych firemných klientov.
2. High-Ticket vývoj na mieru a architektúra: Prémiové softvérové inžinierstvo v jazyku Go (Golang) pre enterprise klientov, ktorí vyžadujú ultra rýchle backendové systémy s nulovými externými závislosťami.
3. Integrácia hlasu do CRM (Voice-to-CRM): Licencovanie technológie stojacej za naším hlasovým dopytom pre právnické, lekárske alebo realitné kancelárie na automatizovaný príjem a analýzu klientov.

---

## ✦ Štruktúra webu a konverzný lievik

Webstránka je štruktúrovaná ako prísny marketingový a konverzný lievik:

1. O nás / Domov (`/`): Budovanie dôvery a autority. Zobrazuje údaje z ORSR pre maximálnu kredibilitu pred korporátnymi partnermi, strategické hodnoty a primárny B2B formulár na zachytenie dopytov.
2. Služby (`/services`): Monetizovaná ponuka. Jasné rozdelenie do troch prémiových úrovní (Tierov):
   - *Tier 1:* Autonómne AI Systémy (AI agenti na mieru a ladenie/fine-tuning LLM modelov).
   - *Tier 2:* Softvérová Architektúra (Enterprise inžinierstvo v Go, vysokorýchlostné systémy).
   - *Tier 3:* Výskum a Vývoj (R&D, rýchle prototypovanie pre integráciu hardvéru a softvéru).
3. Ako fungujeme (`/process`): Znižovanie nákupného rizika. 4-kroková B2B časová os (Analýza -> Návrh -> Vývoj -> Nasadenie) navrhnutá tak, aby vybudovala dôveru u konzervatívnych firemných nákupcov.
4. AI Copilot Showcase (`/kompas`): Interaktívna návnada (Product-Led Growth). Živé demo AI agenta so streamovaním textu. Klienti zažijú našu technologickú úroveň v praxi ešte pred tým, ako si dohodnú úvodný hovor.
5. Hlasový dopyt (`/voice-inquiry`): Odstraňovanie bariér. Prémiový 30-sekundový formulár na nahrávanie hlasu. Rieši problém vyťažených manažérov, ktorí nechcú písať dlhé texty. AI na pozadí správu automaticky prepíše a kategorizuje.
6. Ochrana osobných údajov (`/privacy`): Právny súlad. Kompletná dokumentácia v súlade s GDPR v slovenskom jazyku, nevyhnutná pre spracovanie dát podnikových klientov.

---

## 🛠️ Technologický stack ad nákladová efektivita (Vysoká marža)

Architektúra je optimalizovaná na maximálny výkon a minimálne prevádzkové náklady (základná infraštruktúra za 0 €/mesiac), čo umožňuje dosahovať bezkonkurenčné marže.

- Go 1.26+: Čistý backend využívajúci štandardný multiplexer (`net/http`). Nasadenie formou jedného binárneho súboru, ultra nízka spotreba RAM, schopnosť obslúžiť milióny požiadaviek na lacnom $5 VPS.
- HTMX 2.x: Dynamické interakcie na fronte. Eliminuje potrebu prekomplikovaných frameworkov (React/Node.js), čím dramaticky znižuje čas vývoja a náklady na údržbu.
- SQLite: Štruktúrovaná databáza využívajúca driver bez CGO (`modernc.org/sqlite`). Bezúdržbová databáza, ktorá zjednodušuje zálohovanie a replikáciu.
- Vanilla CSS & SVGs: Žiadny Tailwind alebo ťažký JavaScript. Čistý dizajn postavený na CSS premenných zaručuje 100% skóre v Google Lighthouse, čo maximalizuje organické SEO.

---

## 🎨 Dizajnový systém: Prémiový B2B vizuál pre rok 2026
Vizuálna identita je navrhnutá tak, aby pôsobila futuristicky, responzívne a technologicky dominantne, čím obhajuje vysoké ceny našich služieb a oslovuje trh v roku 2026.

### Farebná paleta:
- Dominantné pozadie: Ascent Deep Blue (`#0A192F`)
- Akcentová farba: Cyber Cyan (`#00F0FF`)
- Sekundárne farby: Neural Grey (`#F4F6F9`), Code Charcoal (`#1E2229`)

---

## 🤖 AI-Agent Optimization (GEO & SEO 2026)

Web je plne optimalizovaný pre modernú éru vyhľadávania, kde informácie konzumujú a spracovávajú predovšetkým LLM crawlerov.

1. Kompletné sémantické JSON-LD štruktúrovanie: Každá podstránka obsahuje mikrodáta typu Organization, ProfessionalService a Product.
2. LLM-Friendly textová vrstva.
3. Zákaz blokovania AI botov (Robots.txt).
4. LLM Citation Hooks.

---

## 🧠 Pluggable AI Engine s vyhodnocovaním leadov (Kompas & Hlas)

Backendový AI modul (`/ai`) je navrhnutý tak, aby abstrahoval poskytovateľov LLM, čo nám umožňuje flexibilne optimalizovať náklady na API podla rozpočtu klienta.

- Skórovanie a smerovanie leadov.
- Univerzálne rozhranie (`ai.Provider`).
- Streamovanie v reálnom čase (SSE).
