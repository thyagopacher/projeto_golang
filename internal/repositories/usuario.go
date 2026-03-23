package repositories

import (
	"context"
	"projeto_go/internal/database" // importe seu pacote de conexão
	"projeto_go/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "usuarios" // nome da coleção no MongoDB
)

// UsuarioRepository implementa CRUD no MongoDB
type UsuarioRepository struct {
	collection *mongo.Collection
}

// NewUsuarioRepository cria o repositório usando a conexão global do database
func NewUsuarioRepository() *UsuarioRepository {
	// Usa o DB global que você conectou em database.ConnectMongo
	return &UsuarioRepository{
		collection: database.DB.Collection(collectionName),
	}
}

// GetAll - Listar todos os usuários
func (r *UsuarioRepository) GetAll() ([]models.Usuario, error) {
	ctx := context.Background() // ou context.TODO() / WithTimeout em prod

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var usuarios []models.Usuario
	if err = cursor.All(ctx, &usuarios); err != nil {
		return nil, err
	}

	return usuarios, nil
}

// GetByID - Buscar por ID (recebe int)
func (r *UsuarioRepository) GetByID(id int) (models.Usuario, error) {
	ctx := context.Background()

	var usuario models.Usuario
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&usuario)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Usuario{}, nil // não encontrado
		}
		return models.Usuario{}, err
	}
	return usuario, nil
}

// Create - Criar novo usuário (gera ID automaticamente)
func (r *UsuarioRepository) Create(usuario models.Usuario) (models.Usuario, error) {
	ctx := context.Background()

	// Lógica de auto-increment simples (não concorrência-safe – veja abaixo para versão segura)
	var maxDoc bson.M
	err := r.collection.FindOne(
		ctx,
		bson.M{},
		options.FindOne().SetSort(bson.M{"id": -1}),
	).Decode(&maxDoc)
	if err != nil && err != mongo.ErrNoDocuments {
		return models.Usuario{}, err
	}

	nextID := 1
	if err == nil {
		nextID = int(maxDoc["id"].(int32)) + 1 // ou int64 se for grande
	}

	usuario.ID = nextID
	now := time.Now()
	usuario.DataCriacao = now
	usuario.DataAtualizacao = now
	usuario.Ativo = true

	_, err = r.collection.InsertOne(ctx, usuario)
	if err != nil {
		return models.Usuario{}, err
	}

	return usuario, nil
}

// Update - Atualizar usuário existente
func (r *UsuarioRepository) Update(id int, usuario models.Usuario) (models.Usuario, error) {
	ctx := context.Background()

	// Atualiza apenas os campos permitidos (não mexe no id)
	update := bson.M{
		"$set": bson.M{
			"nome":             usuario.Nome,
			"email":            usuario.Email,
			"ativo":            usuario.Ativo,
			"data_atualizacao": time.Now(), // sempre atualiza a data de modificação
		},
	}

	var updated models.Usuario
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After) // retorna o documento após update

	err := r.collection.FindOneAndUpdate(ctx, bson.M{"id": id}, update, opts).Decode(&updated)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Usuario{}, nil
		}
		return models.Usuario{}, err
	}

	return updated, nil
}

// Delete - Deletar por ID
func (r *UsuarioRepository) Delete(id int) (bool, error) {
	ctx := context.Background()

	result, err := r.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, nil
}
