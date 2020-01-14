package main

import (
	"fmt"
	"strconv"

	"github.com/cwithmichael/advent/day4/puzzle1"
	"github.com/cwithmichael/advent/day4/puzzle2"
)

func main() {
	total := 0
	for i := puzzle1.RangeStart; i < puzzle1.RangeEnd; i++ {
		if puzzle1.CheckPassword(strconv.Itoa(i)) {
			total++
		}
	}
	fmt.Println("Puzzle 1 Solution: ", total)

	total = 0
	for i := puzzle1.RangeStart; i < puzzle1.RangeEnd; i++ {
		if puzzle2.CheckPassword(strconv.Itoa(i)) {
			total++
		}
	}
	fmt.Println("Puzzle 2 Solution: ", total)
}
