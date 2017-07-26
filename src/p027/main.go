package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

// Euler discovered the remarkable quadratic formula:
//
//     n^2 + n + 41
//
// It turns out that the formula will produce 40 primes for the consecutive
// integer values 0 <= n <= 39. However, when n=40,40^2+40+41=40(40+1)+41 is
// divisible by 41, and certainly when n=41,41^2+41+41 is clearly divisible by
// 41.
//
// The incredible formula n^2−79n+1601 was discovered, which produces 80
// primes for the consecutive values 0 <= n <= 79. The product of the
// coefficients, −79 and 1601, is −126479.
//
// Considering quadratics of the form:
//
//     n^2+an+b, where |a| < 1000 and |b| <= 1000
//     where |n| is the modulus/absolute value of n
//     e.g. |11|=11 and |−4|=4
//
// Find the product of the coefficients, a and b, for the quadratic expression
// that produces the maximum number of primes for consecutive values of n,
// starting with n=0

const max = 1000

func main() {
	fmt.Printf("P027A: %d\n", p027A(max))
}

// 全パターンを調べる
func p027A(max int) int {
	if max < 0 {
		return 0
	}

	maxLen := 0
	prod := 0
	gen := common.NewPrimeGenerator()
	for a := -max + 1; a < max; a++ {
		for b := -max; b <= max; b++ {
			l := 0
			for n := 0; ; n++ {
				m := n*n + a*n + b
				if m < 0 || !common.IsPrime(uint(m), gen) {
					break
				}
				l++
			}
			if maxLen < l {
				maxLen = l
				prod = a * b
			}
		}
	}
	return prod
}

// a,bの範囲をなるべく絞る
//
//   f(n) = n^2 + an + b
// とする。
// f(0) = b で、これが素数になる必要があるので、bは素数でなければならない。
// よって、 2 <= b <= 1000 の中の素数のみ調べればよい。
//
// また、
//   f(n) = n(n+a) + b
// より、n もしくは n+a のどちらかがbの倍数になった場合に素数ではなくなる。
//   n > n+a つまり a < 0 の場合
//     n < b の範囲であればbの倍数にはなることはない
//   n <= n+a つまり a >= 0 の場合
//     n+a < b つまり n < b-a の範囲であればbの倍数になることはない
// よって、 a < 0 の方がより多くの n が得られる可能性がある。
//
// さらに、
//   f(n) = n^2 + an + b
//        = (n + a/2)^2 - (a^2)/4 + b
// で、これは n = -a/2 の時に最小値 b-(a^2)/4 をもつが、f(n) <= 0 の範囲は
// 調べる意味がないので
//   b - (a^2)/4 > 0
// の範囲で調べればよい。つまり
//   b > (a^2)/4
//   4b > a^2
// a < 0 より
//   -2*sqrt(b) < a < 0
// である。
//
// bを大きい方から調べていき、得られたnの最大値をmとする。
// nは最大でも b-1 までなので、b <= m となったらそこで探索を打ち切ってよい。
func p027B(max int) int {
	if max < 0 {
		return 0
	}

	// max以下の素数を列挙
	primes := common.Primes(uint(max))

	// 素数を逆順に調べる
	maxLen := 0
	prod := 0
	// gen := common.NewPrimeGenerator()
	for i := len(primes) - 1; i >= 0; i-- {
		b := int(primes[i])
		if b < maxLen {
			break
		}

		for a := -1; a*a < 4*b; a-- {
			n := 0
			for ; n < b; n++ {
				m := uint(n*n + a*n + b)
				if m < 2 {
					break
				}

				// 素数は列挙されているので、ジェネレータは使わず判定
				isPrime := true
				for _, p := range primes {
					if p*p > m {
						break
					}
					if m%p == 0 {
						isPrime = false
						break
					}
				}
				if !isPrime {
					break
				}
			}
			if maxLen < n {
				maxLen = n
				prod = a * b
			}
		}
	}

	return prod
}
