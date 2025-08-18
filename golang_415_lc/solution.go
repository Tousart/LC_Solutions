package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		return addStrings(num2, num1)
	}

	var shift uint8
	res := make([]byte, len(num1)+1)
	res[0] = '1'
	ind := len(res) - 1

	i, j := len(num1)-1, len(num2)-1
	for j >= 0 {
		sum := num1[i] - '0' + num2[j] - '0' + shift
		res[ind] = sum%10 + '0'
		shift = sum / 10
		ind--
		j--
		i--
	}

	for i >= 0 {
		sum := num1[i] - '0' + shift
		res[ind] = sum%10 + '0'
		shift = sum / 10
		ind--
		i--
	}

	if shift == 0 {
		return string(res[1:])
	}

	return string(res)
}

func main() {
	num1 := "1"
	num2 := "9"
	fmt.Println(addStrings(num1, num2))
}
