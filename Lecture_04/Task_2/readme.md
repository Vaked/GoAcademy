Given this function:
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

Write a function that ranges over the two slices that citiesAndPrices
returns, and groups them in a single map, using the city name as the key.
func groupSlices(keySlice []string, valueSlice) map[string][]int {
    // ...
}
{
    "Berlin": [12, 23, 22, 12],
    "Tokyo": [45, 56, 1, 45],
}