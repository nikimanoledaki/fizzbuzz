package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nikimanoledaki/fizzbuzz/pkg/fizzbuzz"
)

func main() {
	numberAsString := os.Args[1]
	number, err := strconv.Atoi(numberAsString)
	if err != nil {
		fmt.Println("argument must be a number")
		os.Exit(1)
	}
	fmt.Println(fizzbuzz.Says(number))

}
