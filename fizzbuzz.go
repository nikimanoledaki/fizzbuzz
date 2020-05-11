package fizzbuzz

import "fmt"

// IsDivisibleBy returns true if remainder equals zero.
func IsDivisibleBy(number, divisor int) bool {
	return number%divisor == 0
}

// Says 'fizz' if it's divisible by 3, 'buzz' by 5, 'fizzbuzz' by 15.
func Says(number int) string {
	switch {
	case IsDivisibleBy(number, 15):
		return "fizzbuzz"
	case IsDivisibleBy(number, 3):
		return "fizz"
	case IsDivisibleBy(number, 5):
		return "buzz"
	default:
		return fmt.Sprintf("%d", number)
	}
}
