package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	filename := "testinput.txt"
	//squareSize := 10
	squareSize := 1000
	fabric := make([][]int, squareSize)
	for row := range fabric {
		fabric[row] = make([]int, squareSize)
	}
	inputs := readAndSplitFile(filename)
	claimAmount := make([]bool, 1266)
	for i, input := range inputs {
		claimFabric2(fabric, input, i + 1, claimAmount)
	}
	for i := range claimAmount {
		if !claimAmount[i] {
			fmt.Println(i)
		}
	}
}

func firstStar(fabric [][]int, inputs []string) {
	overlappingPieces := 0
	for _, input := range inputs {
		overlappingPieces += claimFabric(fabric, input)
	}
	fmt.Println(overlappingPieces)
}

func claimFabric2(fabric [][]int, claim string, currentClaimNumber int, claimAmount []bool) {
	indexOfAt := strings.Index(claim, "@")
	indexOfColon := strings.Index(claim, ":")
	x, y := getXandYPos(claim[indexOfAt+2 : indexOfColon])
	i, j := getDimensions(claim[indexOfColon+2:])
	for row := x; row < x+i; row += 1 {
		for col := y; col < y+j; col += 1 {
			if fabric[row][col] != 0 {
				claimAmount[currentClaimNumber] = true
				claimAmount[fabric[row][col]] = true
			} else {
				fabric[row][col] = currentClaimNumber
			}
		}
	}
}

func claimFabric(fabric [][]int, claim string) int {
	overlapping := 0
	indexOfAt := strings.Index(claim, "@")
	indexOfColon := strings.Index(claim, ":")
	x, y := getXandYPos(claim[indexOfAt+2 : indexOfColon])
	i, j := getDimensions(claim[indexOfColon+2:])
	for row := x; row < x+i; row += 1 {
		for col := y; col < y+j; col += 1 {
			fabric[row][col] += 1
			if fabric[row][col] == 2 {
				overlapping += 1
			}
		}
	}
	return overlapping
}

func getXandYPos(claim string) (int, int) {
	indexOfComma := strings.Index(claim, ",")
	x, err := strconv.Atoi(claim[:indexOfComma])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(claim[indexOfComma+1:])
	if err != nil {
		panic(err)
	}
	return x, y
}

func getDimensions(claim string) (int, int) {
	indexOfX := strings.Index(claim, "x")
	i, err := strconv.Atoi(claim[:indexOfX])
	if err != nil {
		panic(err)
	}
	j, err := strconv.Atoi(claim[indexOfX+1:])
	if err != nil {
		panic(err)
	}
	return i, j
}

func readAndSplitFile(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(input), "\n")
	return inputs[:len(inputs)-1]
}
