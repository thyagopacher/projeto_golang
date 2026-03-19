package models

type Fatura struct {
	ID   int    `json:"id"`
	IdUsuario int `json:"id_usuario"`
	IdCliente int `json:"id_cliente"`
	DataCriacao string `json:"data_criacao"`
	DataAtualizacao string `json:"data_atualizacao"`
}
