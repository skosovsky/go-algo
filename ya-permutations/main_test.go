package main

import "testing"

func Test_permutation(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{2},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{1},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutation(tt.args.num); got != tt.want {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
