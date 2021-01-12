package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)
// Rota é a estrutura das rotas
type Rota struct {
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
// Configurar retorna a rota dentro do router
func Configurar(router *mux.Router) *mux.Router {

	rotas := rotasLogin

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))


	for _, rota := range rotas{
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return router
}