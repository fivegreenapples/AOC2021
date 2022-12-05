package days

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2021/utils"
)

func (r *Runner) Day8Part1(in string) string {

	data := utils.MultilineCsvToStrings(in, " | ")
	numOf1478 := 0
	for _, dets := range data {

		digits := utils.CsvToStringsSep(dets[1], " ")
		for _, d := range digits {
			if len(d) == 2 || len(d) == 4 || len(d) == 3 || len(d) == 7 {
				numOf1478++
			}
		}

	}

	return strconv.Itoa(numOf1478)
}

func (r *Runner) Day8Part2(in string) string {

	data := utils.MultilineCsvToStrings(in, " | ")
	total := 0
	for _, dets := range data {

		wires := utils.CsvToStringsSep(dets[0], " ")
		digits := utils.CsvToStringsSep(dets[1], " ")

		wiresToDigit := map[string]int{}
		sixes := []runeSet{}
		var setCF, setACF, setCBDF, setABCDEFG runeSet
		for _, w := range wires {
			sorted := r.d8SortString(w)

			switch len(w) {
			case 2:
				setCF = r.d8ToSet(sorted)
				wiresToDigit[sorted] = 1
			case 3:
				setACF = r.d8ToSet(sorted)
				wiresToDigit[sorted] = 7
			case 4:
				setCBDF = r.d8ToSet(sorted)
				wiresToDigit[sorted] = 4
			case 6:
				sixes = append(sixes, r.d8ToSet(sorted))
			case 7:
				setABCDEFG = r.d8ToSet(sorted)
				wiresToDigit[sorted] = 8
			}
		}

		setA := r.d8Diff(setACF, setCF)
		setBD := r.d8Diff(setCBDF, setCF)
		setEG := r.d8Diff(r.d8Diff(setABCDEFG, setCBDF), setA)

		setCDE := runeSet{}
		sixSegmentCounts := map[rune]int{}
		for _, thisSixSet := range sixes {
			for r := range thisSixSet {
				sixSegmentCounts[r]++
			}
		}
		for r, cnt := range sixSegmentCounts {
			if cnt == 2 {
				setCDE[r] = true
			}
		}

		setF := r.d8Diff(setCF, setCDE)
		setC := r.d8Diff(setCF, setF)

		setB := r.d8Diff(setBD, setCDE)
		setD := r.d8Diff(setBD, setB)

		setG := r.d8Diff(setEG, setCDE)
		setE := r.d8Diff(setEG, setG)

		setABCEFG := r.d8Diff(setABCDEFG, setD)                // 0
		setABDEFG := r.d8Diff(setABCDEFG, setC)                // 6
		setABCDFG := r.d8Diff(setABCDEFG, setE)                // 9
		setACDEG := r.d8Diff(r.d8Diff(setABCDEFG, setB), setF) // 2
		setACDFG := r.d8Diff(r.d8Diff(setABCDEFG, setB), setE) // 3
		setABDFG := r.d8Diff(r.d8Diff(setABCDEFG, setC), setE) // 5

		wiresToDigit[r.d8FromSet(setABCEFG)] = 0
		wiresToDigit[r.d8FromSet(setABDEFG)] = 6
		wiresToDigit[r.d8FromSet(setABCDFG)] = 9
		wiresToDigit[r.d8FromSet(setACDEG)] = 2
		wiresToDigit[r.d8FromSet(setACDFG)] = 3
		wiresToDigit[r.d8FromSet(setABDFG)] = 5

		thisNum := 0
		for idx, d := range digits {
			thisNum += int(math.Pow10(3-idx)) * wiresToDigit[r.d8SortString(d)]
		}

		total += thisNum
	}

	return strconv.Itoa(total)
}

type runeSet map[rune]bool

func (r *Runner) d8SortString(a string) string {
	chars := strings.Split(a, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
func (r *Runner) d8ToSet(a string) runeSet {
	set := runeSet{}
	for _, c := range a {
		set[c] = true
	}
	return set
}
func (r *Runner) d8CloneSet(a runeSet) runeSet {
	set := runeSet{}
	for c := range a {
		set[c] = true
	}
	return set
}
func (r *Runner) d8FromSet(a runeSet) string {
	runes := []string{}
	for c := range a {
		runes = append(runes, string(c))
	}
	sort.Strings(runes)
	return strings.Join(runes, "")
}
func (r *Runner) d8Diff(a, b runeSet) runeSet {
	acloned := r.d8CloneSet(a)
	for r := range b {
		delete(acloned, r)
	}
	return acloned
}
