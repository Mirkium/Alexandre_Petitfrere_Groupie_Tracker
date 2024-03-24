package main

import (
	Templates "cocktail/Templates"
	"cocktail/root"
)

func main() {
	Templates.InitTemplate()
	root.InitServer()
}
