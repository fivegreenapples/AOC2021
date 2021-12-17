package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2021/utils"
)

type d5Segment struct {
	a utils.Coord
	b utils.Coord
}

func (s d5Segment) isHorizontal() bool {
	return s.a.Y == s.b.Y
}
func (s d5Segment) isVertical() bool {
	return s.a.X == s.b.X
}
func (s d5Segment) sorted() d5Segment {
	if s.a.X > s.b.X || s.a.Y > s.b.Y {
		return d5Segment{
			a: s.b,
			b: s.a,
		}
	}
	return s
}
func (s d5Segment) coordAt(n int) utils.Coord {
	x, y := s.a.X, s.a.Y
	if s.a.X != s.b.X {
		sign := s.b.X - s.a.X/utils.AbsInt(s.b.X-s.a.X)
		x += sign * n
	}
	if s.a.Y != s.b.Y {
		sign := s.b.Y - s.a.Y/utils.AbsInt(s.b.Y-s.a.Y)
		y += sign * n
	}
	return utils.Coord{
		X: x,
		Y: y,
	}
}

func (s d5Segment) isCoordOnLine(c utils.Coord) bool {
	if s.isVertical() {
		return c.X == s.a.X && c.Y >= s.a.Y && c.Y <= s.b.Y
	}
	if s.isHorizontal() {
		return c.Y == s.a.Y && c.X >= s.a.X && c.X <= s.b.X
	}
	return false
}

func (r *Runner) Day5Part1(in string) string {

	lines := utils.StringsFromRegex(in, `^([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)$`)

	allPts := map[utils.Coord]bool{}
	overlapPts := map[utils.Coord]bool{}
	for _, line := range lines {

		// Vertical line
		if line[1] == line[3] {
			x := utils.MustAtoi(line[1])
			y1, y2 := utils.MustAtoi(line[2]), utils.MustAtoi(line[4])
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				pt := utils.Coord{X: x, Y: y}
				if allPts[pt] {
					overlapPts[pt] = true
				} else {
					allPts[pt] = true
				}
			}
			continue
		}

		// Horizontal line
		if line[2] == line[4] {
			y := utils.MustAtoi(line[2])
			x1, x2 := utils.MustAtoi(line[1]), utils.MustAtoi(line[3])
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				pt := utils.Coord{X: x, Y: y}
				if allPts[pt] {
					overlapPts[pt] = true
				} else {
					allPts[pt] = true
				}
			}
			continue
		}
	}

	return strconv.Itoa(len(overlapPts))
}

func (r *Runner) Day5Part2(in string) string {

	lines := utils.StringsFromRegex(in, `^([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)$`)

	allPts := map[utils.Coord]bool{}
	overlapPts := map[utils.Coord]bool{}
	for _, line := range lines {

		// Vertical line
		if line[1] == line[3] {
			x := utils.MustAtoi(line[1])
			y1, y2 := utils.MustAtoi(line[2]), utils.MustAtoi(line[4])
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				pt := utils.Coord{X: x, Y: y}
				if allPts[pt] {
					overlapPts[pt] = true
				} else {
					allPts[pt] = true
				}
			}
			continue
		}

		// Horizontal line
		if line[2] == line[4] {
			y := utils.MustAtoi(line[2])
			x1, x2 := utils.MustAtoi(line[1]), utils.MustAtoi(line[3])
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				pt := utils.Coord{X: x, Y: y}
				if allPts[pt] {
					overlapPts[pt] = true
				} else {
					allPts[pt] = true
				}
			}
			continue
		}

		// Diagonal line
		x1, x2 := utils.MustAtoi(line[1]), utils.MustAtoi(line[3])
		y1, y2 := utils.MustAtoi(line[2]), utils.MustAtoi(line[4])
		yDelta := 1
		if x1 > x2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		if y1 > y2 {
			yDelta = -1
		}
		for x, y := x1, y1; x <= x2; x, y = x+1, y+yDelta {
			pt := utils.Coord{X: x, Y: y}
			if allPts[pt] {
				overlapPts[pt] = true
			} else {
				allPts[pt] = true
			}
		}

	}

	return strconv.Itoa(len(overlapPts))
}
