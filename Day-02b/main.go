package main

// A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
// A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.

// Ribbons:
// A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
// A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Rectangle struct {
	height int
	width  int
	length int
}

// totalPaper is the total of square feet of paper needed
var totalPaper int
var totalLint int

// type Tests struct {
// 	rectangle Rectangle
// 	answer    int
// }

// var tests = []Tests{
// 	{
// 		Rectangle{
// 			2,
// 			3,
// 			4,
// 		},
// 		58,
// 	},
// 	{
// 		Rectangle{
// 			1,
// 			1,
// 			10,
// 		},
// 		43,
// 	},
// }

func main() {
	input, _ := readLines("input.txt")
	for _, line := range input {
		sides := strings.Split(line, "x")
		a, _ := strconv.Atoi(sides[0])
		b, _ := strconv.Atoi(sides[1])
		c, _ := strconv.Atoi(sides[2])
		var rectangle = Rectangle{a, b, c}
		totalPaper += calcPaper(rectangle)
		totalLint += calcLint(rectangle)
	}
	fmt.Printf("Total square feet of paper needed: %d\nTotal lint needed: %d\n", totalPaper, totalLint)
}

// The total paper should be 2*l*w + 2*w*h + 2*h*l
func calcPaper(rectangle Rectangle) (paper int) {
	var smallside int
	a := 2 * rectangle.length * rectangle.width
	b := 2 * rectangle.width * rectangle.height
	c := 2 * rectangle.height * rectangle.length
	if b > a {
		smallside = a
	} else {
		smallside = b
	}
	if c < smallside {
		smallside = c
	}
	return a + b + c + smallside/2
}

// calcLint Calculates the needed lint according to the formula supplied.
func calcLint(rectangle Rectangle) (lint int) {
	// A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
	sorted := sortSidesbySize(rectangle)
	return 2*sorted[0] + 2*sorted[1] + sorted[0]*sorted[1]*sorted[2]
}

// sortSidesbySize sorts the sides of the rectangle by size, from small to big.
func sortSidesbySize(rectangle Rectangle) []int {
	sidesBySide := []int{rectangle.height, rectangle.length, rectangle.width}
	sort.Ints(sidesBySide)
	return sidesBySide
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
