package root

import (
	"cocktail/controller"
	"net/http"
)

func InitServer() {

	FileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", FileServer))
	http.HandleFunc("/Acceuil", controller.AccueilHandler)
	http.HandleFunc("/Cocktail", controller.CocktailHandler)
	http.HandleFunc("/Favorie", controller.FavorieHandler)
}
