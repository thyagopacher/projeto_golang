package main

import (
	"log"
	"os"
	"projeto_go/internal/controllers"
	"projeto_go/internal/database"
	"projeto_go/internal/logger"
	"projeto_go/internal/repositories"
	"projeto_go/internal/services"
	"projeto_go/routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {

	logger.SetupLog()

	gin.DefaultWriter = log.Writer()

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEW_RELIC_APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
	)
	if err != nil {
		log.Fatal("Failed to create New Relic application:", err)
	}

	// Wait for the application to connect
	if err := app.WaitForConnection(5 * time.Second); err != nil {
		log.Println("Warning: New Relic application did not connect:", err)
	}

	log.Println("Iniciando aplicação...")

	// Conecta ao MongoDB (chame antes de iniciar o servidor)
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	if uri == "" || dbName == "" {
		log.Fatal("MONGO_URI e MONGO_DB_NAME devem estar definidos nas variáveis de ambiente")
	}
	if err := database.ConnectMongo(uri, dbName); err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}
	defer database.Disconnect() // garante desconexão no shutdown

	// Conecta ao Redis (chame antes de iniciar o servidor)
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	if redisHost == "" || redisPort == "" {
		log.Fatal("REDIS_HOST e REDIS_PORT devem estar definidos nas variáveis de ambiente")
	}
	if err := database.ConnectRedis(redisHost, redisPort); err != nil {
		log.Fatalf("Erro ao conectar ao Redis: %v", err)
	}
	defer database.DisconnectRedis() // garante desconexão no shutdown

	// Gin
	r := gin.New()

	// Middlewares básicos
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	/**
	* Middleware do New Relic (ESSENCIAL)
	* captura error automaticamente e monitora transações
	 */
	r.Use(nrgin.Middleware(app))

	// Home Service (sem repo)
	homeService := services.NewHomeService()
	homeController := controllers.NewHomeController(homeService)

	usuarioRepo := repositories.NewUsuarioRepository()
	usuarioService := services.NewUsuarioService(usuarioRepo)
	usuarioController := controllers.NewUsuarioController(usuarioService)

	// Produto Service
	produtoRepo := repositories.NewProdutoRepository()
	produtoService := services.NewProdutoService(produtoRepo)
	produtoController := controllers.NewProdutoController(produtoService)

	// Fatura Service (sem repo)
	faturaService := services.NewFaturaService()
	faturaController := controllers.NewFaturaController(faturaService)

	routes.SetupRoutes(r, &routes.Controllers{
		Usuario: usuarioController,
		Fatura:  faturaController,
		Produto: produtoController,
		Home:    homeController,
	})

	log.Println("O servidor está rodando na porta 8080")
	r.Run(":8080")
}
