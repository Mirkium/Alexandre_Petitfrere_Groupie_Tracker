package structs

type User struct {
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Favorites []SentCocktail `json:"favorites"`
}

type RawCocktail struct {
	Id           string `json:"idDrink"`
	Name         string `json:"strDrink"`
	Alcoholic    string `json:"strAlcoholic"`
	Instructions string `json:"strInstructions"`
	Ingredient1  string `json:"strIngredient1"`
	Ingredient2  string `json:"strIngredient2"`
	Ingredient3  string `json:"strIngredient3"`
	Ingredient4  string `json:"strIngredient4"`
	Ingredient5  string `json:"strIngredient5"`
	Ingredient6  string `json:"strIngredient6"`
	Ingredient7  string `json:"strIngredient7"`
	Ingredient8  string `json:"strIngredient8"`
	Ingredient9  string `json:"strIngredient9"`
	Ingredient10 string `json:"strIngredient10"`
	Ingredient11 string `json:"strIngredient11"`
	Ingredient12 string `json:"strIngredient12"`
	Ingredient13 string `json:"strIngredient13"`
	Ingredient14 string `json:"strIngredient14"`
	Ingredient15 string `json:"strIngredient15"`
	Measure1     string `json:"strMeasure1"`
	Measure2     string `json:"strMeasure2"`
	Measure3     string `json:"strMeasure3"`
	Measure4     string `json:"strMeasure4"`
	Measure5     string `json:"strMeasure5"`
	Measure6     string `json:"strMeasure6"`
	Measure7     string `json:"strMeasure7"`
	Measure8     string `json:"strMeasure8"`
	Measure9     string `json:"strMeasure9"`
	Measure10    string `json:"strMeasure10"`
	Measure11    string `json:"strMeasure11"`
	Measure12    string `json:"strMeasure12"`
	Measure13    string `json:"strMeasure13"`
	Measure14    string `json:"strMeasure14"`
	Measure15    string `json:"strMeasure15"`
	ImageSrc     string `json:"strDrinkThumb"`
}

type SentCocktail struct {
	Id           string
	Name         string
	Alcoholic    bool
	Ingredients  []string
	Measures     []string
	ImageSrc     string
	Instructions string
}

func RawToClean(raw RawCocktail) SentCocktail {
	var clean SentCocktail

	clean.Name = raw.Name
	clean.Id = raw.Id
	clean.ImageSrc = raw.ImageSrc
	clean.Instructions = raw.Instructions

	if raw.Alcoholic == "Alcoholic" {
		clean.Alcoholic = true
	} else {
		clean.Alcoholic = false
	}

	clean.Ingredients = append(clean.Ingredients, raw.Ingredient1)
	if raw.Ingredient2 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient2)
	}
	if raw.Ingredient3 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient3)
	}
	if raw.Ingredient4 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient4)
	}
	if raw.Ingredient5 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient5)
	}
	if raw.Ingredient6 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient6)
	}
	if raw.Ingredient7 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient7)
	}
	if raw.Ingredient8 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient8)
	}
	if raw.Ingredient9 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient9)
	}
	if raw.Ingredient10 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient10)
	}
	if raw.Ingredient11 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient11)
	}
	if raw.Ingredient12 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient12)
	}
	if raw.Ingredient13 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient13)
	}
	if raw.Ingredient14 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient14)
	}
	if raw.Ingredient15 != "" {
		clean.Ingredients = append(clean.Ingredients, raw.Ingredient15)
	}

	clean.Measures = append(clean.Measures, raw.Measure1)
	if raw.Measure2 != "" {
		clean.Measures = append(clean.Measures, raw.Measure2)
	}
	if raw.Measure3 != "" {
		clean.Measures = append(clean.Measures, raw.Measure3)
	}
	if raw.Measure4 != "" {
		clean.Measures = append(clean.Measures, raw.Measure4)
	}
	if raw.Measure5 != "" {
		clean.Measures = append(clean.Measures, raw.Measure5)
	}
	if raw.Measure6 != "" {
		clean.Measures = append(clean.Measures, raw.Measure6)
	}
	if raw.Measure7 != "" {
		clean.Measures = append(clean.Measures, raw.Measure7)
	}
	if raw.Measure8 != "" {
		clean.Measures = append(clean.Measures, raw.Measure8)
	}
	if raw.Measure9 != "" {
		clean.Measures = append(clean.Measures, raw.Measure9)
	}
	if raw.Measure10 != "" {
		clean.Measures = append(clean.Measures, raw.Measure10)
	}
	if raw.Measure11 != "" {
		clean.Measures = append(clean.Measures, raw.Measure11)
	}
	if raw.Measure12 != "" {
		clean.Measures = append(clean.Measures, raw.Measure12)
	}
	if raw.Measure13 != "" {
		clean.Measures = append(clean.Measures, raw.Measure13)
	}
	if raw.Measure14 != "" {
		clean.Measures = append(clean.Measures, raw.Measure14)
	}
	if raw.Measure15 != "" {
		clean.Measures = append(clean.Measures, raw.Measure15)
	}

	return clean
}
