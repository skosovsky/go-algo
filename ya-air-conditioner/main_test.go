package main

import "testing"

func Test_detectTemp(t *testing.T) {
	type args struct {
		tempRoom int
		tempCond int
		modeCond string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{10, 20, "heat"},
			want: 20,
		},
		{
			name: "Example 2",
			args: args{10, 20, "freeze"},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectTemp(tt.args.tempRoom, tt.args.tempCond, tt.args.modeCond); got != tt.want {
				t.Errorf("detectTemp() = %v, want %v", got, tt.want)
			}
		})
	}
}
