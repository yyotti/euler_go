package main

import (
	"fmt"
	"github.com/yyotti/euler_go/src/common"
	"math"
)

const num = 600851475143

// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?
func main() {
	fmt.Printf("P003A: %d\n", p003A(num))
}

// メモ化
var primes = map[uint64]struct{}{}

type uint64p uint64

func (n uint64p) isPrime() bool {
	if n < 2 {
		return false
	}

	if _, ok := primes[uint64(n)]; ok {
		return true
	}

	for i := uint64p(2); float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}

	primes[uint64(n)] = struct{}{}
	return true
}

// nを純粋に素因数分解して、素因数の最大値を得る
func p003A(n uint64) uint64 {
	pf := primeFactors(n)

	max := uint64(1)
	for p := range pf {
		if p > max {
			max = p
		}
	}

	return max
}

func primeFactors(n uint64) map[uint64]int {
	pCounts := map[uint64]int{}
	if n < 2 {
		return pCounts
	}

	k := n
	for i := uint64(2); float64(i) <= math.Sqrt(float64(n)); i++ {
		if !uint64p(i).isPrime() {
			continue
		}

		for k%i == 0 {
			cnt, ok := pCounts[i]
			if ok {
				pCounts[i] = cnt + 1
			} else {
				pCounts[i] = 1
			}

			k /= i
		}
	}

	if k != 1 {
		pCounts[k] = 1
	}

	return pCounts
}

// p003Aの素因数分解を無限素数ジェネレータを使ってやる
func p003B(n uint64) uint64 {
	// 実装はp003Aと同じ
	pf := primeFactors(n)

	max := uint64(1)
	for p := range pf {
		if p > max {
			max = p
		}
	}

	return max
}

func primeFactorsGen(n uint64) map[uint64]int {
	pCounts := map[uint64]int{}
	if n < 2 {
		return pCounts
	}

	gen := common.NewPrimeGenerator()
	k := n
	for i := gen.Next(); float64(i) <= math.Sqrt(float64(n)); i = gen.Next() {
		for k%i == 0 {
			cnt, ok := pCounts[i]
			if ok {
				pCounts[i] = cnt + 1
			} else {
				pCounts[i] = 1
			}

			k /= i
		}
	}

	if k != 1 {
		pCounts[k] = 1
	}

	return pCounts
}

// 素因数分解はせずに、nを割り切れる最大の素数を探す
func p003C(n uint64) uint64 {
	if n < 2 {
		return 1
	}

	// とりあえず2で割れるだけ割る
	k := uint64p(n)
	for ; k%2 == 0; k /= 2 {
	}
	if k == 1 {
		return 2
	}

	i := uint64p(3)
	for ; float64(i) <= math.Sqrt(float64(k)); i += 2 {
		if !i.isPrime() {
			continue
		}

		for ; k%i == 0; k /= i {
		}
	}

	if k == 1 {
		return uint64(i)
	}

	return uint64(k)
}

// p003Cのジェネレータ利用版
func p003D(n uint64) uint64 {
	if n < 2 {
		return 1
	}

	gen := common.NewPrimeGenerator()
	k := n
	max := uint64(0)
	for p := gen.Next(); float64(p) <= math.Sqrt(float64(k)); p = gen.Next() {
		for ; k%p == 0; k /= p {
		}
		max = p
	}

	if k == 1 {
		return max
	}

	return k
}
