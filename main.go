package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nikimanoledaki/fizzbuzz/pkg/fizzbuzz"
)

func main() {
	numbersAsStrings := os.Args[1:]

	for _, arg := range numbersAsStrings {
		number, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("arguments must be numbers")
			os.Exit(1)
		}
		fmt.Println(fizzbuzz.Says(number))
	}
}
