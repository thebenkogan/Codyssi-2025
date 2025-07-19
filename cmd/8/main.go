package main

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
)

func main() {
	input := lib.GetInput()

	fmt.Println(lo.CountBy([]rune(input), isAlphabetical))

	p2 := 0
	p3 := 0
	for _, line := range strings.Split(input, "\n") {
		p2 += len(reduce(line, p2Tester))
		p3 += len(reduce(line, p3Tester))
	}
	fmt.Println(p2)
	fmt.Println(p3)
}

func p2Tester(a, b rune) bool {
	return isNumber(a) && !isNumber(b) || !isNumber(a) && isNumber(b)
}

func p3Tester(a, b rune) bool {
	return isAlphabetical(a) && isNumber(b) || isNumber(a) && isAlphabetical(b)
}

func reduce(line string, tester func(a, b rune) bool) string {
	curr := []rune(line)
	for len(curr) > 0 {
		var next strings.Builder
		reduced := false
		i := 0
		for i < len(curr) {
			if i < len(curr)-1 && tester(curr[i], curr[i+1]) {
				reduced = true
				i += 2
				continue
			}
			next.WriteRune(curr[i])
			i += 1
		}

		curr = []rune(next.String())
		if !reduced {
			return string(curr)
		}
	}

	return ""
}

func isAlphabetical(r rune) bool {
	return r >= 97 && r <= 122
}

func isNumber(r rune) bool {
	return r >= 48 && r <= 57
}
