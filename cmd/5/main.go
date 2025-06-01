package main

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
)

func main() {
	input := lib.GetInput()
	lines := strings.Split(input, "\n")
	origin := lib.Vector{X: 0, Y: 0}

	points := map[string]lib.Vector{}
	for _, line := range lines {
		nums := lib.ParseNums(line)
		v := lib.Vector{X: nums[0], Y: nums[1]}
		points[v.Hash()] = v
	}

	nearest := closest(origin, points)
	farthest := lo.MaxBy(lo.Values(points), func(v1, v2 lib.Vector) bool {
		return v1.ManhattanDist(origin) > v2.ManhattanDist(origin)
	})

	fmt.Println(origin.ManhattanDist(farthest) - origin.ManhattanDist(nearest))

	curr := origin
	total := 0
	length := 0
	for len(points) > 0 {
		next := closest(curr, points)
		d := curr.ManhattanDist(next)
		if length == 1 {
			fmt.Println(d)
		}
		total += d
		curr = next
		delete(points, next.Hash())
		length += 1
	}
	fmt.Println(total)
}

func closest(start lib.Vector, points map[string]lib.Vector) lib.Vector {
	return lo.MinBy(lo.Values(points), func(v1, v2 lib.Vector) bool {
		d1, d2 := start.ManhattanDist(v1), start.ManhattanDist(v2)
		if d1 != d2 {
			return d1 < d2
		} else if v1.X != v2.X {
			return v1.X < v2.X
		} else {
			return v1.Y < v2.Y
		}
	})
}
