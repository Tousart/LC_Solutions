package main

import (
	"fmt"
	"math"
)

func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}

	codes := make(map[uint32]struct{})
	codesCount := int(math.Pow(2, float64(k)))

	for i := 0; i <= len(s)-k; i++ {
		code := bitMask(s[i : i+k])
		if _, ok := codes[code]; !ok {
			codes[code] = struct{}{}
			if len(codes) == codesCount {
				return true
			}
		}
	}

	return len(codes) == codesCount
}

func bitMask(str string) uint32 {
	var code uint32

	for i, num := range str {
		code |= uint32(num-'0') << i
	}

	return code
}

func main() {
	s := "00110110"
	k := 2
	fmt.Println(hasAllCodes(s, k))
}
