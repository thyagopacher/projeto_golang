package models

import (
	"time"
)

type Usuario struct {
	ID              int `bson:"id" json:"id"`
	Nome            string             `bson:"nome" json:"nome"`
	Email           string             `bson:"email" json:"email"`
	Ativo           bool               `bson:"ativo" json:"ativo"`
	DataCriacao     time.Time          `bson:"data_criacao" json:"data_criacao"`
	DataAtualizacao time.Time          `bson:"data_atualizacao" json:"data_atualizacao"`
}