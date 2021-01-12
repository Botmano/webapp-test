package utils

import (
	"net/http"
	"html/template"
)

var templates *template.Template
// CarregarTemplates insere os templates html na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("view/*.html"))
}

// ExecutarTemplate renderiza a pagina html
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}){
	templates.ExecuteTemplate(w, template, dados)
}