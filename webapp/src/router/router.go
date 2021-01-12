package router

import (
	"github.com/gorilla/mux"
)


// Gerar retorna um router com as rotas
func Gerar() *mux.Router {
	return mux.NewRouter()
}