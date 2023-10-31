package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{"A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "Example 2",
			args: args{"race a car"},
			want: false,
		},
		{
			name: "Example 3",
			args: args{" "},
			want: true,
		},
		{
			name: "Example 4",
			args: args{"0P"},
			want: false,
		},
		{
			name: "Example 5",
			args: args{"a"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
