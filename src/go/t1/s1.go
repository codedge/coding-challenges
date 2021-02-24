package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Maskify(cc string) string {
	var masked = make([]string, len(cc))
	var splitted = strings.Split(cc, "")

	for idx, item := range splitted {
		// Exclude first and last four items
		if idx == 0 || idx >= len(splitted) - 4 {
			masked = append(masked, item)
			continue
		}

		// Check if is numeric
		_, err := strconv.Atoi(item)

		if err != nil {
			masked = append(masked, item)
			continue
		}

		masked = append(masked, "#")
	}

	return strings.Join(masked, "")
}

func main()  {
	cc := ""

	if len(os.Args) > 1 {
		cc = os.Args[1]
	} else {
		os.Exit(0)
	}

	// Do no mask if cc less than 6 or empty string
	if len(cc) < 6 {
		fmt.Println(cc)
		os.Exit(0)
	}

	fmt.Printf(Maskify(cc))
}