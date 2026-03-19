package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nome            string             `bson:"nome" json:"nome"`
	Email           string             `bson:"email" json:"email"`
	Ativo           bool               `bson:"ativo" json:"ativo"`
	DataCriacao     time.Time          `bson:"data_criacao" json:"data_criacao"`
	DataAtualizacao time.Time          `bson:"data_atualizacao" json:"data_atualizacao"`
}