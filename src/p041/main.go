package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

// We shall say that an n-digit number is pandigital if it makes use of all
// the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital
// and is also prime.
//
// What is the largest n-digit pandigital prime that exists?
func main() {
	fmt.Printf("P041A: %d\n", p041A())
	fmt.Printf("P041B: %d\n", p041B())
}

// 素直にやる
//
// パンデジタル数になるのは9桁までで、それを満たす最大の数は 987654321 である。
// しかし、1から9の順列は各桁の数字の和が3の倍数(45)になるので、どれも素数では
// ない。同様に、1から8の順列も各桁の数字の和が3の倍数(36)になるので、素数には
// ならない。
// よって、1から7の順列の最大である 7654321 まで調べればよい。
func p041A() int {
	gen := common.NewPrimeGenerator()
	maxP := 0
	for p := gen.Next(); p <= 7654321; p = gen.Next() {
		if common.IsPandigital(int(p), common.DigitCount(uint64(p))) {
			maxP = int(p)
		}
	}

	return maxP
}

// チェビシェフの定理から、任意の自然数 n に対して
//   n < p <= 2n
// を満たす素数pが存在する。よって、n=3827160 とすれば、
//   n = 3827160 < p <= 2n = 7654320 < 7654321
// を満たす素数pが存在するので、1から7の順列を逆順に調べていき、最初に見つけた
// 素数が最大のパンデジタル素数である。
func p041B() int {
	ns := []int{7, 6, 5, 4, 3, 2, 1}
	maxP := 0
	gen := common.NewPrimeGenerator()
	for n := 1; ; n++ {
		perm := common.NthPermutation(ns, n)
		if len(perm) == 0 {
			break
		}
		m := 0
		for _, d := range perm {
			m = m*10 + d
		}
		if common.IsPrime(uint(m), gen) {
			maxP = m
			break
		}
	}

	return maxP
}
