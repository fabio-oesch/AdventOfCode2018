package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
)

var devices []device

var minX, minY, maxX, maxY int

type device struct {
	x int
	y int
}

func main() {
	filename := "testinput.txt"
	inputs := readAndSplitFile(filename)
	createDeviceSlice(inputs)
	getMinMaxXYVals()
	fmt.Println(devices, minX, minY, maxX, maxY)
}

func manhattenDistance(a device, b device) int {
	return abs(a.x - b.x) + abs(a.y - b.y)
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func getMinMaxXYVals() {
	minX = devices[0].x
	maxX = devices[0].x
	minY = devices[0].y
	maxY = devices[0].y
	for _, device := range devices[1:] {
		if device.x < minX {
			minX = device.x
		}
		if device.x > maxX {
			maxX = device.x
		}
		if device.y < minY {
			minY = device.y
		}
		if device.y > maxY {
			maxY = device.y
		}
	}
}

func createDeviceSlice(inputs []string) {
	for _,input := range inputs {
		splitInput := strings.Split(input, ", ")
		tempX, err := strconv.Atoi(splitInput[0])
		if err != nil {
			log.Fatal(err)
		}
		tempY, err2 := strconv.Atoi(splitInput[1])
		if err2 != nil {
			log.Fatal(err2)
		}
		devices = append(devices, device{tempX, tempY})
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
