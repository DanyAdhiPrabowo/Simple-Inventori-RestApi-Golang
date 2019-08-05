package category_api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../../config"
	"../../entities"
	"../../models"
	"github.com/gorilla/mux"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		categoryModel := models.CategoryModel{
			Db: db,
		}
		categories, err2 := categoryModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, categories)
		}
	}
}

func FindSpecific(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		categoryModel := models.CategoryModel{
			Db: db,
		}
		categories, err2 := categoryModel.FindSpecific(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, categories)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		categoryModel := models.CategoryModel{
			Db: db,
		}
		categories, err2 := categoryModel.Search(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, categories)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	var category entities.Category
	err := json.NewDecoder(request.Body).Decode(&category)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		categoryModel := models.CategoryModel{
			Db: db,
		}
		err2 := categoryModel.Create(&category)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, category)
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	var category entities.CategoryEdit
	err := json.NewDecoder(request.Body).Decode(&category)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		categoryModel := models.CategoryModel{
			Db: db,
		}
		_, err2 := categoryModel.Update(id, &category)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, category)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		categoryModel := models.CategoryModel{
			Db: db,
		}
		_, err2 := categoryModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, nil)
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
