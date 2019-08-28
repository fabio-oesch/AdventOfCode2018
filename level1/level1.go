package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	filename := "testinput.txt"
	inputs := readAndSplitFile(filename)
	fmt.Println(sumUpValues(inputs))
}

func sumUpValues(inputs []string) (int, int) {
	result := 0
	alreadyOccuredState := make(map[int]bool)
	alreadyOccuredState[0] = true
	firstOccurence := 0
	found := false
	for i := 0; !found;i++ {
		if i == len(inputs) {
			i = 0
		}
		if len(inputs[i]) == 0 {
			continue
		}
		result = addToResult(result, inputs[i])
		if alreadyOccuredState[result] {
			firstOccurence = result
			found = true
		} else {
			alreadyOccuredState[result] = true
		}
	}
	return result, firstOccurence
}

func addToResult(result int, toAdd string) int {
	toInt, err := strconv.Atoi(toAdd[1:])
	if err != nil {
		panic(err)
	}
	if toAdd[0] == 45 {
		toInt *= -1
	}

	return result + toInt
}

func readAndSplitFile(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(input), "\n")
}
