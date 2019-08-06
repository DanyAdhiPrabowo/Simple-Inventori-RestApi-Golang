package productapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../../config"
	"../../entities"
	"../../models"
	"github.com/gorilla/mux"
)

// GetProduct menampilkan semua product atau menampilkan product berdasarkan id_category
func GetProduct(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id_category")
	idCategory, _ := strconv.ParseInt(id, 10, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.GetProduct(idCategory)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, products)
		}
	}
}

// FindSpecific menampilkan product secara spesifict berdasarkan id_product
func FindSpecific(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindSpecific(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, products)
		}
	}
}

// Search untuk mencari product berdasarkan nama product yang di input
func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Search(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, products)
		}
	}
}

// Create menambah Product baru
func Create(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		err2 := productModel.Create(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, product)
		}
	}
}

// Update = mengupdate secara spesificasi product berdasarkan id_product
func Update(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	var product entities.ProductEdit
	err := json.NewDecoder(request.Body).Decode(&product)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		_, err2 := productModel.Update(id, &product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, product)
		}
	}
}

// Delete = menghapus secara spesifikasi product berdasarkan id_product
func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		_, err2 := productModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, nil)
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

// respondWithJSON merubah response atau return berupa json
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
