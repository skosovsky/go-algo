package main

import "testing"

func Test_convertToValidPhone(t *testing.T) {
	type args struct {
		phoneNum string
	}
	tests := []struct {
		name           string
		args           args
		wantValidPhone int
	}{
		{
			name:           "Case 1",
			args:           args{"8(495)430-23-97"},
			wantValidPhone: 74954302397,
		},
		{
			name:           "Case 2",
			args:           args{"+7-4-9-5-43-023-97"},
			wantValidPhone: 74954302397,
		},
		{
			name:           "Case 3",
			args:           args{"4-3-0-2-3-9-7"},
			wantValidPhone: 74954302397,
		},
		{
			name:           "Case 4",
			args:           args{"8-495-430"},
			wantValidPhone: 74958495430,
		},
		{
			name:           "Case 5",
			args:           args{"86406361642"},
			wantValidPhone: 76406361642,
		},
		{
			name:           "Case 6",
			args:           args{"+78047952807"},
			wantValidPhone: 78047952807,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValidPhone := convertToValidPhone(tt.args.phoneNum); gotValidPhone != tt.wantValidPhone {
				t.Errorf("convertToValidPhone() = %v, want %v", gotValidPhone, tt.wantValidPhone)
			}
		})
	}
}
