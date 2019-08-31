package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type instructions []*instruction

type instruction struct {
	name     string
	follows  *instruction
	precedes []*instruction
}

func main() {
	filename := "testinput.txt"
	inputs := readAndSplitFile(filename)
	inst := instructions{}
	for _, input := range inputs {
		split := strings.Fields(input)
		parent := split[1]
		child := split[7]
		var parentInst *instruction
		var childInst *instruction
		if inst.exists(parent) {
			parentInst = inst.find(parent)
		} else {
			parentInst = inst.add(parent)
		}
		if inst.exists(child) {
			childInst = inst.find(child)
		} else {
			childInst = inst.add(child)
		}
		childInst.follows = parentInst
		parentInst.precedes = append(parentInst.precedes, childInst)
		fmt.Println(inst)
	}
	//instructions := make(map[string]instruction)
	//createHierarchy(inputs, instructions)
	fmt.Println(inst[0].name)
}

func (inst instructions) find(name string) *instruction {
	for _, i := range inst {
		if (*i).name == name {
			return i
		}
	}
	return nil
}

func (inst instructions) exists(name string) bool {
	for _, i := range inst {
		if (*i).name == name {
			return true
		}
	}
	return false
}

func (inst *instructions) add(name string) *instruction {
	i := instruction{name, nil, nil}
	*inst = append(*inst, &i)
	return &i
}

func (inst instructions) String() string {
	var out strings.Builder
	for _, i := range inst {
		fmt.Fprintf(&out, "name: "+i.name)
		if i.follows != nil {
			fmt.Fprintf(&out, ", follows: "+i.follows.name)
		}
		if len(i.precedes) > 0 {
			fmt.Fprintf(&out, ", precedes: ["+i.precedes[0].name)
			for _, p := range i.precedes[1:] {
				fmt.Fprintf(&out, " "+p.name)
			}
			fmt.Fprintf(&out, "]")
		}
		fmt.Fprintf(&out, "; ")
	}
	return out.String()
}

func readAndSplitFile(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(input), "\n")
	return inputs[:len(inputs)-1]
}
