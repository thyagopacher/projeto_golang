package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"projeto_go/internal/controllers"
	"projeto_go/internal/database"
	"projeto_go/internal/repositories"
	"projeto_go/internal/services"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

type Produto struct {
	ID              int       `json:"id"`
	Nome            string    `json:"nome"`
	Foto            string    `json:"foto"`
	Preco           float64   `json:"preco"`
	Ativo           bool      `json:"ativo"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataAtualizacao time.Time `json:"data_atualizacao"`
}

func TestCreateProduto(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	database.ConnectMongo("mongodb://root:root123@localhost:27017", "testdb") // Conecte ao MongoDB de teste
	defer database.Disconnect()

	// registra só a rota que quer testar
	ProdutoRepository := repositories.NewProdutoRepository()              // passe nil ou um mock do banco
	ProdutoService := services.NewProdutoService(ProdutoRepository)       // passe nil ou um mock do repository
	ProdutoController := controllers.NewProdutoController(ProdutoService) // passe nil ou um mock do service
	router.POST("/Produtos", ProdutoController.CreateProduto)

	produto := Produto{
		Nome:  "Produto Teste",
		Foto:  "foto.png",
		Preco: 99.90,
		Ativo: true,
	}

	jsonValue, _ := json.Marshal(produto)

	req, _ := http.NewRequest(http.MethodPost, "/Produtos", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated && w.Code != http.StatusOK {
		t.Errorf("Erro: %d", w.Code)
	}
}
