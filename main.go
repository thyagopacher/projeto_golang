package main 

import ( 
 	"fmt" 
	"net/http"
	"log"
	"os"
	"time"
	"github.com/newrelic/go-agent/v3/newrelic"
    "github.com/gin-gonic/gin"
    "projeto_go/routes"
)

func homeRoute (w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintf(w, "Olá, aplicação Go em Docker!" ) 
} 

func  main () { 
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

	// Gin
	r := gin.New()

	// Middlewares básicos
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

    routes.SetupRoutes(r)

	fmt.Println( "O servidor está rodando na porta 8080" ) 
	r.Run(":8080")
}