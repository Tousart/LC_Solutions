package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	sort.Slice(events, func(i, j int) bool {
		t1, _ := strconv.Atoi(events[i][1])
		t2, _ := strconv.Atoi(events[j][1])

		if t1 == t2 {
			return events[i][0] == "OFFLINE"
		}

		return t1 < t2
	})

	mentions := make([]int, numberOfUsers)
	offline := make(map[int]int)

	for _, event := range events {
		status := event[0]
		timestamp, _ := strconv.Atoi(event[1])
		payload := event[2]

		for k, v := range offline {
			if v <= timestamp {
				delete(offline, k)
			}
		}

		if status == "MESSAGE" {
			if payload == "ALL" {
				for i := range numberOfUsers {
					mentions[i]++
				}
			} else if payload == "HERE" {
				for i := range numberOfUsers {
					if _, ok := offline[i]; !ok {
						mentions[i]++
					}
				}
			} else {
				users := strings.Split(payload, " ")
				for i := range users {
					id, _ := strconv.Atoi(users[i][2:])
					mentions[id]++
				}
			}

		} else {
			id, _ := strconv.Atoi(payload)
			offline[id] = 60 + timestamp
		}
	}

	return mentions
}

func main() {
	numberOfUsers := 2
	events := [][]string{
		{"MESSAGE", "10", "id1 id0"},
		{"OFFLINE", "11", "0"},
		{"MESSAGE", "71", "HERE"},
	}
	fmt.Println(countMentions(numberOfUsers, events))
}
