docker run --rm -v $(pwd):/app -w /app golang:1.21 go mod tidy
docker compose down && docker compose up -d