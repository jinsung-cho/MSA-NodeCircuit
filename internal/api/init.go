package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"msa-app/pkg/handler"
)

var myPort string

func InitRouter(port string) {
	r := mux.NewRouter()
	myPort = port
	corsConfig := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})
	corsHandler := corsConfig.Handler(r)
	r.HandleFunc("/api/end", EndPoint).Methods("POST")
	r.HandleFunc("/api/mid", Middleware).Methods("POST")

	err := http.ListenAndServe(":"+string(port), corsHandler)
	handler.CheckErrorAndPanic(err, "Listen Server Err")
}
