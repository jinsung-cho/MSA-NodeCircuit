package main

import (
	"log"
	"msa-app/internal/api"
	"msa-app/pkg/handler"

	"github.com/joho/godotenv"
)

func main() {
	env_err := godotenv.Load(".env")
	handler.CheckErrorAndPanic(env_err, ".env Load fail")

	api.InitRouter()

	log.Println("Server started on: server")
}
