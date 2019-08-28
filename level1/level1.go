package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filename := "input.txt"
	input := readAndSplitFile(filename)
	fmt.Println(input)
}

func readAndSplitFile(filename string) []string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(input), "\n")
}
