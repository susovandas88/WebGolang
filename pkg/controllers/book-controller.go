package controllers

import (
	"encoding/json"
	"firstwebproject/pkg/models"
	"firstwebproject/pkg/utils"
	"net/http"
)

var NewBook models.Book

func GetAllBook(w http.ResponseWriter, r *http.Request) {

	newbooks := models.GetAllBook()
	res, _ := json.Marshal(newbooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
