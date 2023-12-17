package main

import "testing"

func Test_well(t *testing.T) {
	type args struct {
		x []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{[]string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"}},
			want: "I smell a series!",
		},
		{
			name: "Example 2",
			args: args{[]string{"bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 3",
			args: args{[]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 4",
			args: args{[]string{"bad", "bad", "bad", "good", "bad", "bad", "good", "bad", "bad", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 5",
			args: args{[]string{"bad", "bad", "bad", "bad", "good", "bad", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 6",
			args: args{[]string{"bad", "bad"}},
			want: "Fail!",
		},
		{
			name: "Example 7",
			args: args{[]string{"bad", "bad", "bad", "bad", "bad"}},
			want: "Fail!",
		},
		{
			name: "Example 8",
			args: args{[]string{"bad", "bad", "bad", "bad", "good", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 9",
			args: args{[]string{"bad"}},
			want: "Fail!",
		},
		{
			name: "Example 10",
			args: args{[]string{"bad", "bad", "bad", "good", "bad", "good", "good", "bad", "bad", "bad", "bad", "good", "good"}},
			want: "I smell a series!",
		},
		{
			name: "Example 11",
			args: args{[]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "bad", "bad", "good"}},
			want: "Publish!",
		},
		{
			name: "Example 12",
			args: args{[]string{"good", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "good", "bad"}},
			want: "I smell a series!",
		},
		{
			name: "Example 13",
			args: args{[]string{"bad", "bad", "bad", "bad", "bad", "good", "bad", "good", "good", "good", "bad", "bad", "good"}},
			want: "I smell a series!",
		},
		{
			name: "Example 14",
			args: args{[]string{"bad", "good", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 15",
			args: args{[]string{"bad", "bad", "bad", "good", "bad", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 16",
			args: args{[]string{"good", "bad", "bad", "bad", "bad", "good", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 17",
			args: args{[]string{"good"}},
			want: "Publish!",
		},
		{
			name: "Example 18",
			args: args{[]string{"good", "good"}},
			want: "Publish!",
		},
		{
			name: "Example 19",
			args: args{[]string{"bad", "bad", "bad", "good", "bad", "bad", "bad", "good", "bad", "bad", "bad", "bad", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 20",
			args: args{[]string{"good", "bad", "good", "good", "bad", "bad", "bad"}},
			want: "I smell a series!",
		},
		{
			name: "Example 21",
			args: args{[]string{"bad", "bad", "bad", "bad", "bad"}},
			want: "Fail!",
		},
		{
			name: "Example 22",
			args: args{[]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "bad", "good"}},
			want: "Publish!",
		},
		{
			name: "Example 23",
			args: args{[]string{"bad", "good", "bad", "bad", "good", "bad", "good", "bad", "bad", "bad", "bad", "bad", "bad", "bad"}},
			want: "I smell a series!",
		},
		{
			name: "Example 24",
			args: args{[]string{"bad", "good"}},
			want: "Publish!",
		},
		{
			name: "Example 25",
			args: args{[]string{"bad", "good", "bad", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 26",
			args: args{[]string{"good", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good"}},
			want: "Publish!",
		},
		{
			name: "Example 27",
			args: args{[]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"}},
			want: "Fail!",
		},
		{
			name: "Example 28",
			args: args{[]string{"bad", "bad", "good", "bad", "bad", "good", "bad", "bad", "bad", "bad", "bad", "good", "good", "bad", "good", "bad"}},
			want: "I smell a series!",
		},
		{
			name: "Example 29",
			args: args{[]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"}},
			want: "Fail!",
		},
		{
			name: "Example 30",
			args: args{[]string{"bad", "bad", "bad", "bad", "good", "bad", "bad", "bad"}},
			want: "Publish!",
		},
		{
			name: "Example 31",
			args: args{[]string{"bad", "bad", "good", "bad", "bad", "good", "bad", "good", "bad", "bad", "bad"}},
			want: "I smell a series!",
		},
		{
			name: "Example 32",
			args: args{[]string{"bad", "bad", "bad", "bad", "bad"}},
			want: "Fail!",
		},
		{
			name: "Example 33",
			args: args{[]string{"good", "bad", "bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad", "bad", "good", "bad", "bad"}},
			want: "I smell a series!",
		},
		{
			name: "Example 34",
			args: args{[]string{"bad", "bad", "bad", "bad", "bad", "good", "bad", "bad"}},
			want: "Publish!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := well(tt.args.x); got != tt.want {
				t.Errorf("well() = %v, want %v", got, tt.want)
			}
		})
	}
}
