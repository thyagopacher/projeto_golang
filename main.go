package main 
import ( 
 "fmt" 
 "net/http"
 ) 
func  handler (w http.ResponseWriter, r *http.Request) { 
fmt.Fprintf(w, "Olá, aplicação Go em Docker!" ) 
} 
func  main () { 
http.HandleFunc( "/" , handler) 
fmt.Println( "O servidor está rodando na porta 8080" ) 
http.ListenAndServe( ":8080" , nil ) 
}