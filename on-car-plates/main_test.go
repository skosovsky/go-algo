package main

import "testing"

func Test_iterateCarPlates(t *testing.T) {
	type args struct {
		carPlates string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{"R48FAO00OOO0OOA99OKA99OK"},
			want: "R48FA O00OO O0OO A99OK A99OK ",
		},
		{
			name: "Example 2",
			args: args{"R48FAO00OOO0OOA99OKA99O"},
			want: "-",
		},
		{
			name: "Example 3",
			args: args{"A9PQ"},
			want: "A9PQ ",
		},
		{
			name: "Example 4",
			args: args{"A9PQA"},
			want: "-",
		},
		{
			name: "Example 5",
			args: args{"A99AAA99AAA99AAA99AA"},
			want: "A99AA A99AA A99AA A99AA ",
		},
		{
			name: "Example 6",
			args: args{"AP9QA"},
			want: "-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iterateCarPlates(tt.args.carPlates); got != tt.want {
				t.Errorf("iterateCarPlates() = %v, want %v", got, tt.want)
			}
		})
	}
}
