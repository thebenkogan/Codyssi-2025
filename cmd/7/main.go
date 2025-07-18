package main

import (
	"fmt"
	"slices"
	"strings"

	lib "github.com/thebenkogan/Codyssi-2025"
)

func main() {
	input := lib.GetInput()
	sections := strings.Split(input, "\n\n")

	p1Tracks := lib.ParseNums(sections[0])
	p2Tracks := slices.Clone(p1Tracks)
	p3Tracks := slices.Clone(p1Tracks)
	swaps := strings.Split(sections[1], "\n")
	test := lib.ParseNums(sections[2])[0] - 1
	for i, swap := range swaps {
		// p1, direct swap
		a, b := parseSwap(swap)
		p1Tracks[a], p1Tracks[b] = p1Tracks[b], p1Tracks[a]

		// p2, triple swap with next
		nextSwap := swaps[(i+1)%len(swaps)]
		c, _ := parseSwap(nextSwap)
		p2Tracks[b], p2Tracks[c], p2Tracks[a] = p2Tracks[a], p2Tracks[b], p2Tracks[c]

		// p3, swap with non-overlapping blocks
		low := min(a, b)
		high := max(a, b)
		startHigh := high
		for low < startHigh && high < len(p3Tracks) {
			p3Tracks[low], p3Tracks[high] = p3Tracks[high], p3Tracks[low]
			low += 1
			high += 1
		}
	}

	fmt.Println(p1Tracks[test])
	fmt.Println(p2Tracks[test])
	fmt.Println(p3Tracks[test])
}

func parseSwap(s string) (int, int) {
	ns := lib.ParseNums(s)
	return ns[0] - 1, -1*ns[1] - 1
}
