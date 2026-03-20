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


## Testes unitários
go test -v ./tests

## Rotas exemplo de uso para Curl:
Cadastro de usuários
```
curl -X POST http://localhost:8080/usuarios/ \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Ana",
    "email": "Ana@email.com",
    "ativo": true
  }'
```