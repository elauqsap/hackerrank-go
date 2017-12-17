package main

import "fmt"

func main() {
	var year, date int
	fmt.Scanf("%d", &year)
	switch {
	case year >= 1700 && year <= 1917:
		if year%4 == 0 {
			date = 12
		} else {
			date = 13
		}
	case year == 1918:
		date = 26
	case year >= 1919:
		if ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0) {
			date = 12
		} else {
			date = 13
		}
	}
	fmt.Printf("%d.09.%d\n", date, year)
}
