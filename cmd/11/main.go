package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/samber/lo/mutable"
	lib "github.com/thebenkogan/Codyssi-2025"
)

func main() {
	input := lib.GetInput()

	best := 0
	total := 0
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		n := []rune(split[0])
		base, _ := strconv.Atoi(split[1])
		var number int
		for i := range n {
			number += int(math.Pow(float64(base), float64(i)) * float64(runeToInt(n[len(n)-i-1])))
		}
		best = max(best, number)
		total += number
	}
	fmt.Println(best)

	var base68 strings.Builder
	running := total
	for running > 0 {
		base68.WriteRune(intToRune(running % 68))
		running /= 68
	}
	s := []rune(base68.String())
	mutable.Reverse(s)
	fmt.Println(string(s))

	l, r := 0, 10_000
	for l < r {
		base := (l + r) / 2
		// max = (base-1)*(base^3) + (base-1)*(base^2) + (base-1)*base + (base-1)
		// max = (base-1)*(base^3 + base^2 + base + 1)
		maxValue := (base - 1) * (base*base*base + base*base + base + 1)
		if maxValue > total {
			r = base
		} else {
			l = base + 1
		}
	}
	fmt.Println(l)
}

func runeToInt(r rune) int {
	// digits
	if r >= 48 && r <= 57 {
		return int(r) - 48
	}
	// uppercase
	if r >= 65 && r <= 90 {
		return int(r) - 55
	}
	// lowercase
	if r >= 97 && r <= 122 {
		return int(r) - 61
	}
	panic("invalid rune: " + string(r))
}

func intToRune(i int) rune {
	if i < 10 {
		return rune(i + 48)
	}
	if i < 36 {
		return rune(i + 55)
	}
	if i < 62 {
		return rune(i + 61)
	}
	switch i {
	case 62:
		return '!'
	case 63:
		return '@'
	case 64:
		return '#'
	case 65:
		return '$'
	case 66:
		return '%'
	case 67:
		return '^'
	}
	panic("invalid int: " + strconv.Itoa(i))
}
