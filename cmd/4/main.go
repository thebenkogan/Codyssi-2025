package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
)

func main() {
	input := lib.GetInput()
	lines := strings.Split(input, "\n")

	total := lo.SumBy(lines, func(line string) int {
		return size(line)
	})
	fmt.Println(total)

	total = lo.SumBy(lines, func(line string) int {
		return size(lossyCompress(line))
	})
	fmt.Println(total)

	total = lo.SumBy(lines, func(line string) int {
		return size(losslessCompress(line))
	})
	fmt.Println(total)
}

func lossyCompress(line string) string {
	newLen := len(line) / 10
	removed := len(line) - (2 * newLen)
	return line[:newLen] + strconv.Itoa(removed) + line[len(line)-newLen:]
}

func losslessCompress(line string) string {
	var s strings.Builder
	curr := rune(line[0])
	length := 1
	for _, c := range line[1:] {
		if c == curr {
			length++
		} else {
			s.WriteRune(curr)
			s.WriteString(strconv.Itoa(length))
			curr = c
			length = 1
		}
	}
	s.WriteRune(curr)
	s.WriteString(strconv.Itoa(length))
	return s.String()
}

func size(line string) int {
	total := 0
	for _, c := range line {
		if int(c) < 65 {
			total += int(c) - 48
		} else {
			total += int(c) - 64
		}
	}
	return total
}
