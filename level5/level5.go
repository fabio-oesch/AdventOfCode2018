package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// testinput.txt => dabCBAcaDA

func main() {
	filename := "input.txt"
	input := readAndSplitFile(filename)[0]
	inputlength := len(input)
	for i := 0; i < inputlength-1; i += 1 { // 'a' = 97, 'A' = 65
		if sub(input[i], input[i+1]) == 32 {
			input = input[:i] + input[i+2:]
			if i == 0 {
				i -= 1
			} else {
				i -= 2
			}
			inputlength = inputlength - 2
		}
	}
	fmt.Println(len(input))
}

func sub(x byte, y byte) byte {
	if x > y {
		return x - y
	}
	return y - x
}

func readAndSplitFile(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(input), "\n")
	return inputs[:len(inputs)-1]
}
