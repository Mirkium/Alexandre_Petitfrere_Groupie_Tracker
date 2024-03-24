package root

import (
	"cocktail/controller"
	"net/http"
	"os"
)

func InitServer() {
	rootDoc, _ := os.Getwd()
	FileServer := http.FileServer(http.Dir(rootDoc + "/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", FileServer))

	http.HandleFunc("/Acceuil", controller.AccueilHandler)
	http.HandleFunc("/Cocktail", controller.CocktailHandler)
	http.HandleFunc("/Favorie", controller.FavorieHandler)
	http.HandleFunc("/Register", controller.Register)
	http.HandleFunc("/enregistrement", controller.RegisterHandler)
	http.HandleFunc("/Login", controller.Login)
	http.HandleFunc("/LoginHandler", controller.LoginHandler)
	http.HandleFunc("/rechercher", controller.Rechercher)
	http.HandleFunc("/Not_Found", controller.NotFound)

	http.ListenAndServe("localhost:8080", nil)
}
