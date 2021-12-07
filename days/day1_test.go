package days

import "testing"

func TestDay1Part1(t *testing.T) {

	testInputs := map[string]string{
		`199
		200
		208
		210
		200
		207
		240
		269
		260
		263`: "7",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(1, 1, in)
		if out != expectedOut {
			t.Errorf("day1 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay1Part2(t *testing.T) {

	testInputs := map[string]string{
		`199
		200
		208
		210
		200
		207
		240
		269
		260
		263`: "5",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(1, 2, in)
		if out != expectedOut {
			t.Errorf("day1 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
