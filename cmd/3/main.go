package main

import (
	"fmt"
	"slices"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
)

func main() {
	input := lib.GetInput()

	nums := lib.ParseNums(input)
	nums = lo.Map(nums, func(n int, _ int) int {
		if n < 0 {
			return -n
		}
		return n
	})

	ranges := make([]Range, 0)
	for i := 0; i < len(nums); i += 2 {
		ranges = append(ranges, Range{nums[i], nums[i+1]})
	}

	total := lo.SumBy(ranges, func(r Range) int {
		return r.Size()
	})
	fmt.Println(total)

	total = 0
	for i := 0; i < len(ranges); i += 2 {
		total += collapsedSize(ranges[i], ranges[i+1])
	}
	fmt.Println(total)

	best := 0
	for i := 0; i < len(ranges)-2; i += 2 {
		best = max(collapsedSize(ranges[i], ranges[i+1], ranges[i+2], ranges[i+3]), best)
	}
	fmt.Println(best)
}

func collapsedSize(ranges ...Range) int {
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.Start - b.Start
	})

	collapsed := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		if ranges[i].Overlaps(collapsed[len(collapsed)-1]) {
			collapsed[len(collapsed)-1] = collapsed[len(collapsed)-1].Merge(ranges[i])
		} else {
			collapsed = append(collapsed, ranges[i])
		}
	}

	return lo.SumBy(collapsed, func(r Range) int {
		return r.Size()
	})
}

type Range struct {
	Start, End int
}

func (r Range) Size() int {
	return r.End - r.Start + 1
}

func (r Range) Overlaps(other Range) bool {
	return r.Start <= other.End && other.Start <= r.End
}

func (r Range) Merge(other Range) Range {
	return Range{
		Start: min(r.Start, other.Start),
		End:   max(r.End, other.End),
	}
}
