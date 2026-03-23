## BUILD
```
docker compose up -d 
```

## Para atualizar go.sum
```
docker run --rm -v $(pwd):/app -w /app golang:1.21 go mod tidy;
```

## Rodar 
-URL para servir o projeto:
http://localhost:8080/


## Testes unitários
go test -v ./tests

## Rotas exemplo de uso para Curl:

1 - Cadastro de usuários
```
curl -X POST http://localhost:8080/usuarios/ \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Ana",
    "email": "Ana@email.com",
    "ativo": true
  }'
```

2 - Gerando de auth JWT
```
curl -X POST 'http://localhost:8080/api/auth/login' \
  --header 'Content-Type: application/json' \
  --body '{"user":"test","pass":"known"}'
```