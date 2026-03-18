FROM golang:1.21 AS builder

WORKDIR /app

# Copia arquivos de dependência primeiro (cache)
COPY go.mod go.sum ./

# Baixa dependências
RUN go mod download

# Copia o resto do projeto
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# ---- runtime ----
FROM alpine:latest

WORKDIR /root/

# Certificados SSL (importante!)
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]