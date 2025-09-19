package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Spreadsheet struct {
	sheet map[string]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{sheet: make(map[string]int)}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.sheet[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.sheet[cell] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	sliceFormula := strings.Split(formula[1:], "+")
	x, y := 0, 0
	xInFormula, yInFormula := sliceFormula[0], sliceFormula[1]

	if val, err := strconv.Atoi(xInFormula); err == nil {
		x = val
	} else {
		x = this.sheet[xInFormula]
	}

	if val, err := strconv.Atoi(yInFormula); err == nil {
		y = val
	} else {
		y = this.sheet[yInFormula]
	}

	return x + y
}

func main() {
	rows := 3
	obj := Constructor(rows)
	obj.SetCell("A1", 10)
	obj.ResetCell("B2")
	fmt.Println(obj.GetValue("=A1+16"))
}
