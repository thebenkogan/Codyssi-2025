package main

import (
	"fmt"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
)

func main() {
	input := lib.GetInput()

	total := lo.CountBy([]rune(input), func(c rune) bool {
		return lo.Contains(lo.LettersCharset, c)
	})
	fmt.Println(total)

	total = lo.SumBy([]rune(input), func(c rune) int {
		return value(c)
	})
	fmt.Println(total)

	corrected := make([]int, 0, len(input))
	for i, c := range input {
		v := value(c)
		if v > 0 {
			corrected = append(corrected, v)
			continue
		}
		prev := corrected[i-1]
		result := prev*2 - 5
		result = result % 52
		if result <= 0 || result == 52 {
			result += 52
		}
		corrected = append(corrected, result)
	}
	fmt.Println(lo.Sum(corrected))
}

func value(c rune) int {
	ascii := int(c)
	if ascii >= 65 && ascii <= 90 {
		return ascii - 64 + 26
	}
	if ascii >= 97 && ascii <= 122 {
		return ascii - 96
	}
	return 0
}
