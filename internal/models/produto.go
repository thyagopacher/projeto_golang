package models

import (
	"time"
)

type Produto struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	Foto string `json:"foto"`
	Preco float64 `json:"preco"`
	Ativo           bool               `bson:"ativo" json:"ativo"`
	DataCriacao     time.Time          `bson:"data_criacao" json:"data_criacao"`
	DataAtualizacao time.Time          `bson:"data_atualizacao" json:"data_atualizacao"`
}
