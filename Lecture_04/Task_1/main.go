package main

import (
	"fmt"
)

func checkYearValue(year int) bool {
	if year < 0 {
		return false
	}
	return true
}

func checkMonthValue(month int) bool {
	if month < 1 || month > 12 {
		return false
	}
	return true
}

func checkLeapYear(year int) bool {
	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
		return true
	}
	return false
}

func daysInMonth(month int, year int) (int, bool) {
	isMonthValueCorrect := checkMonthValue(month)
	isYearValueCorrect := checkYearValue(year)
	isLeapYear := checkLeapYear(year)

	if isLeapYear && isMonthValueCorrect && isYearValueCorrect {
		return 29, true
	}

	if isMonthValueCorrect && isYearValueCorrect {
		var days int

		switch month {
		case 2:
			days = 28
			fmt.Printf("The month has %d days ", days)
			return 28, true
		case 4, 6, 9, 11:
			days = 30
			fmt.Printf("The month has %d days", days)
			return 30, true
		case 1, 3, 5, 7, 8, 10, 12:
			days = 31
			fmt.Printf("The month has %d days", days)
			return 31, true
		}
	}
	return 0, false
}

func main() {
	fmt.Println("Wrong month")
	fmt.Println(daysInMonth(14, 2016))

	fmt.Println("Negative month")
	fmt.Println(daysInMonth(-4, 2016))

	fmt.Println("Check leap year february")
	fmt.Println(daysInMonth(2, 2016))

	fmt.Println("Check standard year february")
	fmt.Println(daysInMonth(2, 2015))

	fmt.Println("Check december")
	fmt.Println(daysInMonth(2, 2015))
}