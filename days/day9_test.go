package days

import "testing"

func TestDay9Part1(t *testing.T) {

	testInputs := map[string]string{
		`2199943210
3987894921
9856789892
8767896789
9899965678`: "15",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(9, 1, in)
		if out != expectedOut {
			t.Errorf("day9 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay9Part2(t *testing.T) {

	testInputs := map[string]string{
		`2199943210
3987894921
9856789892
8767896789
9899965678`: "1134",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(9, 2, in)
		if out != expectedOut {
			t.Errorf("day9 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
