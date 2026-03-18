## BUILD
```
docker compose up -d 
```

## Para atualizar go.sum
```
docker run --rm -v $(pwd):/app -w /app golang:1.21 go mod tidy;
```

## Rodar 
http://localhost:8080/

