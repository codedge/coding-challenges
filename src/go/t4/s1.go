package main

import "fmt"

func JumpOutOfArray(numbers []int) int {
	sum := 0
	jumps := 0

	for index, number := range numbers {
		sum += numbers[sum]

		// Jump out of array at the start of the array
		if (index + number) < 0 {
			jumps++
			break
		}

		jumps++

		// Jump out of array at the end of the array
		if sum > len(numbers) {
			break
		}

		// End of array but do not jump out
		if (index + 1) == len(numbers) {
			jumps = -1
		}
	}

	return jumps
}

func main() {
	fmt.Println(JumpOutOfArray([]int { 2, 3, -1, 1, 6 }))
}