package main

// For example:

// > delivers presents to 2 houses: one at the starting location, and one to the east.
// ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
// ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := readLines("input.txt")
	for n, directions := range input {
		fmt.Printf("Set %d: %d", n, CountHousesOnDirections(directions))
	}
}

func CountHousesOnDirections(directions string) int {
	type coords struct {
		hor int
		ver int
	}
	coordssanta := coords{0, 0}
	coordsrobot := coords{0, 0}
	visitedHouses := 1
	travelmap := make(map[[2]int]int)
	travelmap[[2]int{0, 0}] = 1
	for n, direction := range directions {
		which := &coordssanta
		if n%2 == 0 {
			which = &coordssanta
		} else {
			which = &coordsrobot
		}
		switch string(direction) {
		case "^":
			which.ver++
		case "v":
			which.ver--
		case ">":
			which.hor++
		case "<":
			which.hor--
		}
		travelmap[[2]int{which.hor, which.ver}]++
		if travelmap[[2]int{which.hor, which.ver}] == 1 {
			visitedHouses++
		}
	}
	return visitedHouses
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
