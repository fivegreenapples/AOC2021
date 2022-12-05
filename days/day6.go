package days

import (
	"fmt"
	"strconv"

	"github.com/fivegreenapples/AOC2021/utils"
)

func (r *Runner) Day6Part1(in string) string {
	return r.d6Do(in, 80)
}

func (r *Runner) Day6Part2(in string) string {
	return r.d6Do(in, 256)
}

func (r *Runner) d6Do(in string, days int) string {
	fish := utils.CsvToInts(in)

	fishByAge := map[int]int{}
	for _, age := range fish {
		fishByAge[age]++
	}

	var total int
	for d := 1; d <= days; d++ {
		total = 0
		numZeroAgeFish := fishByAge[0]
		for age := 1; age <= 8; age++ {
			fishByAge[age-1] = fishByAge[age]
			total += fishByAge[age-1]
		}
		fishByAge[6] += numZeroAgeFish
		fishByAge[8] = numZeroAgeFish
		total += numZeroAgeFish + numZeroAgeFish

		if r.verbose {
			fmt.Printf("After %2d days there are %d fish\n", d, total)
		}
	}

	return strconv.Itoa(total)
}
