package main

import (
	"fmt"
	"strconv"

	"github.com/yyotti/euler_go/src/common"
)

// Surprisingly there are only three numbers that can be written as the sum of
// fourth powers of their digits:
//
//     1634 = 1^4 + 6^4 + 3^4 + 4^4
//     8208 = 8^4 + 2^4 + 0^4 + 8^4
//     9474 = 9^4 + 4^4 + 7^4 + 4^4
//
// As 1 = 1^4 is not a sum it is not included.
//
// The sum of these numbers is 1634 + 8208 + 9474 = 19316.
//
// Find the sum of all the numbers that can be written as the sum of fifth
// powers of their digits.

const power = 5

func main() {
	fmt.Printf("P030A: %d\n", p030A(power))
	fmt.Printf("P030B: %d\n", p030B(power))
}

// 普通にやる
//
// 1を除くので、調べる範囲の最小値は2である。
// 最大値は、桁数をm桁とすると
//   10^(m-1) - 1 <= m*9^k
// を満たす最大のmを求めればよい。
func p030A(k int) int {
	m := 1
	pow9 := pow(9, k)
	for pow(10, m-1)-1 <= m*pow9 {
		m++
	}

	max := pow(10, m-1)
	sum := 0
	for i := 2; i < max; i++ {
		nums, _ := common.SplitNums(strconv.Itoa(i))
		s := 0
		for _, n := range nums {
			s += pow(n, k)
		}
		if s == i {
			sum += i
		}
	}

	return sum
}

// 0の0乗は1と定義する
func pow(n, k int) int {
	if k == 0 {
		return 1
	}

	if k%2 == 0 {
		return pow(n*n, k/2)
	}

	return n * pow(n*n, k/2)
}

// P030Aの高速化をはかってみる
//
// k乗の計算を毎回やるのは無駄なので、最初に作ってキャッシュしておく。
// また、各桁のk乗の和をとるのも数値計算でやる。
func p030B(k int) int {
	m := 1
	pow9 := pow(9, k)
	for pow(10, m-1)-1 <= m*pow9 {
		m++
	}

	pows := make([]int, 10)
	for i := 0; i < len(pows); i++ {
		pows[i] = pow(i, k)
	}

	max := pow(10, m-1)
	sum := 0
	for i := 2; i < max; i++ {
		n := i
		s := 0
		for n > 0 {
			s += pows[n%10]
			n /= 10
		}
		if s == i {
			sum += i
		}
	}

	return sum
}
