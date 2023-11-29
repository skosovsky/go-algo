package main

import "testing"

var commits = []CommitInfo{
	{Hash: "654ec593", BuildTime: 1},
	{Hash: "654ec593", BuildTime: 2},
	{Hash: "654ec593", BuildTime: 3},
	{Hash: "654ec593", BuildTime: 4},
	{Hash: "654ec593", BuildTime: 5},
	{Hash: "654ec593", BuildTime: 6},
	{Hash: "654ec593", BuildTime: 7},
	{Hash: "654ec593", BuildTime: 8},
	{Hash: "654ec593", BuildTime: 9},
	{Hash: "654ec593", BuildTime: 10},
}

var thresholdTime = 8

func BenchmarkFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindTheBrokenOne(commits, thresholdTime)
	}
}
