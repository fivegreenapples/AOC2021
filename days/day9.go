package days

import (
	"sort"
	"strconv"

	"github.com/fivegreenapples/AOC2021/utils"
)

func (r *Runner) Day9Part1(in string) string {

	heights := utils.MultilineCsvToInts(in, "")
	heightMap := map[utils.Coord]int{}
	for y, row := range heights {
		for x, h := range row {
			heightMap[utils.Coord{X: x, Y: y}] = h
		}
	}

	min, max := utils.ExtentsOfIntMap(heightMap)
	lowPointRisk := 0
	utils.Foreach2D(min, max, func(c utils.Coord) {

		current := heightMap[c]
		adjacents := c.OrthAdjacents()
		anyLower := false
		for _, adj := range adjacents {
			if hAdj, ok := heightMap[adj]; ok {
				if hAdj <= current {
					anyLower = true
					break
				}
			}
		}

		if !anyLower {
			lowPointRisk += (1 + current)
		}

	})

	return strconv.Itoa(lowPointRisk)
}

func (r *Runner) Day9Part2(in string) string {

	heights := utils.MultilineCsvToInts(in, "")
	heightMap := map[utils.Coord]int{}
	for y, row := range heights {
		for x, h := range row {
			heightMap[utils.Coord{X: x, Y: y}] = h
		}
	}

	min, max := utils.ExtentsOfIntMap(heightMap)
	basins := []int{}
	utils.Foreach2D(min, max, func(c utils.Coord) {

		current := heightMap[c]
		if current == 9 || current == -1 {
			// not in basin or already seend
			return
		}

		// start depth search for all points in basin
		basinSize := 1
		heightMap[c] = -1
		candidates := c.OrthAdjacents()
		candidateIdx := 0
		for candidateIdx < len(candidates) {

			thisCandidate := candidates[candidateIdx]
			candidateIdx++

			if hCand, ok := heightMap[thisCandidate]; ok {

				if hCand == 9 || hCand == -1 {
					continue
				}

				heightMap[thisCandidate] = -1
				basinSize++
				candidates = append(candidates, thisCandidate.OrthAdjacents()...)
			}
		}

		basins = append(basins, basinSize)
	})

	sort.Ints(basins)
	return strconv.Itoa(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
}
