package main

import (
	"net/http"
	"sinosigorest/controller"
	"github.com/gorilla/mux"
)

func main() {

	//routes.CarregarRotas()
	//http.ListenAndServe(":7000", nil)

	router := mux.NewRouter()

	router.HandleFunc("/", controller.Index)
	router.HandleFunc("/denuncias", controller.GetAll).Methods("GET")
	router.HandleFunc("/denuncia/{id}", controller.GetDenunciaPorId).Methods("GET")
	router.HandleFunc("/denuncia", controller.PostNovaDenuncia).Methods("POST")
	router.HandleFunc("/novadenunciateste", controller.NovoPostTeste).Methods("POST")

	http.ListenAndServe(":7000", router)
}
