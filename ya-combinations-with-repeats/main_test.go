package main

import "testing"

func Test_combinationWithRepeats(t *testing.T) {
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
			args: args{1, 1},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{4, 3},
			want: 20,
		},
		{
			name: "Example 3",
			args: args{2, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationWithRepeats(tt.args.numN, tt.args.numK); got != tt.want {
				t.Errorf("combinationWithRepeats() = %v, want %v", got, tt.want)
			}
		})
	}
}
