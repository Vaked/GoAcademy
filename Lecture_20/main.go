// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"

// 	//"strings"
// 	"time"
// )

// type BartenderBot struct{
// 	recipe string `json"recipe"`
// }

// func NewBartenderBot() *BartenderBot{
// 	return &BartenderBot{}
// }

// func GetUserInput() string{
// 	fmt.Println("Please enter what you would like to drink: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error while reading input.")
// 		return ""
// 	}
// 	return input
// }

// func GetResultFromDrinksApi(drink string) {
// 	url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=margarita"

// 	drinksClient := http.Client{
// 		Timeout: time.Second * 2,
// 	}

// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	res, getErr := drinksClient.Do(req)
// 	if getErr != nil{
// 		log.Fatal(getErr)
// 	}

// 	if res.Body != nil{
// 		defer res.Body.Close()
// 	}

// 	body, readErr := ioutil.ReadAll(res.Body)
// 	if readErr != nil{
// 		log.Fatal(readErr)
// 	}

// 	bot := NewBartenderBot()
// 	jsonErr := json.Unmarshal(body, &bot)

// 	if jsonErr != nil {
// 		log.Fatal(jsonErr)
// 	}

// 	fmt.Println(bot.recipe)

// }

// func (b *BartenderBot) Start() {
// 	prefferedDrink := GetUserInput()

// 	GetResultFromDrinksApi(prefferedDrink)

	
// }

// func fart() {
// 	bot := NewBartenderBot()

// 	bot.Start()
// }