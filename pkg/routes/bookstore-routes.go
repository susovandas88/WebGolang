package routes

import (
	"firstwebproject/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.Createbook).Methods("POST")
}
