package main

import "testing"

func Test_combination(t *testing.T) {
	type args struct {
		numN int
		numK int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{3, 2},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{7, 5},
			want: 21,
		},
		{
			name: "Example 3",
			args: args{1, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combination(tt.args.numN, tt.args.numK); got != tt.want {
				t.Errorf("combination() = %v, want %v", got, tt.want)
			}
		})
	}
}
