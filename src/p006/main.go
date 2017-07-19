package main

import (
	"fmt"
	"math"
)

const max = 100

// The sum of the squares of the first ten natural numbers is,
//
// 1^2 + 2^2 + ... + 10^2 = 385
//
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
//
// Hence the difference between the sum of the squares of the first ten
// natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred
// natural numbers and the square of the sum.
func main() {
	fmt.Printf("P006A: %d\n", p006A(max))
	fmt.Printf("P006B: %d\n", p006B(max))
}

// 問題で言われている通り、それぞれを計算して差をとる
func p006A(max uint) uint {
	sum1 := uint(0)
	sum2 := uint(0)
	for i := uint(1); i <= max; i++ {
		sum1 += uint(math.Pow(float64(i), 2))
		sum2 += i
	}

	return uint(math.Pow(float64(sum2), 2)) - sum1
}

// 数学で解く
//
// Nを自然数とする。
// 1以上N以下の自然数の二乗和は
//   1^2 + 2^2 + ... + N^2 = N(N+1)(2N+1)/6
// 1以上N以下の自然数の和の二乗は
//   (1 + 2 + ... + N)^2 = {N(N+1)/2}^2
// これらの差をとればいいので、
//   {N(N+1)/2}^2 - N(N+1)(2N+1)/6 = N(N+1)(N-1)(3N+2)/12
func p006B(max uint) uint {
	return max * (max + 1) * (max - 1) * (3*max + 2) / 12
}
