package main

import (
	"webapp/src/utils"
	"webapp/src/router"
	"net/http"
	"log"
	"fmt"
)

func main(){
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Rodando o Webapp!")
	log.Fatal(http.ListenAndServe(":9000",r))

}