package models

type ItemFatura struct {
	ID   int    `json:"id"`
	IdFatura int `json:"id_fatura"`
	IdProduto int `json:"id_produto"`
	Quantidade int `json:"quantidade"`
	DataCriacao string `json:"data_criacao"`
	DataAtualizacao string `json:"data_atualizacao"`
}
