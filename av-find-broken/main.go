package main

import "fmt"

type CommitInfo struct {
	Hash      string
	BuildTime int
}

func main() {
	commits := []CommitInfo{
		{Hash: "654ec593", BuildTime: 3},
		{Hash: "7ed9a3d6", BuildTime: 5},
		{Hash: "20c1be38", BuildTime: 8},
		{Hash: "6d9eb971", BuildTime: 9},
		{Hash: "4ed905e2", BuildTime: 10},
		{Hash: "654ec59", BuildTime: 16},
		{Hash: "7ed9a3d", BuildTime: 18},
		{Hash: "20c1be3", BuildTime: 20},
		{Hash: "6d9eb97", BuildTime: 25},
		{Hash: "4ed905e", BuildTime: 30},
	}

	thresholdTime := 16

	fmt.Println(FindTheBrokenOne(commits, thresholdTime)) // standard output
}

func FindTheBrokenOne(commits []CommitInfo, thresholdTime int) string {
	// Быстрая проверка, если вдруг последнее значение меньше
	if len(commits) != 0 && commits[len(commits)-1].BuildTime < thresholdTime {
		return ""
	}

	// Если длина меньше 20, то используем обычный перебор (20 вычислено тестами)
	if len(commits) < 20 {
		for i := 0; i <= len(commits)-1; i++ {
			if commits[i].BuildTime >= thresholdTime {
				return commits[i].Hash
			}
			if i == len(commits) {
				return ""
			}
		}
	}

	// Если длина больше 20, то используем бинарный поиск
	// Можно было заморочиться с другими видами поиска, но тут всего 1 млн значений
	low, high := 0, len(commits)-1

	for low < high {
		mid := low + (high-low)/2

		if commits[mid].BuildTime < thresholdTime {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low < len(commits) {
		return commits[low].Hash
	}

	return ""
}
