package routes

import (
	"example.com/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/actions/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/actions/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/actions/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/actions/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/actions/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
