package main

import (
	"fmt"
	"net/http"

	"./apis/product_api"
	// "./apis/category_api"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/product", product_api.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/api/product/{id}", product_api.FindSpecific).Methods(http.MethodGet)
	router.HandleFunc("/api/product/search/{keyword}", product_api.Search).Methods(http.MethodGet)
	router.HandleFunc("/api/product", product_api.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/product/{id}", product_api.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/product/{id}", product_api.Delete).Methods(http.MethodDelete)

	err := http.ListenAndServe(":6060", router)
	if err != nil {
		fmt.Println(err)
	}

}
