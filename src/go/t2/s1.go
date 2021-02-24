package main

import "strconv"

func NumberToOrdinal(number int) string {
	if number == 0 {
		return "0"
	}

	special := []int{11, 12, 13}
	_, found := Find(special, number % 100)

	ending := "th"

	if !found {
		switch number % 10 {
		case 1:
			ending = "st"
		case 2:
			ending = "nd"
		case 3:
			ending = "rd"
		}
	}

	return strconv.Itoa(number) + ending
}

func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}