package main

import "testing"

func Test_isValidKitDates(t *testing.T) {
	type args struct {
		kitDates []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{[]int{10, 9, 2022}},
			want: true,
		},
		{
			name: "Example 2",
			args: args{[]int{21, 9, 2022}},
			want: true,
		},
		{
			name: "Example 3",
			args: args{[]int{29, 2, 2022}},
			want: false,
		},
		{
			name: "Example 4",
			args: args{[]int{31, 2, 2022}},
			want: false,
		},
		{
			name: "Example 5",
			args: args{[]int{29, 2, 2020}},
			want: true,
		},
		{
			name: "Example 6",
			args: args{[]int{29, 2, 2100}},
			want: false,
		},
		{
			name: "Example 7",
			args: args{[]int{31, 11, 1999}},
			want: false,
		},
		{
			name: "Example 8",
			args: args{[]int{31, 12, 1999}},
			want: true,
		},
		{
			name: "Example 9",
			args: args{[]int{29, 2, 2024}},
			want: true,
		},
		{
			name: "Example 10",
			args: args{[]int{29, 2, 2023}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidKitDates(tt.args.kitDates); got != tt.want {
				t.Errorf("isValidKitDates() = %v, want %v", got, tt.want)
			}
		})
	}
}
