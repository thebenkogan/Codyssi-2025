package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/samber/lo/mutable"
	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func main() {
	input := lib.GetInput()

	lines := strings.Split(input, "\n")
	signs := strings.Split(lines[len(lines)-1], "")
	corrections := lo.Map(lines[:len(lines)-1], func(line string, _ int) int {
		n, _ := strconv.Atoi(line)
		return n
	})

	fmt.Println(getOffset(corrections, signs))

	mutable.Reverse(signs)
	fmt.Println(getOffset(corrections, signs))

	pairCorrections := make([]int, 0)
	for i := 0; i < len(corrections); i += 2 {
		pairCorrections = append(pairCorrections, corrections[i]*10+corrections[i+1])
	}
	fmt.Println(getOffset(pairCorrections, signs))
}

func getOffset(corrections []int, signs []string) int {
	offset := corrections[0]
	for _, c := range lo.Zip2(corrections[1:], signs) {
		correction, sign := c.A, c.B
		if sign == "+" {
			offset += correction
		} else {
			offset -= correction
		}
	}
	return offset
}
