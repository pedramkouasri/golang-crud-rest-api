package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedramkousari/golang-crud-rest-api/controllers"
	"github.com/pedramkousari/golang-crud-rest-api/database"
)

func main() {
	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)
	RegisterProductRoutes(router)

	log.Printf("serving in address localhsot:%v", AppConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte(`{"message": "pong"}`))
	})

	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")

	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")

	router.HandleFunc("/api/product/{id}", controllers.GetProductById).Methods("GET")

	router.HandleFunc("/api/product/{id}", controllers.UpdateProduct).Methods("PUT")

	router.HandleFunc("/api/product/{id}", controllers.DeleteProduct).Methods("DELETE")
}
