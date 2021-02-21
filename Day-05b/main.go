package main

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
	if TestOne(s) && TestTwo(s) {
		return "nice"
	}
	return "naughty"
}

// TestOne checks if it contains a pair of any two letters that appears at least twice in the string without overlapping,
// like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
func TestOne(s string) bool {
	limit := len(s)
	if limit < 4 {
		return false
	}
	for n := range s {
		if n == limit-3 {
			return false
		}
		for j := n + 2; j <= limit-2; j++ {
			if s[j:j+2] == s[n:n+2] {
				return true
			}
		}
	}

	return false
}

// NiceTestTwo checks if it contains at least one letter which repeats with exactly one letter between them,
// like xyx, abcdefeghi (efe), or even aaa.
func TestTwo(s string) bool {
	lim := len(s)
	for n, char := range s {
		if n <= lim-3 {
			if string(char) == string(s[n+2]) {
				return true
			}
		}
	}
	return false
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
