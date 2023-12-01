package _1

import (
	"aoc-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type numStr struct {
	name string
	num  int
}

var nums = []numStr{
	{
		name: "one",
		num:  1,
	},
	{
		name: "two",
		num:  2,
	},
	{
		name: "three",
		num:  3,
	},
	{
		name: "four",
		num:  4,
	},
	{
		name: "five",
		num:  5,
	},
	{
		name: "six",
		num:  6,
	},
	{
		name: "seven",
		num:  7,
	},
	{
		name: "eight",
		num:  8,
	},
	{
		name: "nine",
		num:  9,
	},
}

func SumCalculator() int {
	vals := utils.ReadStringArrayFromFile("01-2")

	total := 0

	for _, val := range vals {
		comb := getCombinedNum(val)
		total += comb
	}

	return total
}

func getCombinedNum(input string) int {
	first, last := convertNumbers(input)
	joined := fmt.Sprintf("%d%d", first, last)
	num, err := strconv.Atoi(joined)
	if err != nil {
		panic(err)
	}

	return num
}

func convertNumbers(input string) (int, int) {
	// Define a map for textual representations of numbers
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		// Add more numbers as needed
	}

	// Define a regular expression to find textual representations of numbers
	//re := regexp.MustCompile(`(?:one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9)`)

	matches := findOverlappingMatches(input, "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9")

	// Find all matches in the input string

	fmt.Println(input)
	fmt.Println(matches)

	// Convert textual representations to numeric values
	var numericValues []int
	for _, match := range matches {
		numericValue, ok := numberMap[match]
		if ok {
			numericValues = append(numericValues, numericValue)
		} else {
			if matchInt, err := strconv.Atoi(match); err == nil {
				numericValues = append(numericValues, matchInt)
			}
		}
	}
	fmt.Println(numericValues)

	// Extract the first and last numeric values
	var first, last int
	if len(numericValues) > 0 {
		first = numericValues[0]
		last = numericValues[len(numericValues)-1]
	}

	return first, last
}

func findOverlappingMatches(input string, words ...string) []string {
	matches := make([]string, 0)

	for i := 0; i < len(input); i++ {
		for _, word := range words {
			if i+len(word) <= len(input) && strings.HasPrefix(input[i:], word) {
				matches = append(matches, word)
			}
		}
	}

	return matches
}
