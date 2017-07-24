package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

const max = 20

// 2520 is the smallest number that can be divided by each of the numbers from
// 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?
func main() {
	fmt.Printf("P005A: %d\n", p005A(max))
	fmt.Printf("P005B: %d\n", p005B(max))
}

// 1からmaxまでを素因数分解し、その結果出てきた素因数を全て掛け合わせる
func p005A(max int) int {
	if max < 1 {
		return 0
	}

	pf := map[int64]int{}
	for i := 1; i <= max; i++ {
		for n, c := range common.PrimeFactors(int64(i)) {
			cnt, ok := pf[n]
			if !ok || cnt < c {
				pf[n] = c
			}
		}
	}

	p := 1
	for n, c := range pf {
		for i := 0; i < c; i++ {
			p *= int(n)
		}
	}

	return p
}

// 1からmaxまでの最小公倍数を計算する
func p005B(max int) int {
	if max < 1 {
		return 0
	}

	lcm := 1
	for i := 2; i <= max; i++ {
		lcm = common.Lcm(lcm, i)
	}

	return lcm
}
