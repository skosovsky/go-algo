package main

import "testing"

func TestCapitalize(t *testing.T) {
	type args struct {
		st  string
		arr []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{"abcdef", []int{1, 2, 5}},
			want: "aBCdeF",
		},
		{
			name: "Example 2",
			args: args{"abcdef", []int{1, 2, 5, 100}},
			want: "aBCdeF",
		},
		{
			name: "Example 3",
			args: args{"codewars", []int{1, 3, 5, 50}},
			want: "cOdEwArs",
		},
		{
			name: "Example 4",
			args: args{"abRacaDabRA", []int{2, 6, 9, 10}},
			want: "abRacaDabRA",
		},
		{
			name: "Example 5",
			args: args{"codewarriors", []int{5}},
			want: "codewArriors",
		},
		{
			name: "Example 6",
			args: args{"indexinglessons", []int{0}},
			want: "Indexinglessons",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.args.st, tt.args.arr); got != tt.want {
				t.Errorf("Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
