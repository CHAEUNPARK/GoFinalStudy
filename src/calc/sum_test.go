package calc

import "testing"

type TestData struct {
	argument1 int
	argument2 int
	result    int
}

var testdata = []TestData{
	{2, 6, 8},
	{-8, 3, -5},
	{-6, 6, 0},
	{0, 0, 0},
}

func TestSum(t *testing.T) {
	for _, d := range testdata {
		r := Sum(d.argument1, d.argument2)
		if r != d.result {
			t.Errorf(
				"%d + %d의 결과값이 %d(이)가 아닙니다. r=%d",
				d.argument1,
				d.argument2,
				d.result,
				r,
			)
		}
	}
}
