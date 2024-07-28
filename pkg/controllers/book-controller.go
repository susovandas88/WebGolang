package controllers

import (
	"encoding/json"
	"firstwebproject/pkg/models"
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
