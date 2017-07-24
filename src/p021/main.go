package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

// Let d(n) be defined as the sum of proper divisors of n (numbers less than n
// which divide evenly into n).
// If d(a) = b and d(b) = a, where a != b, then a and b are an amicable pair
// and each of a and b are called amicable numbers.
//
// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44,
// 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4,
// 71 and 142; so d(284) = 220.
//
// Evaluate the sum of all the amicable numbers under 10000.

const max = 10000

func main() {
	fmt.Printf("P021A: %d\n", p021A(max))
	fmt.Printf("P021B: %d\n", p021B(max))
}

// sum of proper divisors of n
func d(n int) int {
	sum := 0
	for _, i := range common.Divisors(n) {
		if i >= n {
			continue
		}
		sum += i
	}
	return sum
}

// 普通にやる
func p021A(max int) int {
	sum := 0
	for i := 1; i <= max; i++ {
		d1 := d(i)
		if d1 == i || d1 >= max {
			continue
		}

		d2 := d(d1)
		if i == d2 {
			sum += d1
		}
	}
	return sum
}

// キャッシュを使う
func p021B(max int) int {
	cache := make(map[int]int, max)

	sum := 0
	for i := 1; i <= max; i++ {
		d1, ok := cache[i]
		if !ok {
			d1 = d(i)
			cache[i] = d1
		}
		if d1 == i || d1 >= max {
			continue
		}

		d2, ok := cache[d1]
		if !ok {
			d2 = d(d1)
			cache[d1] = d2
		}
		if i == d2 {
			sum += d1
		}
	}
	return sum
}
