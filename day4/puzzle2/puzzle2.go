package puzzle2

import (
	"github.com/cwithmichael/advent/day4/puzzle1"
)

func hasSinglePairOfEqualAdjacentDigits(password string) bool {
	for i := 0; i < len(password)-1; i++ {
		count := 0
		idx := i
		for idx < len(password)-1 && password[idx] == password[idx+1] {
			count += 2
			idx++
		}
		if count == 2 {
			return true
		}
		i = idx
	}
	return false
}

//CheckPassword tests to see if a password meets certain criteria
func CheckPassword(password string) bool {

	return puzzle1.IsWithinRange(password) &&
		puzzle1.IsAllDigitsIncreasingOrSame(password) &&
		hasSinglePairOfEqualAdjacentDigits(password)
}
