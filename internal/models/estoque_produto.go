package models

type EstoqueProduto struct {
	ID   int    `json:"id"`
	IDProduto int `json:"id_produto"`
	Qtde int `json:"qtde"`
	DataCriacao string `json:"data_criacao"`
}
