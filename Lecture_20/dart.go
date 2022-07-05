package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// type BartenderBot struct{
// 	Drinks []Drinks
// }

// type Drinks struct {
// 	ID string `json:"idDrink"`
// 	Drink string `json:"strDrink"`
// 	StrDrinkAlternate string `json:"strDrinkAlternate"`
// 	StrTags string `json:"strTags"`
// 	StrVideo string `json:"strVideo"`
// 	StrAlcoholic string `json:"strAlcoholic"`
// 	StrGlass string `json:"strGlass"`
// 	StrInstructions string `json:"strInstructions"`
// }

func a() {
	url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=margarita"

	drinksClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "drinks")

	res, getErr := drinksClient.Do(req)
	if getErr != nil{
		log.Fatal(getErr)
	}

	if res.Body != nil{
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil{
		log.Fatal(readErr)
	}

	bot := BartenderBot{}
	jsonErr := json.Unmarshal(body, &bot)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	sentences := strings.Split((bot.Drinks[0].StrInstructions), ".")

	for i := 0; i < len(sentences); i++ {
		fmt.Println(sentences[i])
	}
}