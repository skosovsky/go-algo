package main

import "testing"

func Test_checkMonotopy(t *testing.T) {
	type args struct {
		userSlice []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name:       "Example 1",
			args:       args{[]int{1, 2, 3, 4}},
			wantResult: true,
		},
		{
			name:       "Example 2",
			args:       args{[]int{1, 1}},
			wantResult: true,
		},
		{
			name:       "Example 3",
			args:       args{[]int{9, 5, 1}},
			wantResult: true,
		},
		{
			name:       "Example 4",
			args:       args{[]int{9, 5, 10}},
			wantResult: false,
		},
		{
			name:       "Example 5",
			args:       args{[]int{9}},
			wantResult: true,
		},
		{
			name:       "Example 6",
			args:       args{[]int{9, 5}},
			wantResult: true,
		},
		{
			name:       "Example 7",
			args:       args{[]int{5, 6, 7, 9, 11, 12}},
			wantResult: true,
		},
		{
			name:       "Example 8",
			args:       args{[]int{5, 6, 7, 6, 11, 12}},
			wantResult: false,
		},
		{
			name:       "Example 9",
			args:       args{[]int{9, 9, 9}},
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := checkMonotopy(tt.args.userSlice); gotResult != tt.wantResult {
				t.Errorf("checkMonotopy() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func BenchmarkCheckMonotopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkMonotopy([]int{5, 6, 7, 6, 11, 12})
	}
}

func BenchmarkCheckSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkSlice([]int{5, 6, 7, 6, 11, 12})
	}
}

func BenchmarkIsMonotonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isMonotonic([]int{5, 6, 7, 6, 11, 12})
	}
}
