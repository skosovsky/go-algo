package main

import "testing"

func Test_closeCompare(t *testing.T) {
	type args struct {
		a      float64
		b      float64
		margin float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{4.0, 5.0, 0.0},
			want: -1,
		},
		{
			name: "Example 2",
			args: args{5.0, 5.0, 0.0},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{6.0, 5.0, 0.0},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{2.0, 5.0, 3.0},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{5.0, 5.0, 3.0},
			want: 0,
		},
		{
			name: "Example 6",
			args: args{8.0, 5.0, 3.0},
			want: 0,
		},
		{
			name: "Example 7",
			args: args{8.1, 5.0, 3.0},
			want: 1,
		},
		{
			name: "Example 8",
			args: args{1.99, 5.0, 3.0},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closeCompare(tt.args.a, tt.args.b, tt.args.margin); got != tt.want {
				t.Errorf("closeCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
