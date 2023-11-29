package main

import (
	"testing"
	"time"
)

func TestBestInterestingShop(t *testing.T) {
	type args struct {
		vasyaVisits [][]Visit
		petyaVisits [][]Visit
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{[][]Visit{{{ShopName: "B", VisitedFrom: time.Date(2007, time.July, 1, 11, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 1, 12, 0, 0, 0, time.UTC)}},
				{{ShopName: "A", VisitedFrom: time.Date(2007, time.July, 2, 10, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 2, 11, 0, 0, 0, time.UTC)}},
			},
				[][]Visit{{{ShopName: "B", VisitedFrom: time.Date(2007, time.July, 1, 8, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 1, 11, 1, 0, 0, time.UTC)}},
					{{ShopName: "A", VisitedFrom: time.Date(2007, time.July, 2, 8, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 2, 10, 0, 1, 0, time.UTC)}},
				}},
			want: "B",
		},
		{
			name: "Example 2",
			args: args{[][]Visit{{{ShopName: "B", VisitedFrom: time.Date(2007, time.July, 1, 11, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 1, 12, 0, 0, 0, time.UTC)}},
				{{ShopName: "A", VisitedFrom: time.Date(2007, time.July, 2, 10, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 2, 11, 0, 0, 0, time.UTC)}},
			},
				[][]Visit{{{ShopName: "B", VisitedFrom: time.Date(2007, time.July, 1, 8, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 1, 11, 0, 0, 0, time.UTC)}},
					{{ShopName: "A", VisitedFrom: time.Date(2007, time.July, 2, 8, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 2, 10, 0, 1, 0, time.UTC)}},
				}},
			want: "A",
		},
		{
			name: "Example 3",
			args: args{[][]Visit{{{ShopName: "B", VisitedFrom: time.Date(2007, time.July, 1, 11, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 1, 12, 0, 0, 0, time.UTC)}},
				{{ShopName: "A", VisitedFrom: time.Date(2007, time.July, 2, 10, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 2, 11, 0, 0, 0, time.UTC)}},
			},
				[][]Visit{{{ShopName: "B", VisitedFrom: time.Date(2007, time.July, 1, 8, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 1, 11, 0, 0, 0, time.UTC)}},
					{{ShopName: "A", VisitedFrom: time.Date(2007, time.July, 2, 8, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 2, 10, 0, 0, 0, time.UTC)}},
				}},
			want: "None",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestInterestingShop(tt.args.vasyaVisits, tt.args.petyaVisits); got != tt.want {
				t.Errorf("BestInterestingShop() = %v, want %v", got, tt.want)
			}
		})
	}
}
