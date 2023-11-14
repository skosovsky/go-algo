package main

import "testing"

func Test_calcSizeTable(t *testing.T) {
	tests := []struct {
		a1 int
		b1 int
		a2 int
		b2 int
		a  int
		b  int
	}{
		{10, 2, 2, 10, 20, 2},
		{5, 7, 3, 2, 9, 5},
	}
	for _, tt := range tests {
		res1, res2 := calcSizeTable(tt.a1, tt.b1, tt.a2, tt.b2)
		if res1 != tt.a && res2 != tt.b {
			t.Errorf("calcSizeTable(%v, %v, %v, %v) = %v, %v, expected %v, %v",
				tt.a1, tt.b1, tt.a2, tt.b2, res1, res2, tt.a, tt.b)
		}
	}
}
