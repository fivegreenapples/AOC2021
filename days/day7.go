package days

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/fivegreenapples/AOC2021/utils"
)

func (r *Runner) Day7Part1(in string) string {

	positions := utils.CsvToInts(in)

	sort.Ints(positions)
	min := positions[0]
	max := positions[len(positions)-1]

	currentMinPos := 2982398472394
	for testPos := min; testPos <= max; testPos++ {
		sum := 0
		for _, crabPos := range positions {
			sum += utils.AbsInt(crabPos - testPos)
		}
		if r.verbose {
			fmt.Println("Sum for", testPos, "is", sum)
		}
		if sum < currentMinPos {
			currentMinPos = sum
		}
	}

	return strconv.Itoa(currentMinPos)
}

func (r *Runner) Day7Part2(in string) string {

	positions := utils.CsvToInts(in)

	sort.Ints(positions)
	min := positions[0]
	max := positions[len(positions)-1]

	currentMinPos := 2982398472394
	for testPos := min; testPos <= max; testPos++ {
		sum := 0
		for _, crabPos := range positions {
			delta := utils.AbsInt(crabPos - testPos)
			sum += ((1 + delta) * (delta)) / 2
		}
		if r.verbose {
			fmt.Println("Sum for", testPos, "is", sum)
		}
		if sum < currentMinPos {
			currentMinPos = sum
		}
	}

	return strconv.Itoa(currentMinPos)

}
