package controller

import (
	"cocktail/Templates"
	"cocktail/structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var User = structs.User{
	Username:  "",
	Password:  "",
	Favorites: []structs.SentCocktail{},
}

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	api_url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?f=a"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, api_url, nil)
	if errReq != nil {
		fmt.Println("Erreur dans la requête de recherche : ", errReq)
	}

	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		fmt.Println("Erreur dans l'envoi de la requête de recherche : ", errRes)
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Erreur dans la lecture du body de la réponse api : ", errBody)
	}

	var Results struct {
		Drinks []structs.RawCocktail
	}
	err := json.Unmarshal(body, &Results)
	if err != nil {
		fmt.Println("Erreur dans l'interprétation de la réponse api : ", err)
		panic(err)
	}

	Cleans := []structs.SentCocktail{}

	for i := 0; i < len(Results.Drinks); i++ {
		Cleans = append(Cleans, structs.RawToClean(Results.Drinks[i]))
	}

	var Data struct {
		User   structs.User
		Drinks []structs.SentCocktail
	}

	Data.User = User
	Data.Drinks = Cleans

	Templates.Temp.ExecuteTemplate(w, "Accueil", Data)
}

func CocktailHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.RawQuery

	id := url[3:]

	api_url := "https://www.thecocktaildb.com/api/json/v1/1/lookup.php?i=" + id

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, api_url, nil)
	if errReq != nil {
		fmt.Println("Erreur dans la requête de recherche : ", errReq)
	}

	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		fmt.Println("Erreur dans l'envoi de la requête de recherche : ", errRes)
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Erreur dans la lecture du body de la réponse api : ", errBody)
	}

	var Results struct {
		Drinks []structs.RawCocktail
	}
	err := json.Unmarshal(body, &Results)
	if err != nil {
		fmt.Println("Erreur dans l'interprétation de la réponse api : ", err)
		panic(err)
	}

	clean := structs.RawToClean(Results.Drinks[0])

	var Data struct {
		User     structs.User
		Cocktail structs.SentCocktail
	}

	Data.User = User
	Data.Cocktail = clean

	Templates.Temp.ExecuteTemplate(w, "Cocktail", Data)
}

func FavorieHandler(w http.ResponseWriter, r *http.Request) {
	if User.Username == "" {
		http.Redirect(w, r, "/Login", http.StatusSeeOther)
	} else {
		Templates.Temp.ExecuteTemplate(w, "Favorie", User)
	}
}

func AddToFav(w http.ResponseWriter, r *http.Request) {
	if User.Username == "" {
		http.Redirect(w, r, "/Acceuil", http.StatusForbidden)
	} else {
		cocktailId := r.URL.Query().Get("id")
		api_url := "https://www.thecocktaildb.com/api/json/v1/1/lookup.php?i=" + cocktailId

		httpClient := http.Client{
			Timeout: time.Second * 2,
		}

		req, errReq := http.NewRequest(http.MethodGet, api_url, nil)
		if errReq != nil {
			fmt.Println("Erreur dans la requête du cocktail : ", errReq)
		}

		res, errRes := httpClient.Do(req)
		if res.Body != nil {
			defer res.Body.Close()
		} else {
			fmt.Println("Erreur dans l'envoi de la requête du cocktail : ", errRes)
		}

		body, errBody := io.ReadAll(res.Body)
		if errBody != nil {
			fmt.Println("Erreur dans la lecture du body de la réponse api : ", errBody)
		}

		var raw structs.RawCocktail
		err := json.Unmarshal(body, &raw)
		if err != nil {
			fmt.Println("Erreur dans l'interprétation de la réponse api : ", err)
			panic(err)
		}
		Clean := structs.RawToClean(raw)

		User.Favorites = append(User.Favorites, Clean)

		Json, err := os.ReadFile("./user.json")
		if err != nil {
			fmt.Println("Erreur dans le toaster ! Problème dans la lecture du fichier user.json : ", err)
		}

		var allUsers struct {
			Users []structs.User `json:"Users"`
		}
		jsonErr := json.Unmarshal(Json, &allUsers)
		if jsonErr != nil {
			panic(jsonErr)
		}

		for i := 0; i < len(allUsers.Users); i++ {
			if allUsers.Users[i].Username == User.Username {
				allUsers.Users[i] = User
			}
		}

		newJSON, err := json.MarshalIndent(allUsers, "", "    ")
		if err != nil {
			fmt.Println("Erreur dans le toaster ! Problème dans la création des nouvelles données json : ", err)
		}

		err = os.WriteFile("./user.json", newJSON, 0644)
		if err != nil {
			fmt.Println("Erreur dans le toaster ! Problème dans l'écriture du nouveau json : ", err)
		}

		http.Redirect(w, r, "/Favorie", http.StatusSeeOther)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	err := r.URL.Query().Get("err")

	var Error struct {
		Err string
	}

	if err == "already_exists" {
		Error.Err = "Nom d'utilisateur déjà utilisé"
	} else {
		Error.Err = ""
	}

	Templates.Temp.ExecuteTemplate(w, "Create", Error)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	pwd := r.PostFormValue("password")

	fmt.Println("name : ", name, " pwd : ", pwd)

	var Users struct {
		Users []structs.User `json:"Users"`
	}

	usersFile, errFile := os.ReadFile("./users.json")
	if errFile != nil {
		fmt.Println("Erreur dans la lecture du fichier users.json : ", errFile)
	}
	jsonErr := json.Unmarshal(usersFile, &Users)
	if jsonErr != nil {
		panic(jsonErr)
	}

	alreadyUsed := false

	for i := 0; i < len(Users.Users); i++ {
		if name == Users.Users[i].Username {
			alreadyUsed = true
		}
	}

	if alreadyUsed {
		http.Redirect(w, r, "/Register?err=already_used", http.StatusSeeOther)
	} else {
		newUser := structs.User{
			Username:  name,
			Password:  pwd,
			Favorites: []structs.SentCocktail{},
		}

		Users.Users = append(Users.Users, newUser)
		Users, err := json.MarshalIndent(Users, "", "    ")
		if err != nil {
			fmt.Println("Erreur dans l'écriture du json : ", err)
		}

		os.WriteFile("./users.json", Users, 0644)

		http.Redirect(w, r, "/Login", http.StatusSeeOther)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.URL.Query().Get("err")

	var Error struct {
		Err string
	}

	if err == "not_found" {
		Error.Err = "utilisateur introuvable/mauvais mot de passe"
	} else {
		Error.Err = ""
	}

	Templates.Temp.ExecuteTemplate(w, "Login", Error)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	pwd := r.PostFormValue("pwd")

	var Users struct {
		Users []structs.User `json:"Users"`
	}

	usersFile, errFile := os.ReadFile("./users.json")
	if errFile != nil {
		fmt.Println("Erreur dans la lecture du fichier users.json : ", errFile)
	}
	err := json.Unmarshal(usersFile, &Users)
	if err != nil {
		panic(err)
	}
	found := false
	goodPwd := false

	for i := 0; i < len(Users.Users); i++ {
		if Users.Users[i].Username == name {
			found = true
			if Users.Users[i].Password == pwd {
				goodPwd = true
				User = Users.Users[i]
			}
		}
	}

	if !found || !goodPwd {
		http.Redirect(w, r, "/Login?err=not_found", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/Acceuil", http.StatusSeeOther)
	}
}

func Rechercher(w http.ResponseWriter, r *http.Request) {
	search := r.Form.Get("q")

	api_url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + search

	res, errRes := http.Get(api_url)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		fmt.Println("Erreur dans l'envoi de la requête de recherche : ", errRes)
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Erreur dans la lecture du body de la réponse api : ", errBody)
	}

	var Response struct {
		Results []structs.RawCocktail `json:"drinks"`
	}

	err := json.Unmarshal(body, &Response)
	if err != nil {
		fmt.Println("Erreur dans l'interprétation de la réponse api : ", err)
	}

	fmt.Println(Response)

	if len(Response.Results) > 0 {
		Clean := structs.RawToClean(Response.Results[0])
		http.Redirect(w, r, "/Cocktail?id="+Clean.Id, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/Not_Found", http.StatusSeeOther)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	Templates.Temp.ExecuteTemplate(w, "NotFound", User)
}
