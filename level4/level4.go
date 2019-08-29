package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	inputs := readAndSplitFile(filename)
	sortTimes(inputs)
	guardsSleepSchedule := totalSleepTimePerGuard(inputs)
	for k, v := range guardsSleepSchedule {
		fmt.Println(k, v)
	}
	mostSleptTime(guardsSleepSchedule)
}

func firstStar(guardsSleepSchedule map[int][]int) {
	sleepiestguard := getLongestSleepingGuard(guardsSleepSchedule)
	_, mostSleptAt := getTimeMostSleptAt(guardsSleepSchedule, sleepiestguard)
	fmt.Println(sleepiestguard, mostSleptAt, sleepiestguard*mostSleptAt)
}

func sortTimes(inputs []string) {
	sort.Strings(inputs)
}

func mostSleptTime(guardsSleep map[int][]int) {
	mostSleptTime := -1
	guardMostSleptTime := 0
	timeAt := 0
	for k := range guardsSleep {
		currentSleep, currentTimeAt := getTimeMostSleptAt(guardsSleep, k)
		if currentSleep > mostSleptTime {
			mostSleptTime = currentSleep
			guardMostSleptTime = k
			timeAt = currentTimeAt
		}
	}
	fmt.Println(mostSleptTime, guardMostSleptTime, timeAt*guardMostSleptTime)
}

func getLongestSleepingGuard(guardsSleep map[int][]int) int {
	maxSleep := -1
	maxSleepingGuard := -1
	for k, v := range guardsSleep {
		totalSleep := 0
		for _, i := range v {
			totalSleep += i
		}
		if totalSleep > maxSleep {
			maxSleep = totalSleep
			maxSleepingGuard = k
		}
	}
	return maxSleepingGuard
}

func getTimeMostSleptAt(guardsSleep map[int][]int, guard int) (int, int) {
	mostSleptAt := -1
	var timeMostSleptAt int
	for i, v := range guardsSleep[guard] {
		if mostSleptAt < v {
			mostSleptAt = v
			timeMostSleptAt = i
		}
	}
	return mostSleptAt, timeMostSleptAt
}

func totalSleepTimePerGuard(inputs []string) map[int][]int {
	currentGuard := -1
	fellAsleepAt := -1
	sleepTime := make(map[int][]int)
	for _, input := range inputs {
		newGuard := strings.Index(input, "#")
		if newGuard != -1 {
			input = input[newGuard+1:]
			currentGuard, _ = strconv.Atoi(input[:strings.Index(input, " ")])
			if _, ok := sleepTime[currentGuard]; !ok {
				sleepTime[currentGuard] = make([]int, 60)
			}
			continue
		}
		getColonPos := strings.Index(input, ":")
		if strings.Index(input, "falls asleep") != -1 {
			fellAsleepAt, _ = strconv.Atoi(input[getColonPos+1 : getColonPos+3])
			continue
		}
		wakeUpAt, _ := strconv.Atoi(input[getColonPos+1 : getColonPos+3])
		for i := fellAsleepAt; i < wakeUpAt; i++ {
			sleepTime[currentGuard][i] += 1
		}
	}
	return sleepTime
}

func readAndSplitFile(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(input), "\n")
	return inputs[:len(inputs)-1]
}
