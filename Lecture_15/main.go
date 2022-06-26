package main

import (
	"fmt"
	"sort"
	"time"
)

func sortDates(format string, dates ...string) ([]string, error) {
	parsedDates := []time.Time{}
	var err error
	var parsedDate time.Time
	var sortedDateStringSlice = make([]string, 0)

	for _, date := range dates {
		parsedDate, err = time.Parse(format, date)
		parsedDates = append(parsedDates, parsedDate)
	}

	sort.Slice(parsedDates, func(i, j int) bool {
		return parsedDates[i].Before(parsedDates[j])
	})

	for _, date := range parsedDates {
		sortedDateStringSlice = append(sortedDateStringSlice, date.String())
	}

	if err != nil{
		return nil, err
	}
	return sortedDateStringSlice, nil
}

func main() {
	marDate := "Mar-18-2022"
	septDate := "Sep-14-2008"
	decDate := "Dec-03-2021"

	format := ("Jan-2-2006")
	fmt.Println(sortDates(format, marDate, decDate,septDate))

}