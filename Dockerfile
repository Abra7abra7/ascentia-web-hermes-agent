# Fáza 1: Build aplikácie v Go
FROM golang:1.22-alpine AS builder

# Inštalácia potrebných build nástrojov
RUN apk add --no-cache git build-base

WORKDIR /app

# Kopírovanie závislostí a ich stiahnutie
COPY go.mod go.sum ./
RUN go mod download

# Kopírovanie zdrojových kódov aplikácie
COPY . .

# Skompilovanie super-rýchlej produkčnej binárky
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ascentia-web cmd/server/main.go

# Fáza 2: Finálny minimalistický produkčný obraz
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata sqlite

WORKDIR /app

# Kopírovanie skompilovanej binárky z build fázy
COPY --from=builder /app/ascentia-web .

# Kopírovanie statických šablón a štýlov
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Fly.io odporúča držať dáta na upevnenom disku kvôli SQLite. 
# Predvolene nastavíme cestu k DB na persistentný adresár
ENV PORT=8080
ENV DB_PATH=/data/ascentia.db
ENV AI_PROVIDER=mock

EXPOSE 8080

CMD ["./ascentia-web"]
