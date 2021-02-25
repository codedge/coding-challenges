package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(input string) (float64, error) {
	result := 0.0

	if len(input) < 1 {
		return result, nil
	}

	splitted := strings.Split(input, " ")

	if isNumeric(splitted[len(splitted)-1]) {
		return strconv.ParseFloat(splitted[len(splitted)-1], 64)
	}

	var numbers []string
	validOperations := []string{"+", "-", "*", "/"}

	for _, element := range splitted {
		if !IsValidElement(element, validOperations) {
			return -1, fmt.Errorf("invalid syntax")
		}

		if isNumeric(element) {
			numbers = append(numbers, element)
		} else {
			number1, numbers := numbers[len(numbers)-1], numbers[:len(numbers)-1]
			number2, numbers := numbers[len(numbers)-1], numbers[:len(numbers)-1]

			op1, _ := strconv.ParseFloat(number1, 64)
			op2, _ := strconv.ParseFloat(number2, 64)

			switch element {
			case "+":
				result = op2 + op1
			case "-":
				result = op2 - op1
			case "*":
				result = op2 * op1
			case "/":
				result = op2 / op1
			}

			numbers = append(numbers, strconv.FormatFloat(result, 'E', -1, 64))
		}


	}

	return result, nil
}

func Find(slice []string, val string) (int, error) {
	for i, item := range slice {
		if item == val {
			return i, nil
		}
	}
	return -1, fmt.Errorf("cannot find element '%s'", val)
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsValidElement(element string, validOperations []string) bool {
	_, err := Find(validOperations, element)
	if !isNumeric(element) && err != nil {
		return false
	}

	return true
}

func main() {
	r, _:= Calculate("6 3 /")

	fmt.Printf("ergebnis: %v", r)
}