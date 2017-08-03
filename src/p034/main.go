package main

import (
	"fmt"
	"strconv"

	"github.com/yyotti/euler_go/src/common"
)

// 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
//
// Find the sum of all numbers which are equal to the sum of the factorial of
// their digits.
//
// Note: as 1! = 1 and 2! = 2 are not sums they are not included.
func main() {
	fmt.Printf("P034A: %d\n", p034A())
	fmt.Printf("P034B: %d\n", p034B())
}

// 1!と2!は除くので、1桁の数は該当しない。
//
// nの桁数をmとして
//   n = a[m-1]*10^(m-1) + ... + a2*10^2 + a1*10^1 + a0
// とすると、
//   a[m-1]! + a[m-2]! + ... + a2! + a1! + a0! < n
// を満たす最大のmを考える。a[k]の最大値は9!であり、m桁の数値の最小値は10^(m-1)
// であるから、
//   m*9! < 10^(m-1)
// を満たせばよい。
//   m = 7のとき 7*9! = 2540160 > 10^6
//   m = 8のとき 8*9! = 2903040 < 10^7
// であるから、7桁の数まで調べる。
func p034A() int {
	sum := 0
	for i := 10; i < 10000000; i++ {
		nums, _ := common.SplitNums(strconv.Itoa(i))

		s := 0
		for _, n := range nums {
			s += int(common.Permutation(n, n))
		}
		if s == i {
			sum += i
		}
	}
	return sum
}

// P034Aの階乗計算をキャッシュ
// 数値の1桁ごとの分割を自力でやる
func p034B() int {
	perms := make([]int, 10)
	p := 1
	for i := 0; i < 10; i++ {
		if i > 0 {
			p *= i
		}
		perms[i] = p
	}

	sum := 0
	for i := 10; i < 10000000; i++ {
		n := i
		s := 0
		for n > 0 {
			// 桁が逆順になるが階乗を足すだけなので問題ない
			s += perms[n%10]
			n /= 10
		}

		if s == i {
			sum += i
		}
	}
	return sum
}
