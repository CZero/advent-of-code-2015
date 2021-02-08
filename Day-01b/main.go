package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := readLines("input.txt")
	solution, posbasement := calcfloor(input[0])
	fmt.Printf("Solution: %d, first time in basement: %d\n", solution, posbasement)

}

// calcfloor takes an instructionset (string) and calculates the correct floor (int).
func calcfloor(instructions string) (floor, posbasement int) {
	floor = 0
	posbasement = -1
	for pos, instruction := range instructions {
		switch string(instruction) {
		case "(":
			floor++
		case ")":
			floor--
		}
		if floor == -1 && posbasement == -1 {
			posbasement = pos + 1
		}
	}
	return floor, posbasement
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
