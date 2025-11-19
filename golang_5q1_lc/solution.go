package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	type studentsCount struct {
		ones   int
		zeroes int
	}

	stud := studentsCount{}

	for i := range len(students) {
		if students[i] == 1 {
			stud.ones++
		} else {
			stud.zeroes++
		}
	}

	idx := 0
	for idx < len(sandwiches) {
		if sandwiches[idx] == 1 {
			if stud.ones > 0 {
				stud.ones--
			} else {
				break
			}
		} else {
			if stud.zeroes > 0 {
				stud.zeroes--
			} else {
				break
			}
		}
		idx++
	}

	return len(sandwiches) - idx
}

func main() {
	st := []int{1, 1, 1, 0, 0, 1}
	sd := []int{1, 0, 0, 0, 1, 1}
	fmt.Println(countStudents(st, sd))
}
