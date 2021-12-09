package days

import "testing"

func TestDay3Part1(t *testing.T) {

	testInputs := map[string]string{
		`00100
		11110
		10110
		10111
		10101
		01111
		00111
		11100
		10000
		11001
		00010
		01010`: "198",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(3, 1, in)
		if out != expectedOut {
			t.Errorf("day3 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay3Part2(t *testing.T) {

	testInputs := map[string]string{
		`00100
		11110
		10110
		10111
		10101
		01111
		00111
		11100
		10000
		11001
		00010
		01010`: "230",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(3, 2, in)
		if out != expectedOut {
			t.Errorf("day3 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
