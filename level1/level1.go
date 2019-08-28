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

func sumUpValues(inputs []string) int {
	result := 0
	for _, input := range inputs {
		if len(input) == 0 {
			break
		}
		toInt, err := strconv.Atoi(input[1:])
		if err != nil {
			panic(err)
		}
		if input[0] == 43 {
			result += toInt
		} else {
			result -= toInt
		}
	}
	return result
}

func readAndSplitFile(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(input), "\n")
}
