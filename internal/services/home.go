package services

import (
	"projeto_go/internal/database"
	"time"
)

type HomeService struct {
	// Sem repositório, é um serviço simples
}

func NewHomeService() *HomeService {
	return &HomeService{}
}

/**
* GET /health
 */
func (s *HomeService) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status":    "ok",
		"message":   "API rodando",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "projeto-go",
		"version":   "1.0.0",
	}
}

/**
* GET / (opcional - página inicial)
 */
func (s *HomeService) GetHome() map[string]interface{} {
	isMongoConnected := database.IsMongoConnected()
	isRedisConnected := database.IsRedisConnected()

	return map[string]interface{}{
		"name":           "Projeto Go API",
		"description":    "API RESTful em Go com Gin e MongoDB",
		"version":        "1.0.0",
		"databaseStatus": isMongoConnected,
		"redisStatus":    isRedisConnected,
		"timestamp":      time.Now().Format(time.RFC3339),
	}
}
