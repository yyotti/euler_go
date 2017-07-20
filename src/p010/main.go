package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

const max = 2000000

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
func main() {
	fmt.Printf("P010A: %d\n", p010A(max))
	fmt.Printf("P010B: %d\n", p010B(max))
}

// 素数ジェネレータで1つ1つ生成して加算する
func p010A(max uint) uint {
	gen := common.NewPrimeGenerator()

	sum := uint(0)
	for p := gen.Next(); p <= max; p = gen.Next() {
		sum += p
	}

	return sum
}

// 最大値が決まっているので、エラトステネスの篩にかける
func p010B(max uint) uint {
	if max < 2 {
		return 0
	}

	sieve := make([]bool, max+1)
	for i := 2; i < len(sieve); i++ {
		sieve[i] = true
	}

	// ふるいにかけつつ足していく
	sum := uint(0)
	for i := uint(2); i < max+1; i++ {
		if !sieve[i] {
			continue
		}

		sum += i
		for j := 2 * i; j < max+1; j += i {
			sieve[j] = false
		}
	}

	return sum
}
