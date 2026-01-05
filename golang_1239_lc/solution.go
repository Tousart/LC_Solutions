package main

import (
	"fmt"
	"math/bits"
)

func maxLength(arr []string) int {
	var res int
	bitArr := makeBitArr(arr)
	concString(bitArr, 0, uint32(0), &res)
	return res
}

func concString(bitArr []uint32, idx int, seq uint32, res *int) {
	for i := idx; i < len(bitArr); i++ {
		if (seq & bitArr[i]) == 0 {
			*res = max(*res, bits.OnesCount32(seq|bitArr[i]))
			concString(bitArr, i+1, seq|bitArr[i], res)
		}
	}
}

func makeBitArr(arr []string) []uint32 {
	bitArr := make([]uint32, len(arr))
	for i, word := range arr {
		num := uint32(0)
		j := 0
		for j < len(word) {
			bit := uint32(1) << (rune(word[j]) - 'a')
			if bit&num == 0 {
				num |= bit
			} else {
				break
			}
			j++
		}
		if j == len(word) {
			bitArr[i] = num
		}
	}
	return bitArr
}

func main() {
	arr := []string{"un", "iq", "ue"}
	fmt.Println(maxLength(arr))

	arr1 := []string{"aa", "bb"}
	fmt.Println(maxLength(arr1))
}
