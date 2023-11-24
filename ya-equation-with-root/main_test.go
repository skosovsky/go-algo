package main

import (
	"reflect"
	"testing"
)

func Test_solveEquation(t *testing.T) {
	type args struct {
		numA int
		numB int
		numC int
	}
	tests := []struct {
		name       string
		args       args
		wantResult interface{} //nolint:gofmt
	}{
		{
			name:       "Example 1",
			args:       args{1, 0, 0},
			wantResult: 0,
		},
		{
			name:       "Example 2",
			args:       args{1, 2, 3},
			wantResult: 7,
		},
		{
			name:       "Example 3",
			args:       args{1, 2, -3},
			wantResult: "NO SOLUTION",
		},
		{
			name:       "Example 4",
			args:       args{0, 0, 0},
			wantResult: "MANY SOLUTIONS",
		},
		{
			name:       "Example 5",
			args:       args{3, 13, 7},
			wantResult: 12,
		},
		{
			name:       "Example 6",
			args:       args{0, 0, 5},
			wantResult: "NO SOLUTION",
		},
		{
			name:       "Example 7",
			args:       args{0, 4, 2},
			wantResult: "MANY SOLUTIONS",
		},
		{
			name:       "Example 8",
			args:       args{0, 2, 2},
			wantResult: "NO SOLUTION",
		},
		{
			name:       "Example 9",
			args:       args{2, 2, 3},
			wantResult: "NO SOLUTION",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solveEquation(tt.args.numA, tt.args.numB, tt.args.numC); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("solveEquation() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
