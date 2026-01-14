package main

import "fmt"

func pushDominoes(dominoes string) string {
	dominoes = "L" + dominoes + "R"
	res := []byte(dominoes)

	points := make([][]int, 0)
	start := 0
	for i := range dominoes {
		if dominoes[i] == '.' {
			if i > 0 && dominoes[i-1] != '.' {
				start = i
				points = append(points, []int{start, start})
			}
		} else {
			if i > 0 && dominoes[i-1] == '.' {
				points[len(points)-1][1] = i - 1
			}
		}
	}

	for _, lr := range points {
		l, r := lr[0], lr[1]
		if res[l-1] == 'R' && res[r+1] == 'L' {
			for l < r {
				res[l] = 'R'
				res[r] = 'L'
				l++
				r--
			}
		} else if res[l-1] == 'R' {
			for l <= r {
				res[l] = 'R'
				l++
			}
		} else if res[r+1] == 'L' {
			for r >= l {
				res[r] = 'L'
				r--
			}
		}
	}

	return string(res[1 : len(res)-1])
}

func main() {
	dominoes := ".RR..L"
	fmt.Println(pushDominoes(dominoes))

	dominoes1 := "..R.."
	fmt.Println(pushDominoes(dominoes1))
}
