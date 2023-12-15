package main

import "testing"

func Test_calcCost(t *testing.T) {
	type args struct {
		tariffCost  int
		tariffSize  int
		mbCost      int
		trafficSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{100, 10, 12, 15},
			want: 160,
		},
		{
			name: "Example 2",
			args: args{100, 10, 12, 1},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcCost(tt.args.tariffCost, tt.args.tariffSize, tt.args.mbCost, tt.args.trafficSize); got != tt.want {
				t.Errorf("calcCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
