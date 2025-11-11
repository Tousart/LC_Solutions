package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxPower(k int, columns []int) (int, []int) {
	if k == len(columns) {
		return 1, []int{1}
	}

	prefix := make(map[int][4]int)
	for i := range k - 1 {
		prefix[i] = [4]int{0, 0, 0, 0}
	}
	prefix[k-1] = [4]int{powerOfTower(columns[:k]), 1, 0, 1}

	for i := k; i < len(columns); i++ {
		a := prefix[i-1][0]
		b := prefix[i-k][0] + powerOfTower(columns[i-k+1:i+1])

		if a >= b {
			prefix[i] = prefix[i-1]
		} else {
			prefixTower := prefix[i-k]
			prefix[i] = [4]int{b, prefixTower[1] + 1, prefixTower[3], i - k + 2}
		}
	}

	fmt.Println(prefix)

	resCount := prefix[len(columns)-1][1]
	resSlice := []int{}

	current := prefix[len(prefix)-1]
	for current[3] != 0 {
		resSlice = append(resSlice, current[3])
		current = prefix[current[2]+k-2]
	}

	i, j := 0, len(resSlice)-1
	for i < j {
		resSlice[i], resSlice[j] = resSlice[j], resSlice[i]
		i++
		j--
	}

	return resCount, resSlice
}

func powerOfTower(tower []int) int {
	minLength := 1001
	sum := 0
	for _, val := range tower {
		sum += val
		minLength = min(minLength, val)
	}
	return sum * minLength
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	nk := strings.Split(strings.TrimSpace(line1), " ")
	n, _ := strconv.Atoi(nk[0])
	k, _ := strconv.Atoi(nk[1])

	line2, _ := reader.ReadString('\n')
	inputColumns := strings.Split(strings.TrimSpace(line2), " ")
	columns := make([]int, n)
	for i := range n {
		val, _ := strconv.Atoi(inputColumns[i])
		columns[i] = val
	}

	count, nums := maxPower(k, columns)
	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')

	numsAnswer := make([]string, len(nums))
	for i := range len(nums) {
		numsAnswer[i] = strconv.Itoa(nums[i])
	}
	writer.WriteString(strings.Join(numsAnswer, " "))
}
