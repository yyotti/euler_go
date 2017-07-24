package main

import (
	"fmt"
	"math"

	"github.com/yyotti/euler_go/src/common"
)

const num = 600851475143

// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?
func main() {
	fmt.Printf("P003A: %d\n", p003A(num))
	fmt.Printf("P003B: %d\n", p003B(num))
	fmt.Printf("P003C: %d\n", p003C(num))
	fmt.Printf("P003D: %d\n", p003D(num))
}

// メモ化
var primes = map[int64]struct{}{}

type intp int64

func (n intp) isPrime() bool {
	if n < 2 {
		return false
	}

	if _, ok := primes[int64(n)]; ok {
		return true
	}

	for i := int64(2); i < math.MaxInt32 && i*i <= int64(n); i++ {
		if n%intp(i) == 0 {
			return false
		}
	}

	primes[int64(n)] = struct{}{}
	return true
}

// nを純粋に素因数分解して、素因数の最大値を得る
func p003A(n int64) int64 {
	pf := primeFactors(n)

	max := int64(1)
	for p := range pf {
		if p > max {
			max = p
		}
	}

	return max
}

func primeFactors(n int64) map[int64]int {
	pCounts := map[int64]int{}
	if n < 2 {
		return pCounts
	}

	k := n
	for i := int64(2); i < math.MaxInt32 && i*i <= n; i++ {
		if !intp(i).isPrime() {
			continue
		}

		for k%int64(i) == 0 {
			cnt, ok := pCounts[int64(i)]
			if ok {
				pCounts[int64(i)] = cnt + 1
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
func p003B(n int64) int64 {
	// 実装はp003Aと同じ
	pf := common.PrimeFactors(n)

	max := int64(1)
	for p := range pf {
		if p > max {
			max = p
		}
	}

	return max
}

// 素因数分解はせずに、nを割り切れる最大の素数を探す
func p003C(n int64) int64 {
	if n < 2 {
		return 1
	}

	// とりあえず2で割れるだけ割る
	k := intp(n)
	for ; k%2 == 0; k /= 2 {
	}
	if k == 1 {
		return 2
	}

	i := intp(3)
	for ; float64(i) <= math.Sqrt(float64(k)); i += 2 {
		if !i.isPrime() {
			continue
		}

		for ; k%i == 0; k /= i {
		}
	}

	if k == 1 {
		return int64(i)
	}

	return int64(k)
}

// p003Cのジェネレータ利用版
func p003D(n int64) int64 {
	if n < 2 {
		return 1
	}

	gen := common.NewPrimeGenerator()
	k := n
	max := int64(0)
	for p := int64(gen.Next()); p*p <= k; p = int64(gen.Next()) {
		for ; k%p == 0; k /= p {
		}
		max = p
	}

	if k == 1 {
		return max
	}

	return k
}
