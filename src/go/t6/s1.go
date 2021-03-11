package main

import "strconv"

func Fib(number int) int  {
	if number < 0 {
		return -1
	}

	if number <= 1 {
		return number
	}

	return Fib(number-1) + Fib(number-2)
}

func Fibonacci(number int) (int, error) {
	result := Fib(number)
	s := strconv.Itoa(result)

	if len(s) > 6 {
		return strconv.Atoi(s[len(s)-6:])
	}

	return strconv.Atoi(s)
}