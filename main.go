package main

import (
	"log"
	"msa-app/internal/api"
	"msa-app/pkg/handler"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env_err := godotenv.Load(".env")
	handler.CheckErrorAndPanic(env_err, ".env Load fail")

	if len(os.Args) < 2 {
		log.Panic("No argument")
	}

	port := os.Args[1]

	api.InitRouter(port)
	log.Println("Server started on: server")
}
