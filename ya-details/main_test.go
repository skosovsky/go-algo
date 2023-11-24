package main

import "testing"

func Test_countM(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		m      int
		mCount int
	}{
		{10, 5, 2, 4},
		{13, 5, 3, 3},
		{14, 5, 3, 4},
	}
	for _, tt := range tests {
		res := countM(tt.n, tt.k, tt.m)
		if res != tt.mCount {
			t.Errorf("countM(%v, %v, %v) = %v, expected %v",
				tt.n, tt.k, tt.m, tt.mCount, res)
		}
	}
}
