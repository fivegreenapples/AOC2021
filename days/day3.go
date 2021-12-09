package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2021/utils"
)

func (r *Runner) Day3Part1(in string) string {

	numbers := utils.Lines(in)
	length := len(numbers[0])
	totals := make([]int, length)

	for _, n := range numbers {
		for i := 0; i < length; i++ {
			if n[i] == '1' {
				totals[i]++
			}
		}
	}

	gamma, epsilon := 0, 0
	unit := 1
	for i := 0; i < length; i++ {
		if totals[length-i-1] > len(numbers)/2 {
			gamma += unit
		} else {
			epsilon += unit
		}
		unit *= 2
	}

	return strconv.Itoa(gamma * epsilon)
}

func (r *Runner) Day3Part2(in string) string {

	numbers := utils.Lines(in)

	var oxygenRatingCandidates = make([]string, len(numbers))
	var scrubberRatingCandidates = make([]string, len(numbers))
	copy(oxygenRatingCandidates, numbers)
	copy(scrubberRatingCandidates, numbers)

	idx := 0
	for len(oxygenRatingCandidates) > 1 || len(scrubberRatingCandidates) > 1 {
		var totalOnes int
		var target byte
		var newCandidates []string

		if len(oxygenRatingCandidates) > 1 {
			totalOnes = 0
			for _, n := range oxygenRatingCandidates {
				totalOnes += int(n[idx] - '0')
			}
			target = '0'
			if (2 * totalOnes) >= len(oxygenRatingCandidates) {
				target = '1'
			}
			for _, n := range oxygenRatingCandidates {
				if n[idx] == target {
					newCandidates = append(newCandidates, n)
				}
			}
			oxygenRatingCandidates = newCandidates
		}

		if len(scrubberRatingCandidates) > 1 {
			totalOnes = 0
			for _, n := range scrubberRatingCandidates {
				totalOnes += int(n[idx] - '0')
			}
			target = '1'
			if (2 * totalOnes) >= len(scrubberRatingCandidates) {
				target = '0'
			}
			newCandidates = nil
			for _, n := range scrubberRatingCandidates {
				if n[idx] == target {
					newCandidates = append(newCandidates, n)
				}
			}
			scrubberRatingCandidates = newCandidates
		}

		idx++
	}

	oxygenRating, _ := strconv.ParseInt(oxygenRatingCandidates[0], 2, 0)
	scrubberRating, _ := strconv.ParseInt(scrubberRatingCandidates[0], 2, 0)

	return strconv.FormatInt(oxygenRating*scrubberRating, 10)
}
