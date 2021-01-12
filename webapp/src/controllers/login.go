package controllers

import (
	"webapp/src/utils"
	"net/http"
)
//CarregarTelaDeLogin faz oque diz o nome
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request){
	utils.ExecutarTemplate(w, "login.html", nil)
}