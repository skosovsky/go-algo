package main

import (
	"reflect"
	"testing"
)

func Test_mergeArrays(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "Example 2",
			args: args{[]int{1, 3, 5, 7, 9}, []int{10, 8, 6, 4, 2}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "Example 3",
			args: args{[]int{1, 3, 5, 7, 9, 11, 12}, []int{1, 2, 3, 4, 5, 10, 12}},
			want: []int{1, 2, 3, 4, 5, 7, 9, 10, 11, 12},
		},
		{
			name: "Example 4",
			args: args{[]int{}, []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeArrays(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
