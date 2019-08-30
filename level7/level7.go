package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type instruction struct {
	name     string
	follows  *instruction
	precedes []*instruction
}

func main() {
	filename := "testinput.txt"
	inputs := readAndSplitFile(filename)
	instructions := make(map[string]instruction)
	createHierarchy(inputs, instructions)
	fmt.Println(instructions["A"].follows.precedes)
	fmt.Println(instructions)
}

func createHierarchy(inputs []string, instructions map[string]instruction) {
	for _, input := range inputs {
		split := strings.Fields(input)
		parent := split[1]
		child := split[7]
		childInstruction, ok := instructions[child]
		if !ok {
			 childInstruction = instruction{child, nil, nil}
		}
		parentInstruction, ok := instructions[parent]
		if !ok {
			parentInstruction = instruction{parent, nil, []*instruction{}}
		}
		childInstruction.follows = &parentInstruction
		parentInstruction.precedes = append(parentInstruction.precedes, &childInstruction)
		instructions[parent], instructions[child] = parentInstruction, childInstruction
	}
}

func readAndSplitFile(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(input), "\n")
	return inputs[:len(inputs)-1]
}
