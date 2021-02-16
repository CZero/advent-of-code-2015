package main

// A nice string is one with all of the following properties:

// It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
// It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
// It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, _ := readLines("input.txt")
	var niceStrings int
	for _, string := range lines {
		if NaughtyOrNice(string) == "nice" {
			niceStrings++
		}
	}
	fmt.Printf("Nice strings foudn: %d", niceStrings)
}

// NaughtyOrNice checks a string to see if the person is naughty or nice by doing tests.
func NaughtyOrNice(s string) string {
	if NiceTestOne(s) && NiceTestTwo(s) && NiceTestThree(s) {
		return "nice"
	}
	return "naughty"
}

// NiceTestOne checks if the string contains at least three vowels (aeiou)
func NiceTestOne(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	neededVowels := 3

	var n int
	for _, char := range s {
		for _, vowel := range vowels {
			if vowel == string(char) {
				n++
				if n == neededVowels {
					return true
				}
				break
			}
		}
	}
	return false
}

// NiceTestTwo checks if it contains at least one letter that appears twice in a row (xx, dd, cc ..)
func NiceTestTwo(s string) bool {
	var previous string
	for _, char := range s {
		if string(char) == previous {
			return true
		}
		previous = string(char)
	}
	return false
}

// NiceTestThree checks if it does NOT contain the strings ab, cd, pq or xy
func NiceTestThree(s string) bool {
	var forbidden = []string{"ab", "cd", "pq", "xy"}
	var previous string
	for _, char := range s {
		for _, forbid := range forbidden {
			if previous+string(char) == forbid {
				return false
			}
		}
		previous = string(char)
	}
	return true
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
