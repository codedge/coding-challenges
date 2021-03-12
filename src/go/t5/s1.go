package main

func Kcomplement(k int, numbers []int) int {
	pairs := 0
	occurrencesNumbers := make(map[int]int)

	for _, number := range numbers {
		if _, ok := occurrencesNumbers[number]; ok {
			occurrencesNumbers[number]++
		} else {
			occurrencesNumbers[number] = 1
		}
	}

	for number, occurrence := range occurrencesNumbers {
		result := k - number
		if _, ok := occurrencesNumbers[number]; ok {
			pairs += occurrence * occurrencesNumbers[result]
		}
	}

	return pairs
}
