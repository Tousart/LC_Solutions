package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		return addBinary(b, a)
	}

	res := make([]rune, len(b)+1)
	res[0] = '1'
	idx := len(res) - 1
	shift := 0

	i, j := len(a)-1, len(b)-1
	for i >= 0 && j >= 0 {
		sum := shift

		if a[i] == '1' {
			sum++
		}

		if b[j] == '1' {
			sum++
		}

		if sum%2 == 0 {
			res[idx] = '0'
		} else {
			res[idx] = '1'
		}

		shift = sum / 2

		idx--
		i--
		j--
	}

	for j >= 0 {
		sum := shift
		if b[j] == '1' {
			sum++
		}

		if sum%2 == 0 {
			res[idx] = '0'
		} else {
			res[idx] = '1'
		}

		shift = sum / 2

		idx--
		j--
	}

	if shift == 1 {
		return string(res)
	}

	return string(res[1:])
}

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}
