package models

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
	Ativo bool `json:"ativo"`
	DataCriacao string `json:"data_criacao"`
	DataAtualizacao string `json:"data_atualizacao"`
}
