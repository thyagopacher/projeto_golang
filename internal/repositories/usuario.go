package repositories

import (
	"context"
	"errors"

	"projeto_go/internal/database" // importe seu pacote de conexão
	"projeto_go/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetByID - Buscar por ID (recebe string hex do ObjectID)
func (r *UsuarioRepository) GetByID(id string) (*models.Usuario, error) {
	ctx := context.Background()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID inválido (deve ser hex de 24 caracteres)")
	}

	var usuario models.Usuario
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&usuario)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // não encontrado → retorna nil sem erro
		}
		return nil, err
	}

	return &usuario, nil
}

// Create - Criar novo usuário (gera ID automaticamente)
func (r *UsuarioRepository) Create(usuario models.Usuario) (models.Usuario, error) {
	ctx := context.Background()

	// Gera novo ObjectID
	usuario.ID = primitive.NewObjectID()
	usuario.DataCriacao = time.Now()
	usuario.DataAtualizacao = time.Now()
	usuario.Ativo = true // valor padrão comum
	result, err := r.collection.InsertOne(ctx, usuario)
	if err != nil {
		return models.Usuario{}, err
	}

	// Atualiza o ID no struct retornado
	usuario.ID = result.InsertedID.(primitive.ObjectID)

	return usuario, nil
}

// Update - Atualizar usuário existente
func (r *UsuarioRepository) Update(id string, usuario models.Usuario) (*models.Usuario, error) {
	ctx := context.Background()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	// Atualiza apenas os campos permitidos (não mexe no _id)
	update := bson.M{
		"$set": bson.M{
			"nome":  usuario.Nome,
			"email": usuario.Email,
		},
	}

	var updated models.Usuario
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After) // retorna o documento após update

	err = r.collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update, opts).Decode(&updated)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &updated, nil
}

// Delete - Deletar por ID
func (r *UsuarioRepository) Delete(id string) (bool, error) {
	ctx := context.Background()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, errors.New("ID inválido")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, nil
}