package managerJS

import (
	"fmt"
	"net/http"
)

func initAPI() {

	url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête GET:", err)
		return

	}

	defer response.Body.Close()
}
