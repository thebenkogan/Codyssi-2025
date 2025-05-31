package lib

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func GetInput() string {
	day := os.Args[1]
	var inputFileName string
	if len(os.Args) > 2 && os.Args[2] == "test" {
		inputFileName = "test"
	} else {
		inputFileName = "in"
	}
	pwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/cmd/%s/%s.txt", pwd, day, inputFileName)
	b, _ := os.ReadFile(path)
	return string(b)
}

var NumRegex = regexp.MustCompile(`-?\d+`)

func ParseNums(s string) []int {
	ss := NumRegex.FindAllString(s, -1)
	nums := make([]int, len(ss))
	for i, str := range ss {
		n, _ := strconv.Atoi(str)
		nums[i] = n
	}
	return nums
}

type Vector struct {
	X, Y int
}

func (v *Vector) Hash() string {
	return fmt.Sprintf("%d,%d", v.X, v.Y)
}

var DIRS = []Vector{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var ALL_DIRS = []Vector{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}
