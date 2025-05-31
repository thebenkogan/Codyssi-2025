package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
)

const (
	maxPrice = 15000000000000
)

func main() {
	input := lib.GetInput()

	ns := strings.Split(input, "\n\n")[1]
	nums := lib.ParseNums(ns)
	slices.Sort(nums)
	median := nums[len(nums)/2]

	fmt.Println(price(median))

	evens := lo.Filter(nums, func(n int, _ int) bool {
		return n%2 == 0
	})
	fmt.Println(price(lo.Sum(evens)))

	options := lo.Filter(nums, func(n int, _ int) bool {
		return price(n) < maxPrice
	})
	fmt.Println(lo.Max(options))
}

func price(n int) int {
	return a(b(c(n)))
}

func a(n int) int {
	return n + 957
}

func b(n int) int {
	return n * 71
}

func c(n int) int {
	return int(math.Pow(float64(n), 3))
}
