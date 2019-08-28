package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filename := "testinput.txt"
	inputs := readAndSplitFile(filename)
	containsTwo := 0
	containsThree := 0
	for _, input := range inputs {
		if hasLetterCount(input, 2) {
			containsTwo += 1
		}
		if hasLetterCount(input, 3) {
			containsThree += 1
		}
	}
	fmt.Println(containsTwo, containsThree)
	fmt.Println(containsTwo * containsThree)
}

func hasLetterCount(input string, specificCount int) bool {
	count := make(map[rune]int)
	for _, c := range input {
		count[c] += 1
	}
	for _, entry := range count {
		if entry == specificCount {
			return true
		}
	}
	return false
}

func readAndSplitFile(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(input), "\n")
	return inputs[:len(inputs)-1]
}
