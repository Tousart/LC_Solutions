package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	res := make([]int, rowIndex+1)
	mid := 1

	for i := 1; i <= rowIndex; i++ {
		idx := len(res)/2 - 1

		if i%2 != 0 {
			prev := mid

			for idx >= 0 {
				newVal := prev + res[idx]
				prev = res[idx]
				res[idx] = newVal
				idx--
			}
		} else {
			mid = res[idx] * 2

			for idx > 0 {
				if idx > 0 {
					res[idx] += res[idx-1]
				}
				idx--
			}
		}
	}

	SecondPartidx := len(res) / 2

	if rowIndex%2 == 0 {
		res[SecondPartidx] = mid
		SecondPartidx++
	}

	for i := SecondPartidx; i < len(res); i++ {
		res[i] = res[len(res)-i-1]
	}

	return res
}

func main() {
	rowIndex := 4
	fmt.Println(getRow(rowIndex))
}
