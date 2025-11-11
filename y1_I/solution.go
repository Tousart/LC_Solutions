package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func countMoves(vStart, hStart, vFinish, hFinish int) int {
	if vStart == vFinish && hStart == hFinish {
		return 0
	} else if vStart != vFinish && hStart != hFinish {
		return (int(math.Abs(float64(hStart-hFinish)))-1)*3 + 1 + (int(math.Abs(float64(vStart-vFinish)))-1)*3
	} else if vStart == vFinish {
		return (int(math.Abs(float64(hStart-hFinish))) - 1) * 3
	}
	return (int(math.Abs(float64(vStart-vFinish))) - 1) * 3
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	input := strings.Split(strings.TrimSpace(line1), " ")
	vStart, _ := strconv.Atoi(input[0])
	hStart, _ := strconv.Atoi(input[1])

	line2, _ := reader.ReadString('\n')
	input = strings.Split(strings.TrimSpace(line2), " ")
	vFinish, _ := strconv.Atoi(input[0])
	hFinish, _ := strconv.Atoi(input[1])

	writer.WriteString(strconv.Itoa(countMoves(vStart, hStart, vFinish, hFinish)))
}
