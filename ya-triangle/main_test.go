package main

import "testing"

func Test_isPossibleTriangle(t *testing.T) {
	type args struct {
		side1 int
		side2 int
		side3 int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{3, 4, 5},
			want: "YES",
		},
		{
			name: "Example 2",
			args: args{3, 5, 4},
			want: "YES",
		},
		{
			name: "Example 3",
			args: args{3, 5, 4},
			want: "YES",
		},
		{
			name: "Example 4",
			args: args{1, 4, 3},
			want: "NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossibleTriangle(tt.args.side1, tt.args.side2, tt.args.side3); got != tt.want {
				t.Errorf("isPossibleTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
