package main 

import ( 
 	"fmt" 
	"net/http"
	"log"
	"os"
	"time"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
    "github.com/gin-gonic/gin"
    "projeto_go/routes"
	"projeto_go/internal/logger"
	"projeto_go/internal/controllers"
	"projeto_go/internal/repositories"
	"projeto_go/internal/services"
	"projeto_go/internal/database"
)

func homeRoute (w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintf(w, "Olá, aplicação Go em Docker!" ) 
} 

func  main () { 

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
	
	// Gin
	r := gin.New()

	// Middlewares básicos
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Middleware do New Relic (ESSENCIAL)
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

	log.Println( "O servidor está rodando na porta 8080" ) 
	r.Run(":8080")
}