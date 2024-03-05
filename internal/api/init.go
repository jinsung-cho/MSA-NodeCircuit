package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var port string

func init() {
	port = "20080"
}

func InitRouter() {
	r := mux.NewRouter()

	corsConfig := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})
	handler := corsConfig.Handler(r)
	r.HandleFunc("/api/end", EndPoint).Methods("POST")
	r.HandleFunc("/api/mid", Middleware).Methods("POST")

	http.ListenAndServe(":"+string(port), handler)
}
