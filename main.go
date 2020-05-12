package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nikimanoledaki/fizzbuzz/pkg/fizzbuzz"
)

func main() {
	numberAsString := os.Args[1]
	number, _ := strconv.Atoi(numberAsString)
	fmt.Println(fizzbuzz.Says(number))
}
