package controller

import (
	"cocktail/templates"
	"net/http"
)

func AccueilHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "Accueil", nil)
}
