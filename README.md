## 📬 Postman Collection

A collection do Postman está disponível em:

[Download da Collection](docs/postman_collection.json)

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

## ERRO para autenticação em rotas que usam JWT
-É necessário gerar o token antes no endpoint `/api/auth/login` 
```json
{
    "error": "Token não informado"
}
```

## Rotas exemplo de uso para Curl:

1 - Cadastro de usuários

```shell
curl -X POST http://localhost:8080/api/usuarios/ \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Ana",
    "email": "Ana@email.com",
    "ativo": true
  }'
```
-retorno OK:
```json
[
    {
        "id": 1,
        "nome": "Thyago",
        "email": "thyago@email.com",
        "ativo": true,
        "data_criacao": "2026-03-20T10:19:20.498Z",
        "data_atualizacao": "2026-03-20T10:19:20.498Z"
    },
    {
        "id": 2,
        "nome": "Ana",
        "email": "Ana@email.com",
        "ativo": true,
        "data_criacao": "2026-03-20T10:19:38.954Z",
        "data_atualizacao": "2026-03-20T10:19:38.954Z"
    }
]
```
--------------------------------------------------

2 - Gerando de auth JWT
```
curl -X POST 'http://localhost:8080/api/auth/login' \
  --header 'Content-Type: application/json' \
  --body '{"user":"test","pass":"known"}'
```