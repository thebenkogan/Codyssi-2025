package main

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
	"github.com/zyedidia/generic/heap"
)

func main() {
	input := lib.GetInput()

	grid := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, lib.ParseNums(line))
	}

	fmt.Println(bestRowOrCol(grid))
	fmt.Println(minPath(grid, 14, 14))
	fmt.Println(minPath(grid, len(grid[0])-1, len(grid)-1))
}

func bestRowOrCol(grid [][]int) int {
	best := int(1e9)
	for _, row := range grid {
		best = min(best, lo.Sum(row))
	}
	for col := 0; col < len(grid[0]); col++ {
		total := 0
		for _, row := range grid {
			total += row[col]
		}
		best = min(best, total)
	}
	return best
}

type node struct {
	pos  lib.Vector
	cost int
}

func minPath(grid [][]int, endX, endY int) int {
	q := heap.New(func(a, b node) bool { return a.cost < b.cost })
	q.Push(node{pos: lib.Vector{X: 0, Y: 0}, cost: grid[0][0]})
	seen := make(map[string]struct{})

	for q.Size() > 0 {
		n, _ := q.Pop()
		if n.pos.X == endX && n.pos.Y == endY {
			return n.cost
		}

		if _, ok := seen[n.pos.Hash()]; ok {
			continue
		}
		seen[n.pos.Hash()] = struct{}{}

		for _, dir := range []lib.Vector{{X: 0, Y: 1}, {X: 1, Y: 0}} {
			next := lib.Vector{X: n.pos.X + dir.X, Y: n.pos.Y + dir.Y}
			if next.X > endX || next.Y > endY {
				continue
			}
			q.Push(node{pos: next, cost: n.cost + grid[next.Y][next.X]})
		}
	}

	panic("unreachable")
}
