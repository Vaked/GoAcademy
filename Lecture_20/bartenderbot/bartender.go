package bartenderbot

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const(
	cocktailDbApiUrl string = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="
)

type BartenderBot struct{
	Drinks []Drinks
}

type Drinks struct {
	ID string `json:"idDrink"`
	Drink string `json:"strDrink"`
	StrDrinkAlternate string `json:"strDrinkAlternate"`
	StrTags string `json:"strTags"`
	StrVideo string `json:"strVideo"`
	StrAlcoholic string `json:"strAlcoholic"`
	StrGlass string `json:"strGlass"`
	StrInstructions string `json:"strInstructions"`
}
func NewBartenderBot() *BartenderBot{
	return &BartenderBot{}
}

func GetUserInput() string{
	fmt.Println("Please enter what you would like to drink: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading input.")
		return ""
	}
	input = strings.TrimRight(input, "\r\n")

	return input
}

func SetUrl(querryString string) *url.URL{
	url, err := url.Parse(cocktailDbApiUrl)
	if err != nil{
		log.Fatal("Url format is invalid")
	}

	querry := url.Query()
	querry.Add("s", querryString)
	url.RawQuery = querry.Encode()

	return url
}

func GetResultFromDrinksApi(drink *url.URL) {
	drinksClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, drink.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

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

	bot := NewBartenderBot()
	jsonErr := json.Unmarshal(body, &bot)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	SplitAndPrintRecipe(bot.Drinks[0].StrInstructions)
}

func SplitAndPrintRecipe(recpe string) {
	sentences := strings.Split(recpe, ".")

	for i := 0; i < len(sentences); i++{
		fmt.Println(sentences[i])
	}
}

func (b *BartenderBot) Start() {
	for {
		prefferedDrink := GetUserInput()
		if prefferedDrink == "nothing"{
			break
		}

		SetUrl(prefferedDrink)
		GetResultFromDrinksApi(SetUrl(prefferedDrink))
	}
}