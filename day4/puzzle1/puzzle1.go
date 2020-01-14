package puzzle1

import (
	"strconv"
)

const (
	//RangeStart Password can't be less than this
	RangeStart = 382345
	//RangeEnd Password can't be greater than this
	RangeEnd = 843167
)

func hasPairOfEqualAdjacentDigits(password string) bool {
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			return true
		}
	}
	return false
}

func IsAllDigitsIncreasingOrSame(password string) bool {
	for i := 0; i < len(password)-1; i++ {
		if password[i+1] < password[i] {
			return false
		}
	}
	return true
}

func IsWithinRange(password string) bool {
	numericPassword, _ := strconv.Atoi(password)
	return numericPassword >= RangeStart && numericPassword <= RangeEnd
}

//CheckPassword tests to see if a password meets certain criteria
func CheckPassword(password string) bool {

	return IsWithinRange(password) && hasPairOfEqualAdjacentDigits(password) && IsAllDigitsIncreasingOrSame(password)
}
