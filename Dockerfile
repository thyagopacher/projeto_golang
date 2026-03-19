# ---- build ----
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# ---- runtime ----
FROM debian:bookworm-slim

WORKDIR /root/

# Instala wkhtmltopdf + dependências
RUN apt-get update && apt-get install -y \
    wkhtmltopdf \
    ca-certificates \
    fonts-dejavu \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]