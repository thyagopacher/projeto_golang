package main 

import ( 
 	"fmt" 
	"net/http"
	"log"
	"os"
	"time"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func  handler (w http.ResponseWriter, r *http.Request) { 
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

	http.HandleFunc( "/" , handler) 
	fmt.Println( "O servidor está rodando na porta 8080" ) 
	http.ListenAndServe( ":8080" , nil ) 
}