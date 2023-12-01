package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadStringArrayFromFile(fileName string) []string {
	body, err := os.ReadFile(fmt.Sprintf("./input/%s.txt", fileName))
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	splitted := strings.Split(string(body), "\n")

	cleaned := deleteEmpty(splitted)

	return cleaned
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
