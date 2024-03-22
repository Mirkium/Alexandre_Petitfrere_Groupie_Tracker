package controller

import (
	"cocktail/templates"
	"net/http"
)

func AccueilHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "Accueil", nil)
}

func CocktailHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "Cocktail", nil)
}

func FavorieHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "Favorie", nil)
}
