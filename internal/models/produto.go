package models

type Produto struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	Foto string `json:"foto"`
	Preco float64 `json:"preco"`
	DataCriacao string `json:"data_criacao"`
	DataAtualizacao string `json:"data_atualizacao"`
}
