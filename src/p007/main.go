package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

const n = 10001

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.
//
// What is the 10,001st prime number?
func main() {
	fmt.Printf("P007A: %d\n", p007A(n))
	fmt.Printf("P007B: %d\n", p007B(n))
}

// 素数ジェネレータで10001個の素数を生成する
func p007A(n int) int64 {
	if n < 1 {
		return 0
	}

	gen := common.NewPrimeGenerator()

	for i := 1; i < n; i++ {
		gen.Next()
	}

	return int64(gen.Next())
}

// 素数ジェネレータとは異なる方法で素数を探す
func p007B(n int) int64 {
	if n < 1 {
		return 0
	}

	primes := make([]int64, 0, n)
	primes = append(primes, 2)

	i := int64(3)
	for len(primes) < n {
		isPrime := true
		for _, p := range primes {
			if p > i*i {
				break
			}
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
		i += 2
	}

	return primes[len(primes)-1]
}
