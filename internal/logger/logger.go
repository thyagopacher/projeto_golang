package logger

import (
	"io"
	"log"
	"os"
)

func SetupLog() {
	// cria pasta logs se não existir
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo de log:", err)
	}

	// escreve no arquivo E no console
	multi := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multi)
}