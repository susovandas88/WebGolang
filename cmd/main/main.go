package main

import (
	"firstwebproject/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	//fmt.Println("Hello Susovan!!")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

}
