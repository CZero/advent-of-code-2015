package main

import (
	"bufio"
	"fmt"
	"os"
)

/* Here's an easy puzzle to warm you up.

Santa is trying to deliver presents in a large apartment building, but he can't find the right floor - the directions he got are a little confusing. He starts on the ground floor (floor 0) and then follows the instructions one character at a time.

An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ), means he should go down one floor.

The apartment building is very tall, and the basement is very deep; he will never find the top or bottom floors.

For example:

(()) and ()() both result in floor 0.
((( and (()(()( both result in floor 3.
))((((( also results in floor 3.
()) and ))( both result in floor -1 (the first basement level).
))) and )())()) both result in floor -3.
To what floor do the instructions take Santa?
*/

var tests = []struct {
	test   string
	answer int
}{
	{
		"(())",
		0,
	},
	{
		"(((",
		3,
	},
	{
		"))(((((",
		3,
	},
	{
		"())",
		-1,
	},
	{
		")))",
		-3,
	},
}

func main() {
	input, _ := readLines("input.txt")
	for n, test := range tests {
		fmt.Printf("%d: %s = %d (should be %d)\n", n, test.test, calcfloor(test.test), test.answer)
	}
	fmt.Printf("Solution: %d\n", calcfloor(input[0]))

}

func calcfloor(instructions string) int {
	floor := 0
	for _, instruction := range instructions {
		switch string(instruction) {
		case "(":
			floor++
		case ")":
			floor--
		}
	}
	return floor
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
