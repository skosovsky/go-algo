package main

import (
	"fmt"
)

// Решения взято с репо автора задачи, т.к. мое решение не являлось полным
// https://github.com/Yankovsky/yandex-algos-training/blob/master/hw1/e.py
// при этом я разобрался в решении и переписал его с учетом особенностей языка

func main() {
	var k, m, k2, p2, n2 int
	fmt.Scan(&k, &m, &k2, &p2, &n2) // standard input

	fmt.Println(findPN(k, m, k2, p2, n2)) // standard output
}

func apartOnFloor(m, k, p, n int) (minApart, maxApart int) {
	minApart = k / (m*(p-1) + n)
	maxApart = (k - 1) / (m*(p-1) + n - 1)

	return minApart, maxApart
}

func findPN(k, m, k2, p2, n2 int) (p1, n1 int) {
	var minApart, maxApart, idxApart int
	var optApart []int

	if n2 > m && n2 > 0 {
		p1, n1 = -1, -1

		return p1, n1
	}

	if p2 == 1 && n2 == 1 {
		if k <= k2 {
			p1, n1 = 1, 1

			return p1, n1
		} else {
			minApart = min(k2, k+1)
			maxApart = max(k2, k+1)
		}
	} else {
		minApart, maxApart = apartOnFloor(m, k2, p2, n2)
	}

	for q := minApart; q <= maxApart; q++ {
		if q != 0 && (m*(p2-1)+n2-1)*q+(k2-1)%q == k2-1 {
			optApart = append(optApart, q)
		}
	}

	var p, n int
	p1, n1 = -1, -1

	for _, q := range optApart {
		idxApart = ((k - 1 - (k-1)%q) / q) + 1
		n = idxApart % m
		p = ((idxApart - n) / m) + 1

		if n == 0 {
			n = m
			p--
		}

		if p1 == -1 && n1 == -1 {
			p1, n1 = p, n
		} else {
			if p != p1 {
				p1 = 0
			}
			if n != n1 {
				n1 = 0
			}
		}
	}

	return p1, n1
}
