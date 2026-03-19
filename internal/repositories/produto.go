package repositories

import (
	"context"
	"time"
	"projeto_go/internal/database" // importe seu pacote de conexão
	"projeto_go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionNameProdutos = "Produtos" // nome da coleção no MongoDB
)

// ProdutoRepository implementa CRUD no MongoDB
type ProdutoRepository struct {
	collection *mongo.Collection
}

// NewProdutoRepository cria o repositório usando a conexão global do database
func NewProdutoRepository() *ProdutoRepository {
	// Usa o DB global que você conectou em database.ConnectMongo
	return &ProdutoRepository{
		collection: database.DB.Collection(collectionNameProdutos),
	}
}

// GetAll - Listar todos os usuários
func (r *ProdutoRepository) GetAll() ([]models.Produto, error) {
	ctx := context.Background() // ou context.TODO() / WithTimeout em prod

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var Produtos []models.Produto
	if err = cursor.All(ctx, &Produtos); err != nil {
		return nil, err
	}

	return Produtos, nil
}

// GetByID - Buscar por ID (recebe int)
func (r *ProdutoRepository) GetByID(id int) (models.Produto, error) {
	ctx := context.Background()

	var Produto models.Produto
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&Produto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Produto{}, nil // não encontrado
		}
		return models.Produto{}, err
	}
	return Produto, nil
}

// Create - Criar novo usuário (gera ID automaticamente)
func (r *ProdutoRepository) Create(Produto models.Produto) (models.Produto, error) {
	ctx := context.Background()

	// Lógica de auto-increment simples (não concorrência-safe – veja abaixo para versão segura)
	var maxDoc bson.M
	err := r.collection.FindOne(
		ctx,
		bson.M{},
		options.FindOne().SetSort(bson.M{"id": -1}),
	).Decode(&maxDoc)
	if err != nil && err != mongo.ErrNoDocuments {
		return models.Produto{}, err
	}

	nextID := 1
	if err == nil {
		nextID = int(maxDoc["id"].(int32)) + 1 // ou int64 se for grande
	}

	Produto.ID = nextID

	Produto.DataCriacao = time.Now()
	Produto.DataAtualizacao = time.Now()
	Produto.Ativo = true
	Produto.Foto = "123" // valor padrão, pode ser modificado depois
	Produto.Preco = 0.0

	_, err = r.collection.InsertOne(ctx, Produto)
	if err != nil {
		return models.Produto{}, err
	}

	return Produto, nil
}

// Update - Atualizar usuário existente
func (r *ProdutoRepository) Update(id int, Produto models.Produto) (models.Produto, error) {
	ctx := context.Background()

	// Atualiza apenas os campos permitidos (não mexe no id)
	update := bson.M{
		"$set": bson.M{
			"nome":             Produto.Nome,
			"ativo":            Produto.Ativo,
			"preco":            Produto.Preco,
			"foto":             Produto.Foto,
			"data_atualizacao": time.Now(), // sempre atualiza a data de modificação
		},
	}

	var updated models.Produto
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After) // retorna o documento após update

	err := r.collection.FindOneAndUpdate(ctx, bson.M{"id": id}, update, opts).Decode(&updated)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Produto{}, nil
		}
		return models.Produto{}, err
	}

	return updated, nil
}

// Delete - Deletar por ID
func (r *ProdutoRepository) Delete(id int) (bool, error) {
	ctx := context.Background()

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, nil
}