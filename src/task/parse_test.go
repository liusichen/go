package main

import "testing"

func TestCalculate(t *testing.T) {
	var tests = []struct {
		byteSent int64
		retTime  float64
		tcpRtt   int64
		want     float64
	}{
		{0,0.0,0,0.0},
		{3,0.0,0,0.0},
		{300000,0.0,100000,24.0},
		{300000,2.0,0,1.2},
		{300000,2.0,1000000,0.8},
		{300,2.0,1000000,0.0008},
	}

	for  _,test := range tests {
		got := Calculate(test.byteSent, test.retTime, test.tcpRtt)
		if got != test.want {
			t.Errorf("Calculate(%d,%f,%d) = %f",test.byteSent,test.retTime,test.tcpRtt,got)
		}
	}
}

