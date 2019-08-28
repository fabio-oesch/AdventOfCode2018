package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


func main() {
	filename := "testinput.txt"
	inputs := readAndSplitFile(filename)
	fmt.Println(inputs)
}

func sortTimes(inputs []string) []string {

}

func readAndSplitFile(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(input), "\n")
	return inputs[:len(inputs)-1]
}
