package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
)

func main() {
	input := lib.GetInput()
	sections := strings.Split(input, "\n\n")

	p1 := make([][]int, 0)
	p2 := make([][]int, 0)
	p3 := make([][]int, 0)
	for _, line := range strings.Split(sections[0], "\n") {
		p1 = append(p1, lib.ParseNums(line))
		p2 = append(p2, lib.ParseNums(line))
		p3 = append(p3, lib.ParseNums(line))
	}

	p2Instructions := strings.Split(sections[1], "\n")
	p3Instructions := slices.Clone(p2Instructions)
	flowControls := strings.Split(sections[2], "\n")

	for _, instruction := range p2Instructions {
		applyInstruction(p1, instruction)
	}

	for _, control := range flowControls {
		p2Instructions = applyFlowControl(control, p2, p2Instructions)
	}

	for {
		done := false
		for _, control := range flowControls {
			p3Instructions = applyFlowControl(control, p3, p3Instructions)
			if len(p3Instructions) == 0 {
				done = true
				break
			}
		}
		if done {
			break
		}
	}

	fmt.Println(maxRowOrCol(p1))
	fmt.Println(maxRowOrCol(p2))
	fmt.Println(maxRowOrCol(p3))
}

func applyFlowControl(control string, grid [][]int, instructions []string) []string {
	switch control {
	case "TAKE":
		return instructions
	case "CYCLE":
		first := instructions[0]
		instructions = instructions[1:]
		return append(instructions, first)
	case "ACT":
		applyInstruction(grid, instructions[0])
		return instructions[1:]
	}
	panic("unknown control: " + control)
}

func applyInstruction(grid [][]int, instruction string) {
	command := strings.Split(instruction, " ")[0]
	switch command {
	case "SHIFT":
		shift(grid, instruction)
	case "MULTIPLY":
		applyOperation(grid, instruction, func(a, b int) int { return a * b })
	case "SUB":
		applyOperation(grid, instruction, func(a, b int) int { return a - b })
	case "ADD":
		applyOperation(grid, instruction, func(a, b int) int { return a + b })
	}
}

func maxRowOrCol(grid [][]int) int {
	best := 0
	for _, row := range grid {
		best = max(best, lo.Sum(row))
	}
	for col := 0; col < len(grid[0]); col++ {
		total := 0
		for _, row := range grid {
			total += row[col]
		}
		best = max(best, total)
	}
	return best
}

func applyOperation(grid [][]int, instruction string, op func(int, int) int) {
	ns := lib.ParseNums(instruction)
	rhs := ns[0]
	isAll := len(ns) == 1
	var pos int
	var isCol bool
	if !isAll {
		pos = ns[1] - 1
		isCol = strings.Contains(instruction, "COL")
	}
	for y, row := range grid {
		for x, n := range row {
			if isAll || (isCol && x == pos) || (!isCol && y == pos) {
				grid[y][x] = correctValue(op(n, rhs))
			}
		}
	}
}

func shift(grid [][]int, instruction string) {
	ns := lib.ParseNums(instruction)
	pos := ns[0] - 1
	steps := ns[1]
	isCol := strings.Contains(instruction, "COL")
	if !isCol {
		newRow := make([]int, len(grid[pos]))
		for i, n := range grid[pos] {
			newRow[(i+steps)%len(grid[pos])] = n
		}
		grid[pos] = newRow
	} else {
		newCol := make([]int, len(grid))
		for i, row := range grid {
			newCol[(i+steps)%len(grid)] = row[pos]
		}
		for i, row := range grid {
			row[pos] = newCol[i]
		}
	}
}

func correctValue(n int) int {
	return ((n % 1073741824) + 1073741824) % 1073741824
}
