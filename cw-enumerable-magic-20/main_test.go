package main

import (
	"reflect"
	"testing"
)

func Test_eachCons(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{[]int{1, 2, 3, 4}, 2},
			want: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			name: "Example 2",
			args: args{[]int{1, 2, 3, 4}, 3},
			want: [][]int{{1, 2, 3}, {2, 3, 4}},
		},
		{
			name: "Example 3",
			args: args{[]int{3, 5, 8, 13}, 1},
			want: [][]int{{3}, {5}, {8}, {13}},
		},
		{
			name: "Example 4",
			args: args{[]int{3, 5, 8, 13}, 2},
			want: [][]int{{3, 5}, {5, 8}, {8, 13}},
		},
		{
			name: "Example 5",
			args: args{[]int{3, 5, 8, 13}, 3},
			want: [][]int{{3, 5, 8}, {5, 8, 13}},
		},
		{
			name: "Example 6",
			args: args{[]int{}, 3},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eachCons(tt.args.arr, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eachCons() = %v, want %v", got, tt.want)
			}
		})
	}
}
