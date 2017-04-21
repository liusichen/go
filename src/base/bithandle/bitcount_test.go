package bithandle

import "testing"

func TestcountBitint64(t *testing.T) {
	var bitCounttests = []struct {
		inputValue  int
		countResult int
	}{
		{0, 0},
		{3, 2},
		{127, 8},
	}

	for _, bitCountTest := range bitCounttests {
		gotCount := countBitint64(bitCountTest.inputValue)
		if gotCount != bitCountTest.countResult {
			t.Errorf("countBitint64(%d) = %d\n", bitCountTest.inputValue, gotCount)
		}
	}
}
