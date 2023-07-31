package routes

import (
	"Go_Bookstore_Management/controllers"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		m := map[string]string{
			"msg": "hello",
		}
		res, _ := json.Marshal(m)
		writer.Write(res)
	})
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
