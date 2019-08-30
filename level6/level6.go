package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var devices []device

var minX, minY, maxX, maxY int
var grid [][]int

type device struct {
	x int
	y int
}

func main() {
	filename := "input.txt"
	inputs := readAndSplitFile(filename)
	createDeviceSlice(inputs)
	getMinMaxXYVals()
	limit, area := 10000, 0
	for y := minY - 1; y < maxY+2; y += 1 {
		for x := minX - 1; x < maxX+2; x += 1 {
			if distanceToAllDevices(x, y) < limit {
				area += 1
			}
		}
	}
	fmt.Println(area)
}

func firstStar() {
	getMinMaxXYVals()
	prepareGrid()
	areas := make(map[int]int)
	for y, values := range grid {
		for x := range values {
			nearestDevice := findNearestDevice(x, y)
			if x == minX-1 || y == minY-1 || x == maxX+1 || y == maxY+1 {
				areas[nearestDevice] = -1
			} else if areas[nearestDevice] != -1 {
				areas[nearestDevice] += 1
			}
		}
	}
	fmt.Println(largestArea(areas))
}

func manhattenDistance(a device, b device) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func distanceToAllDevices(x int, y int) int {
	sum := 0
	for _, d := range devices {
		sum += manhattenDistance(device{x, y}, d)
	}
	return sum
}

func largestArea(areas map[int]int) int {
	largestArea := -1
	for _, v := range areas {
		if v > largestArea {
			largestArea = v
		}
	}
	return largestArea
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func findNearestDevice(x int, y int) int {
	smallestDist := 10000
	deviceNumberOfSmallestDist := 0
	tempDevice := device{x, y}
	for i, device := range devices {
		currentDist := manhattenDistance(tempDevice, device)
		if currentDist == smallestDist {
			deviceNumberOfSmallestDist = -1
		}
		if currentDist < smallestDist {
			smallestDist = currentDist
			deviceNumberOfSmallestDist = i
		}
	}
	return deviceNumberOfSmallestDist
}

func prepareGrid() {
	limitX, limitY := maxX-minX+3, maxY-minY+3
	for i := 0; i < limitY; i += 1 {
		grid = append(grid, make([]int, limitX))
	}
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
	for _, input := range inputs {
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
