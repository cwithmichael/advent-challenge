package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cwithmichael/advent/day5/puzzle1"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	opcodes := strings.Split(scanner.Text(), ",")
	intOpcodes := make([]int, len(opcodes))
	for k, v := range opcodes {
		intOpcodes[k], _ = strconv.Atoi(v)
	}

	puzzle1.ParseOpcodes(intOpcodes, 1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
