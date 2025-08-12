package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	var (
		i     int
		j     int
		count int = 1
	)

	for j < len(chars) {
		if j+1 == len(chars) || chars[j] != chars[j+1] {
			chars[i] = chars[j]

			if count > 1 {
				bytesCount := []byte(strconv.Itoa(count))
				for _, val := range bytesCount {
					i++
					chars[i] = val
				}
			}

			i++
			count = 1
		} else {
			count++
		}
		j++
	}

	return i
}

func main() {
	s := "abbccccdd"
	chars := []byte(s)
	fmt.Println(compress(chars))
}
