package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filename := "testinput.txt"
	inputs := readAndSplitFile(filename)
	fmt.Println(getSimilarStrings(inputs))
}

func getSimilarStrings(inputs []string) string {
	var remaining string
	for i := 0; i < len(inputs); i += 1 {
		for j := i + 1; j < len(inputs); j += 1 {
			remaining = getRemaining(inputs[i], inputs[j])
			if remaining != "" {
				return remaining
			}
		}
	}
	panic("No similar ones found")
}

func differsByOneChar(s1 string, s2 string) (bool, int) {
	differByOne := true
	pos := -1
	for i := range s1 {
		if s1[i] != s2[i] {
			if differByOne {
				differByOne = false
				pos = i
			} else {
				return false, -1
			}
		}
	}
	return true, pos
}

func getRemaining(s1 string, s2 string) string {
	differsByOne, pos := differsByOneChar(s1, s2)
	if differsByOne {
		return s1[:pos] + s1[pos + 1:]
	}
	return ""
}

func containsTwoAndThreeOfSameLetter(inputs []string) {
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
