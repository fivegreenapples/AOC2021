package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2021/utils"
)

func (r *Runner) Day1Part1(in string) string {
	depths := utils.LinesAsInts(in)

	increases := 0
	for d := 1; d < len(depths); d++ {
		if depths[d] > depths[d-1] {
			increases++
		}
	}

	return strconv.Itoa(increases)
}

func (r *Runner) Day1Part2(in string) string {
	depths := utils.LinesAsInts(in)

	increases := 0
	for d := 3; d < len(depths); d++ {
		first := depths[d-3] + depths[d-2] + depths[d-1]
		second := depths[d-2] + depths[d-1] + depths[d]
		if second > first {
			increases++
		}
	}

	return strconv.Itoa(increases)
}
