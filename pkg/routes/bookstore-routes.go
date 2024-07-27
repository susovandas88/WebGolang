package routes

import (
	"github.com/gorilla/mux"
	"\\Users\\Susovan\\source\\repos\\golang\\WebGolang\\pkg\\controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router)
{
	router.HandleFunc("/book/",controllers.Createbook).Methods("POST")
}