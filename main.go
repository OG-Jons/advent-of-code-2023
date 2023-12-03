package main

import (
	dayOne "aoc-2023/01"
	dayTwo "aoc-2023/02"
	"fmt"
)

func main() {
	fmt.Println("-------------------")
	fmt.Println("Running Code for Day 01:")
	fmt.Printf("Day One Total: %d \n", dayOne.SumCalculator())
	fmt.Println("-------------------")
	fmt.Println("Running Code for Day 02:")
	fmt.Printf("Day Two Total: %d \n", dayTwo.SumOfInvalidGameIDs())
	fmt.Println("-------------------")
}
