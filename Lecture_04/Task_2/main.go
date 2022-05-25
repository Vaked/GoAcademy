package main

import (
	"fmt"
	"math/rand"
	"time"
)


func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100
	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
	cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}
	prices := make([]int, dataPointCount)
	for i := range prices {
	prices[i] = rand.Intn(100)
	}
	return cities, prices
}

 func groupSlices(keySlice []string, valueSlice []int) map[string][]int {
	var cityResult = make(map[string][]int)

	for idx, key := range keySlice {
		cityResult[key] = append(cityResult[key], valueSlice[idx])
	}
	return cityResult
}


func main() {
	cities, prices := citiesAndPrices()
	mapOfCitiesAndPrices := groupSlices(cities, prices)
	fmt.Println(mapOfCitiesAndPrices)
}