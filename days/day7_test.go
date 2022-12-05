package days

import "testing"

func TestDay7Part1(t *testing.T) {

	testInputs := map[string]string{
		`16,1,2,0,4,2,7,1,2,14`: "37",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(7, 1, in)
		if out != expectedOut {
			t.Errorf("day7 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay7Part2(t *testing.T) {

	testInputs := map[string]string{
		`16,1,2,0,4,2,7,1,2,14`: "168",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(7, 2, in)
		if out != expectedOut {
			t.Errorf("day7 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
