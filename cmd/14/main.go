package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
)

type item struct {
	quality   int
	cost      int
	materials int
}

func main() {
	input := lib.GetInput()

	items := make([]item, 0)
	for _, line := range strings.Split(input, "\n") {
		ns := lib.ParseNums(line)
		items = append(items, item{
			quality:   ns[1],
			cost:      ns[2],
			materials: ns[3],
		})
	}

	slices.SortFunc(items, func(a, b item) int {
		if a.quality == b.quality {
			return b.cost - a.cost
		}
		return b.quality - a.quality
	})

	fmt.Println(lo.SumBy(items[:5], func(i item) int { return i.materials }))
	fmt.Println(knapsack(items, 30))
	fmt.Println(knapsack(items, 300))
}

func knapsack(items []item, totalCost int) int {
	// dp[i, w] = optimal total quality for the first i items with cost limit w
	dp := make([][]int, 0, len(items)+1)
	// mats[i, w] = least amount of unique materials for the optimal combination of the first items with cost limit w
	mats := make([][]int, 0, len(items)+1)
	for range len(items) + 1 {
		dp = append(dp, make([]int, totalCost+1))
		mats = append(mats, make([]int, totalCost+1))
	}

	// recurrence:
	// case 1 (include item i only if cost fits): dp[i, w] = dp[i-1, w-c] + q
	// case 2 (exclude item i): dp[i, w] = dp[i-1, w]

	for i := 1; i <= len(items); i += 1 {
		item := items[i-1]
		for w := 1; w <= totalCost; w += 1 {
			includeItem := item.cost <= w && dp[i-1][w-item.cost]+item.quality > dp[i-1][w]
			if includeItem {
				dp[i][w] = dp[i-1][w-item.cost] + item.quality
				mats[i][w] = mats[i-1][w-item.cost] + item.materials
			} else {
				dp[i][w] = dp[i-1][w]
				mats[i][w] = mats[i-1][w]
			}
		}
	}

	return dp[len(items)][totalCost] * mats[len(items)][totalCost]
}
