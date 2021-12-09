package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2021/utils"
)

func (r *Runner) Day2Part1(in string) string {

	course := utils.StringsFromRegex(in, `^(forward|down|up) ([0-9])$`)
	h, d := 0, 0
	for _, command := range course {
		val := utils.MustAtoi(command[2])
		switch command[1] {
		case "forward":
			h += val
		case "down":
			d += val
		case "up":
			d -= val
		}
	}
	return strconv.Itoa(h * d)

}

func (r *Runner) Day2Part2(in string) string {

	course := utils.StringsFromRegex(in, `^(forward|down|up) ([0-9])$`)
	h, d, a := 0, 0, 0
	for _, command := range course {
		val := utils.MustAtoi(command[2])
		switch command[1] {
		case "forward":
			h += val
			d += (a * val)
		case "down":
			a += val
		case "up":
			a -= val
		}
	}
	return strconv.Itoa(h * d)

}
