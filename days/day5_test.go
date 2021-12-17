package days

import "testing"

func TestDay5Part1(t *testing.T) {

	testInputs := map[string]string{
		`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`: "5",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(5, 1, in)
		if out != expectedOut {
			t.Errorf("day5 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay5Part2(t *testing.T) {

	testInputs := map[string]string{
		`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`: "12",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(5, 2, in)
		if out != expectedOut {
			t.Errorf("day5 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
