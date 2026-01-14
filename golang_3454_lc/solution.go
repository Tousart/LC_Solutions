package main

import (
	"fmt"
	"math"
	"sort"
)

type Interval struct {
	x1 int
	x2 int
}

type Event struct {
	y   int
	del bool
	x1  int
	x2  int
}

type Area struct {
	prevY float64
	heigh float64
	width float64
}

func separateSquares(squares [][]int) float64 {
	events := make([]Event, 0)

	for _, square := range squares {
		events = append(events,
			Event{
				y:  square[1],
				x1: square[0],
				x2: square[0] + square[2],
			},
			Event{
				y:   square[1] + square[2],
				del: true,
				x1:  square[0],
				x2:  square[0] + square[2],
			},
		)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	areas := make([]Area, 0)
	prevY := events[0].y
	totalArea := 0
	intervals := make([]Interval, 0)
	delIntervals := make([]Interval, 0)

	for _, event := range events {
		if event.y > prevY && len(intervals) > 0 {
			intervals = deleteIntervals(intervals, delIntervals)

			heigh := event.y - prevY
			width := makeWidth(intervals)

			areas = append(areas, Area{
				prevY: float64(prevY),
				heigh: float64(heigh),
				width: float64(width),
			})
			totalArea += heigh * width

			delIntervals = make([]Interval, 0)
		}

		if event.del {
			delIntervals = append(delIntervals, Interval{
				x1: event.x1,
				x2: event.x2,
			})
		} else {
			intervals = append(intervals, Interval{
				x1: event.x1,
				x2: event.x2,
			})
		}

		prevY = event.y
	}

	target := float64(totalArea) / 2.0
	prevArea := 0.0

	for _, area := range areas {
		if (prevArea + area.heigh*area.width) >= target {
			return area.prevY + (target-prevArea)/area.width
		}
		prevArea += area.heigh * area.width
	}

	return 0.0
}

func deleteIntervals(intervals, delIntervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].x1 == intervals[j].x1 {
			return intervals[i].x2 < intervals[j].x2
		}
		return intervals[i].x1 < intervals[j].x1
	})

	sort.Slice(delIntervals, func(i, j int) bool {
		if delIntervals[i].x1 == delIntervals[j].x1 {
			return delIntervals[i].x2 < delIntervals[j].x2
		}
		return delIntervals[i].x1 < delIntervals[j].x1
	})

	res := make([]Interval, 0)
	j := 0
	for i := range intervals {
		if j < len(delIntervals) && delIntervals[j] == intervals[i] {
			j++
			continue
		}
		res = append(res, intervals[i])
	}

	return res
}

func makeWidth(intervals []Interval) int {
	res := 0
	end := math.MinInt64

	for _, interval := range intervals {
		if interval.x1 > end {
			res += interval.x2 - interval.x1
			end = interval.x2
		} else if interval.x2 > end {
			res += interval.x2 - end
			end = interval.x2
		}
	}

	return res
}

func main() {
	squares := [][]int{
		{0, 0, 1},
		{2, 2, 1},
	}
	fmt.Println(separateSquares(squares))

	squares1 := [][]int{
		{3, 16, 5},
		{21, 26, 3},
		{21, 24, 3},
		{0, 1, 4},
	}
	fmt.Println(separateSquares(squares1))

	squares2 := [][]int{
		{19, 30, 5},
		{19, 22, 5},
		{4, 21, 3},
	}
	fmt.Println(separateSquares(squares2))
}
