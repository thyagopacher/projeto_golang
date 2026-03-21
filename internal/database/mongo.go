package database

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB     *mongo.Database
	client *mongo.Client

	once      sync.Once
	initErr   error
	connected bool
)

// ConnectMongo estabelece conexão com o MongoDB (singleton - chama apenas uma vez)
// uri: conexão (ex: "mongodb://localhost:27017" ou MongoDB Atlas URI)
// dbName: nome do banco de dados
func ConnectMongo(uri, dbName string) error {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Opções recomendadas para v1.17.9
		clientOpts := options.Client().
			ApplyURI(uri).
			SetConnectTimeout(10 * time.Second).
			SetServerSelectionTimeout(10 * time.Second).
			SetSocketTimeout(30 * time.Second) // útil para queries longas

		var err error
		client, err = mongo.Connect(ctx, clientOpts)
		if err != nil {
			initErr = err
			log.Printf("❌ Erro ao criar client MongoDB: %v", err)
			return
		}

		// Ping para validar conexão
		if err = client.Ping(ctx, nil); err != nil {
			initErr = err
			log.Printf("❌ Erro ao pingar MongoDB: %v", err)
			return
		}

		DB = client.Database(dbName)
		connected = true

		log.Printf("✅ Conectado ao MongoDB | Database: %s | URI: %s", dbName, uri)
	})

	if initErr != nil {
		return initErr
	}

	if !connected {
		return mongo.ErrClientDisconnected
	}

	return nil
}

// GetClient retorna o client MongoDB (útil para desconexão ou uso avançado)
func GetClient() *mongo.Client {
	return client
}

// Disconnect fecha a conexão com o MongoDB (chame no shutdown da aplicação)
func Disconnect() {
	if client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Printf("⚠️ Erro ao desconectar do MongoDB: %v", err)
		return
	}

	log.Println("🛑 Conexão com MongoDB fechada com sucesso")
	connected = false
	client = nil
	DB = nil

	once = sync.Once{} // 🔥 reset
}

func IsMongoConnected() bool {
	if client == nil {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := client.Ping(ctx, nil)
	return err == nil
}