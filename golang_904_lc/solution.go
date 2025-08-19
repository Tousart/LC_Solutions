package main

import "fmt"

func totalFruit(fruits []int) int {
	var (
		res   int
		start int
		end   int
		last1 int
		last2 int = -1
	)

	for end < len(fruits) {
		if fruits[end] == fruits[last1] {
			last1 = end
		} else if last2 == -1 || fruits[end] == fruits[last2] {
			last2 = end
		} else {
			res = max(res, end-start)
			if last1 < last2 {
				start = last1 + 1
			} else {
				start = last2 + 1
			}
			last1 = end
			last2 = end - 1
		}
		end++
	}

	return max(res, end-start)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fruits := []int{1, 2, 1}
	fmt.Println(totalFruit(fruits))
}

// func totalFruit(fruits []int) int {
// 	var (
// 		i   int
// 		res int
// 	)
// 	valInd := make(map[int]int)

// 	for j, fruit := range fruits {
// 		valInd[fruit] = j
// 		if len(valInd) > 2 {
// 			res = max(res, j-i)
// 			i = j - 1
// 			for _, val := range valInd {
// 				if val < i {
// 					i = val
// 				}
// 			}
// 			delete(valInd, fruits[i])
// 			i++
// 		}
// 		j++
// 	}

// 	return max(res, len(fruits)-i)
// }
