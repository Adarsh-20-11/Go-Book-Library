package main

import (
	"log"
	"net/http"

	"example.com/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", fileServer)
	http.Handle("/action", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
