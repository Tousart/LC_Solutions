package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func isSheet(a, b, s float64) int {
	sqrt := math.Pow(math.Pow(a+b, 2)-4*(a*b-s), 0.5)
	if math.Trunc(sqrt) != sqrt {
		return -1
	}
	numerator := int(a + b + sqrt)
	if numerator%2 != 0 {
		return -1
	}
	return numerator / 2
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	abs := strings.Split(strings.TrimSpace(line), " ")
	a, _ := strconv.ParseFloat(abs[0], 64)
	b, _ := strconv.ParseFloat(abs[1], 64)
	c, _ := strconv.ParseFloat(abs[2], 64)

	writer.WriteString(strconv.Itoa(isSheet(a, b, c)))
}
