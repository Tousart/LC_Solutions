package main

import "fmt"

func numOfWays(n int) int {
	mod := 1000000007
	abc := 6
	aba := 6

	for i := 2; i <= n; i++ {
		newAbc := (2*abc + 2*aba) % mod
		newAba := (2*abc + 3*aba) % mod
		abc, aba = newAbc, newAba
	}

	return (abc + aba) % mod
}

func main() {
	n := 5000
	fmt.Println(numOfWays(n))
}
