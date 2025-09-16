package main

import "fmt"

func replaceNonCoprimes(nums []int) []int {
	res := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		res = append(res, nums[i])
		for len(res) >= 2 {
			nod := evklid(res[len(res)-1], res[len(res)-2])
			if nod != 1 {
				res[len(res)-2] = (res[len(res)-1] * res[len(res)-2]) / nod
				res = res[:len(res)-1]
			} else {
				break
			}
		}
	}
	return res
}

func evklid(a, b int) int {
	if a < b {
		return evklid(b, a)
	}
	if b == 0 {
		return a
	}
	return evklid(b, a%b)
}

func main() {
	a := []int{287, 41, 49, 287, 899, 23, 23, 20677, 5, 825}
	fmt.Println(replaceNonCoprimes(a))
}
