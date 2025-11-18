package main

import (
	"fmt"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	timeStack := make([]int, 0)
	badTimeStack := make([]int, 0)

	for _, log := range logs {
		logSlice := strings.Split(log, ":")
		proc, _ := strconv.Atoi(logSlice[0])
		status := logSlice[1]
		time, _ := strconv.Atoi(logSlice[2])

		if status == "start" {
			timeStack = append(timeStack, time)
			badTimeStack = append(badTimeStack, 0)
			continue
		}

		badTime := badTimeStack[len(badTimeStack)-1]
		currTime := time - timeStack[len(timeStack)-1] + 1 - badTime
		timeStack = timeStack[:len(timeStack)-1]
		badTimeStack = badTimeStack[:len(badTimeStack)-1]

		res[proc] += currTime
		if len(badTimeStack) > 0 {
			badTimeStack[len(badTimeStack)-1] += currTime + badTime
		}
	}

	return res
}

func main() {
	n := 1
	logs := []string{"0:start:0", "0:start:1", "0:start:2", "0:end:3", "0:end:4", "0:end:5"}
	fmt.Println(exclusiveTime(n, logs))
}
