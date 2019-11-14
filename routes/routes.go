package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"sinosigorest/controller"
)

func CarregarRotas() {

	router := mux.NewRouter().StrictSlash(true)


	//http.HandleFunc("/", controller.Index)
	router.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
}
