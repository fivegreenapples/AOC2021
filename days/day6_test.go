package days

import "testing"

func TestDay6Part1(t *testing.T) {

	testInputs := map[string]string{
		`3,4,3,1,2`: "5934",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(6, 1, in)
		if out != expectedOut {
			t.Errorf("day6 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay6Part2(t *testing.T) {

	testInputs := map[string]string{
		`3,4,3,1,2`: "26984457539",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(6, 2, in)
		if out != expectedOut {
			t.Errorf("day6 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
