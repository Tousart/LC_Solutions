package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	i, j, count, idx := 0, 0, 0, 0

	for j < len(chars) {
		if chars[i] == chars[j] {
			count++
			j++
		} else {
			chars[idx] = chars[i]
			idx++

			if count > 1 {
				numStr := strconv.Itoa(count)
				for _, numRune := range numStr {
					chars[idx] = byte(numRune)
					idx++
				}
			}

			i = j
			count = 0
		}
	}

	chars[idx] = chars[i]
	idx++
	if count > 1 {
		numStr := strconv.Itoa(count)
		for _, numRune := range numStr {
			chars[idx] = byte(numRune)
			idx++
		}
	}

	return idx
}

func main() {
	chars := []byte("aabbccc")
	fmt.Println(compress(chars))
}
