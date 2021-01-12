package rotas

import (
	"webapp/src/controllers"
	"net/http"
)

var rotasLogin = []Rota{
{
	URI: "/",
	Metodo: http.MethodGet,
	Funcao: controllers.CarregarTelaDeLogin,
	RequerAutenticacao: false,
},

{
	URI: "/login",
	Metodo: http.MethodGet,
	Funcao: controllers.CarregarTelaDeLogin,
	RequerAutenticacao: false,
},

}