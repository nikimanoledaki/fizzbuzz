package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nikimanoledaki/fizzbuzz/pkg/fizzbuzz"
)

func main() {
	numbersAsStrings := os.Args[1:]

	if len(numbersAsStrings) == 0 {
		fmt.Println("usage: fizzbuzz <number> [<number> <number> ...]")
		os.Exit(1)
	}

	for _, arg := range numbersAsStrings {
		number, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("arguments must be numbers")
			continue
		}
		fmt.Println(fizzbuzz.Says(number))
	}
}
