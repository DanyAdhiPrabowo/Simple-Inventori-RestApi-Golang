package main

import (
	"fmt"
	"log"
	"net/http"

	"./apis/categoryapi"
	"./apis/productapi"

	"github.com/gorilla/mux"
)

const (
	// ListeningPort is the API listerner port
	ListeningPort = ":6060"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/product", productapi.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/api/product/{id}", productapi.FindSpecific).Methods(http.MethodGet)
	router.HandleFunc("/api/product/search/{keyword}", productapi.Search).Methods(http.MethodGet)
	router.HandleFunc("/api/product", productapi.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/product/{id}", productapi.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/product/{id}", productapi.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/category", categoryapi.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/api/category/{id}", categoryapi.FindSpecific).Methods(http.MethodGet)
	router.HandleFunc("/api/category/search/{keyword}", categoryapi.Search).Methods(http.MethodGet)
	router.HandleFunc("/api/category", categoryapi.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/category/{id}", categoryapi.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/category/{id}", categoryapi.Delete).Methods(http.MethodDelete)

	log.Printf("Starting http server at %v", ListeningPort)

	err := http.ListenAndServe(ListeningPort, router)
	if err != nil {
		fmt.Println(err)
	}

}
