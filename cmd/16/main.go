package main

import (
	"fmt"
	"math/big"
	"slices"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
)

// the length of a row/column on one cube face
const size int = 80

type face struct {
	grid       [][]int
	absorption int
}

func (f face) updateCol(col, n int) {
	for y, row := range f.grid {
		f.grid[y][col-1] = correctValue(row[col-1] + n)
	}
}

func (f face) updateRow(row, n int) {
	for x, val := range f.grid[row-1] {
		f.grid[row-1][x] = correctValue(val + n)
	}
}

func (f face) rotatePos90() face {
	rotated := make([][]int, 0, len(f.grid))
	for col := range size {
		ns := lo.Map(f.grid, func(row []int, _ int) int { return row[col] })
		slices.Reverse(ns)
		rotated = append(rotated, ns)
	}
	return face{grid: rotated, absorption: f.absorption}
}

func (f face) rotateNeg90() face {
	rotated := make([][]int, 0, len(f.grid))
	for col := range size {
		ns := lo.Map(f.grid, func(row []int, _ int) int { return row[len(row)-col-1] })
		rotated = append(rotated, ns)
	}
	return face{grid: rotated, absorption: f.absorption}
}

func (f face) rotate180() face {
	rotated := make([][]int, 0, len(f.grid))
	for i := len(f.grid) - 1; i >= 0; i-- {
		row := slices.Clone(f.grid[i])
		slices.Reverse(row)
		rotated = append(rotated, row)
	}
	return face{grid: rotated, absorption: f.absorption}
}

func (f face) dominantSum() int {
	best := 0
	for _, row := range f.grid {
		best = max(best, lo.Sum(row))
	}
	for col := range size {
		ns := lo.Map(f.grid, func(row []int, _ int) int { return row[col] })
		best = max(best, lo.Sum(ns))
	}
	return best
}

// orientation represents all faces on a cube. All faces are oriented based on the
// top left corner, except the back face, which uses the bottom left corner. This
// means all vertical turns require no rotations, but horizontal turns require 180
// degree turns on the back face.
type orientation struct {
	front face
	back  face
	left  face
	right face
	down  face
	up    face
}

func initOrientation() *orientation {
	o := &orientation{}
	for range size {
		row := make([]int, 0, size)
		for range size {
			row = append(row, 1)
		}
		o.front.grid = append(o.front.grid, row)
		o.back.grid = append(o.back.grid, slices.Clone(row))
		o.left.grid = append(o.left.grid, slices.Clone(row))
		o.right.grid = append(o.right.grid, slices.Clone(row))
		o.down.grid = append(o.down.grid, slices.Clone(row))
		o.up.grid = append(o.up.grid, slices.Clone(row))
	}
	return o
}

func correctValue(n int) int {
	mod := n % 100
	if mod == 0 {
		mod = 100
	}
	return mod
}

func (o *orientation) updateFace(n int) {
	for y, row := range o.front.grid {
		for x, val := range row {
			o.front.grid[y][x] = correctValue(val + n)
		}
	}
	o.front.absorption += size * size * n
}

func (o *orientation) updateRow(row, n int, linear bool) {
	o.front.updateRow(row, n)
	o.front.absorption += size * n
	if linear {
		o.right.updateRow(row, n)
		o.left.updateRow(row, n)
		o.back.updateRow(size-row+1, n)
	}
}

func (o *orientation) updateCol(col, n int, linear bool) {
	o.front.updateCol(col, n)
	o.front.absorption += size * n
	if linear {
		o.up.updateCol(col, n)
		o.down.updateCol(col, n)
		o.back.updateCol(col, n)
	}
}

func (o *orientation) turnUp() {
	newFront := o.down
	newUp := o.front
	newBack := o.up
	newDown := o.back
	o.front = newFront
	o.up = newUp
	o.back = newBack
	o.down = newDown
	o.right = o.right.rotatePos90()
	o.left = o.left.rotateNeg90()
}

func (o *orientation) turnDown() {
	newFront := o.up
	newUp := o.back
	newBack := o.down
	newDown := o.front
	o.front = newFront
	o.up = newUp
	o.back = newBack
	o.down = newDown
	o.right = o.right.rotateNeg90()
	o.left = o.left.rotatePos90()
}

func (o *orientation) turnRight() {
	newFront := o.left
	newRight := o.front
	newLeft := o.back.rotate180()
	newBack := o.right.rotate180()
	o.front = newFront
	o.right = newRight
	o.left = newLeft
	o.back = newBack
	o.up = o.up.rotateNeg90()
	o.down = o.down.rotatePos90()
}

func (o *orientation) turnLeft() {
	newFront := o.right
	newRight := o.back.rotate180()
	newLeft := o.front
	newBack := o.left.rotate180()
	o.front = newFront
	o.right = newRight
	o.left = newLeft
	o.back = newBack
	o.up = o.up.rotatePos90()
	o.down = o.down.rotateNeg90()
}

func runInstructions(instructions []string, turns []string, linear bool) (int, *big.Int) {
	cube := initOrientation()
	for i, instruction := range instructions {
		ns := lib.ParseNums(instruction)
		switch strings.Split(instruction, " ")[0] {
		case "FACE":
			cube.updateFace(ns[0])
		case "ROW":
			cube.updateRow(ns[0], ns[1], linear)
		case "COL":
			cube.updateCol(ns[0], ns[1], linear)
		}
		if i < len(turns) {
			switch turns[i] {
			case "L":
				cube.turnRight()
			case "R":
				cube.turnLeft()
			case "D":
				cube.turnUp()
			case "U":
				cube.turnDown()
			}
		}
	}

	absorptions := []int{
		cube.front.absorption,
		cube.back.absorption,
		cube.right.absorption,
		cube.left.absorption,
		cube.up.absorption,
		cube.down.absorption,
	}
	slices.Sort(absorptions)
	absorptionProduct := absorptions[4] * absorptions[5]

	sums := []int{
		cube.front.dominantSum(),
		cube.back.dominantSum(),
		cube.right.dominantSum(),
		cube.left.dominantSum(),
		cube.up.dominantSum(),
		cube.down.dominantSum(),
	}
	product := big.NewInt(1)
	for _, n := range sums {
		product.Mul(product, big.NewInt(int64(n)))
	}
	return absorptionProduct, product
}

func main() {
	input := lib.GetInput()
	sections := strings.Split(input, "\n\n")

	instructions := strings.Split(sections[0], "\n")
	turns := strings.Split(sections[1], "")

	absorptionProduct, sumProduct := runInstructions(instructions, turns, false)
	fmt.Println(absorptionProduct)
	fmt.Println(sumProduct)

	_, sumProduct = runInstructions(instructions, turns, true)
	fmt.Println(sumProduct)
}
