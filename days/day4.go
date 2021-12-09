package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2021/utils"
)

func (r *Runner) Day4Part1(in string) string {

	inputLines := utils.Lines(in)

	drawnNumbers := utils.CsvToInts(inputLines[0])

	boards := [][]int{}
	for idx := 2; idx < len(inputLines); idx += 6 {
		var thisBoard []int
		for _, line := range inputLines[idx : idx+5] {
			thisBoard = append(thisBoard, utils.CsvToIntsSep(line, " ")...)
		}
		boards = append(boards, thisBoard)
	}

	var score int
	for _, drawn := range drawnNumbers {

		for boardIdx := range boards {
			for digitIdx := range boards[boardIdx] {
				if boards[boardIdx][digitIdx] == drawn {
					boards[boardIdx][digitIdx] = -1
				}
			}

			if hasBoardWon(boards[boardIdx]) {
				sumUnmarked := sumBoard(boards[boardIdx])
				score = drawn * sumUnmarked
				break
			}
		}

		if score > 0 {
			break
		}

	}

	return strconv.Itoa(score)
}

func (r *Runner) Day4Part2(in string) string {

	inputLines := utils.Lines(in)

	drawnNumbers := utils.CsvToInts(inputLines[0])

	boards := [][]int{}
	for idx := 2; idx < len(inputLines); idx += 6 {
		var thisBoard []int
		for _, line := range inputLines[idx : idx+5] {
			thisBoard = append(thisBoard, utils.CsvToIntsSep(line, " ")...)
		}
		boards = append(boards, thisBoard)
	}

	var score int
	for _, drawn := range drawnNumbers {

		var boardsNotYetWon [][]int
		for boardIdx := range boards {
			for digitIdx := range boards[boardIdx] {
				if boards[boardIdx][digitIdx] == drawn {
					boards[boardIdx][digitIdx] = -1
				}
			}

			if !hasBoardWon(boards[boardIdx]) {
				boardsNotYetWon = append(boardsNotYetWon, boards[boardIdx])
			} else {
				if len(boards) == 1 {
					// on last board
					sumUnmarked := sumBoard(boards[0])
					score = drawn * sumUnmarked
					break
				}
			}

		}
		boards = boardsNotYetWon
		if score > 0 {
			break
		}
	}

	return strconv.Itoa(score)
}

func hasBoardWon(b []int) bool {
	for r := 0; r < 25; r += 5 {
		if b[r]+b[r+1]+b[r+2]+b[r+3]+b[r+4] == -5 {
			return true
		}
	}
	for c := 0; c < 5; c++ {
		if b[c]+b[c+5]+b[c+10]+b[c+15]+b[c+20] == -5 {
			return true
		}
	}
	return false
}

func sumBoard(b []int) (sum int) {
	for i := 0; i < 25; i++ {
		if b[i] > 0 {
			sum += b[i]
		}
	}
	return
}
