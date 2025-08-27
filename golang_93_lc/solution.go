package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := []string{}
	address := []string{}

	if len(s) > 12 {
		return res
	}

	makeComb(&res, &address, s)

	return res
}

func makeComb(res, address *[]string, s string) {
	if len(s) == 0 {
		if len(*address) == 4 {
			*res = append(*res, strings.Join(*address, "."))
		}
		return
	}

	for i := range len(s) {
		strDigit := s[:i+1]
		if len(strDigit) > 1 && strDigit[0] == '0' {
			return
		}

		digit, _ := strconv.Atoi(strDigit)
		if digit > 255 {
			return
		}

		*address = append(*address, strconv.Itoa(digit))
		makeComb(res, address, s[i+1:])
		*address = (*address)[:len(*address)-1]
	}
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))

	s1 := "0000"
	fmt.Println(restoreIpAddresses(s1))

	s2 := "101023"
	fmt.Println(restoreIpAddresses(s2))
}
