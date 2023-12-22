package main

import "testing"

func Test_summation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{2},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{1},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{8},
			want: 36,
		},
		{
			name: "Example 4",
			args: args{22},
			want: 253,
		},
		{
			name: "Example 5",
			args: args{100},
			want: 5050,
		},
		{
			name: "Example 6",
			args: args{213},
			want: 22791,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summation(tt.args.n); got != tt.want {
				t.Errorf("summation() = %v, want %v", got, tt.want)
			}
		})
	}
}
