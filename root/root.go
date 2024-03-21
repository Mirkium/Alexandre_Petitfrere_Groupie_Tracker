package root

import (
	"net/http"
	"cocktail/controller"
)

func InitServer() {

	FileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", FileServer))
	http.HandleFunc("/Acceuil", controller.AccueilHandler)
}
