package main

import (
	"cocktail/templates"
	"cocktail/root"
)

func main() {

	root.InitServer()
	templates.InitTemplate()
}
