package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedramkousari/golang-crud-rest-api/database"
	"github.com/pedramkousari/golang-crud-rest-api/entities"
)

func GetProducts(w http.ResponseWriter, _ *http.Request) {
	var products []entities.Product

	database.Instance.Find(&products)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	var product entities.Product

	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}

	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func checkIfProductExists(productId string) bool {
	var product entities.Product
	database.Instance.First(&product, productId)
	return product.Id != 0
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)

	database.Instance.Create(&product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	product := entities.Product{}
	database.Instance.Delete(&product,productId)

	w.WriteHeader(http.StatusAccepted)
}
