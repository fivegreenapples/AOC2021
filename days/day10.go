package days

import (
	"sort"
	"strconv"
	"strings"
)

func (r *Runner) Day10Part1(in string) string {

	lines := strings.Split(in, "\n")
	errorScore := 0
	for _, line := range lines {
		expectedClosers := []rune{}
		for _, char := range line {
			switch char {
			case '[':
				expectedClosers = append(expectedClosers, ']')
			case '(':
				expectedClosers = append(expectedClosers, ')')
			case '{':
				expectedClosers = append(expectedClosers, '}')
			case '<':
				expectedClosers = append(expectedClosers, '>')
			case ']', ')', '}', '>':
				if expectedClosers[len(expectedClosers)-1] != char {
					// oops
					errorScore += map[rune]int{']': 57, ')': 3, '}': 1197, '>': 25137}[char]
					goto nextline
				} else {
					expectedClosers = expectedClosers[:len(expectedClosers)-1]
				}
			}
		}
	nextline:
	}

	return strconv.Itoa(errorScore)
}

func (r *Runner) Day10Part2(in string) string {
	lines := strings.Split(in, "\n")
	completionScores := []int{}
	for _, line := range lines {
		expectedClosers := []rune{}
		thisScore := 0
		for _, char := range line {
			switch char {
			case '[':
				expectedClosers = append(expectedClosers, ']')
			case '(':
				expectedClosers = append(expectedClosers, ')')
			case '{':
				expectedClosers = append(expectedClosers, '}')
			case '<':
				expectedClosers = append(expectedClosers, '>')
			case ']', ')', '}', '>':
				if expectedClosers[len(expectedClosers)-1] != char {
					// oops
					goto nextline
				} else {
					expectedClosers = expectedClosers[:len(expectedClosers)-1]
				}
			}
		}

		for idx := len(expectedClosers) - 1; idx >= 0; idx-- {
			thisScore = thisScore*5 + map[rune]int{']': 2, ')': 1, '}': 3, '>': 4}[expectedClosers[idx]]
		}
		completionScores = append(completionScores, thisScore)

	nextline:
	}

	sort.Ints(completionScores)

	return strconv.Itoa(completionScores[len(completionScores)/2])
}
