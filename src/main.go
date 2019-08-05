package main

import (
	"fmt"
	"log"
	"net/http"

	"./apis/category_api"
	"./apis/product_api"

	"github.com/gorilla/mux"
)

const (
	// ListeningPort is the API listerner port
	ListeningPort = ":6060"
)

func main() {

	router := mux.NewRouter()

	// router.HandleFunc("/api/product/c/{id}", product_api.GetProductByCategory).Methods(http.MethodGet)
	// router.HandleFunc("/api/product?category_id={id_category}", product_api.GetProductByCategory).Methods(http.MethodGet)

	router.HandleFunc("/api/product", product_api.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/api/product/{id}", product_api.FindSpecific).Methods(http.MethodGet)
	router.HandleFunc("/api/product/search/{keyword}", product_api.Search).Methods(http.MethodGet)
	router.HandleFunc("/api/product", product_api.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/product/{id}", product_api.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/product/{id}", product_api.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/category", category_api.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/api/category/{id}", category_api.FindSpecific).Methods(http.MethodGet)
	router.HandleFunc("/api/category/search/{keyword}", category_api.Search).Methods(http.MethodGet)
	router.HandleFunc("/api/category", category_api.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/category/{id}", category_api.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/category/{id}", category_api.Delete).Methods(http.MethodDelete)

	log.Printf("Starting http server at %v", ListeningPort)

	err := http.ListenAndServe(ListeningPort, router)
	if err != nil {
		fmt.Println(err)
	}

}
