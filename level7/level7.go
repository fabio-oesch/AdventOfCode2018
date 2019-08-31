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
	var result strings.Builder
	createHierarchy(inputs, &inst)
	root := getRoot(inst)
	nextInst := instructions{root}
	for len(nextInst) > 0 {
		//fmt.Println(root.name, nextInst)
		result.WriteString(root.name)
		fmt.Println("before:", nextInst, root.name)
		next := nextInst.getNextInst()
		fmt.Println("during:", nextInst, next.name)
		nextInst.add(root.precedes)
		fmt.Println("after:", nextInst)
		root = next
	}
	fmt.Println(result.String())
}

func createHierarchy(inputs []string, inst *instructions) {
	for _, input := range inputs {
		split := strings.Fields(input)
		parentInst := inst.find(split[1]) // where the parent sits
		childInst := inst.find(split[7])  // where the child sits
		childInst.follows = parentInst
		parentInst.precedes = append(parentInst.precedes, childInst)
	}
}

func (inst *instructions) getNextInst() *instruction {
	var nextInst *instruction
	var nextInstIndex int
	for index, i := range *inst {
		if nextInst == nil {
			nextInst = i
			nextInstIndex = index
			continue
		}
		if i.name < nextInst.name {
			nextInst = i
			nextInstIndex = index
		}
	}
	*inst = append((*inst)[:nextInstIndex], (*inst)[nextInstIndex+1:]...)
	return nextInst
}

func getRoot(inst instructions) *instruction {
	currentRoot := inst[0]
	for currentRoot.follows != nil {
		currentRoot = currentRoot.follows
	}
	return currentRoot
}

func (inst *instructions) find(name string) *instruction {
	if inst.exists(name) {
		for _, i := range *inst {
			if (*i).name == name {
				return i
			}
		}
	}
	return inst.create(name)
}

func (inst instructions) exists(name string) bool {
	for _, i := range inst {
		if (*i).name == name {
			return true
		}
	}
	return false
}

func (inst *instructions) add(i []*instruction) {
	*inst = append(*inst, i...)
}

func (inst *instructions) create(name string) *instruction {
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
