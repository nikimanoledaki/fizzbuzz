package fizzbuzz

// IsDivisibleByThree returns false
func IsDivisibleByThree(number int) bool {
	if number%3 == 0 {
		return true
	}
	return false
}

func IsDivisibleByFive(number int) bool {
	if number%5 == 0 {
		return true
	}
	return false
}
