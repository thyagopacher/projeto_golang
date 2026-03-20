package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"projeto_go/internal/controllers"
	"projeto_go/internal/services"
	"projeto_go/internal/repositories"
	"projeto_go/internal/database"
	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID              int       `json:"id"`
	Nome            string    `json:"nome"`
	Email           string    `json:"email"`
	Ativo           bool      `json:"ativo"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataAtualizacao time.Time `json:"data_atualizacao"`
}

func TestCreateUsuario(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	database.ConnectMongo("mongodb://root:root123@localhost:27017", "testdb") // Conecte ao MongoDB de teste
	defer database.Disconnect()

	// registra só a rota que quer testar
	usuarioRepository := repositories.NewUsuarioRepository() // passe nil ou um mock do banco
	usuarioService := services.NewUsuarioService(usuarioRepository) // passe nil ou um mock do repository
	usuarioController := controllers.NewUsuarioController(usuarioService) // passe nil ou um mock do service
	router.POST("/usuarios", usuarioController.CreateUsuario)

	usuario := Usuario{
		Nome:  "Thyago",
		Email: "thyago@email.com",
		Ativo: true,
	}

	jsonValue, _ := json.Marshal(usuario)

	req, _ := http.NewRequest(http.MethodPost, "/usuarios", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated && w.Code != http.StatusOK {
		t.Errorf("Erro: %d", w.Code)
	}
}