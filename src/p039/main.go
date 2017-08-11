package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

// If p is the perimeter of a right angle triangle with integral length sides,
// {a,b,c}, there are exactly three solutions for p = 120.
//
// {20,48,52}, {24,45,51}, {30,40,50}
//
// For which value of p <= 1000, is the number of solutions maximised?

const p = 1000

func main() {
	fmt.Printf("P039A: %d\n", p039A(p))
}

// P009の延長
//
// 原始ピタゴラス数を順に調べて、三辺の合計値の倍数をカウントしていく
func p039A(p int) int {
	if p < 0 {
		return 0
	}

	counts := make(map[int]int, p)
	for m := 2; ; m++ {
		n := 1
		if m%2 != 0 {
			n = 2
		}

		if m*m+n*n > p {
			break
		}

		for ; n < m; n += 2 {
			if common.Gcd(m, n) != 1 {
				continue
			}
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n

			s := a + b + c
			for d := 1; s*d <= p; d++ {
				_, ok := counts[s*d]
				if ok {
					counts[s*d]++
				} else {
					counts[s*d] = 1
				}
				// s += s
			}
		}
	}

	maxV := 0
	maxK := 0
	for k, v := range counts {
		if maxV < v {
			maxV = v
			maxK = k
		} else if maxV == v {
			if maxK < k {
				maxK = k
			}
		}
	}

	return maxK
}
