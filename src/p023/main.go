package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

// A perfect number is a number for which the sum of its proper divisors is
// exactly equal to the number. For example, the sum of the proper divisors of
// 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect
// number.
//
// A number n is called deficient if the sum of its proper divisors is less
// than n and it is called abundant if this sum exceeds n.
//
// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
// number that can be written as the sum of two abundant numbers is 24. By
// mathematical analysis, it can be shown that all integers greater than 28123
// can be written as the sum of two abundant numbers. However, this upper
// limit cannot be reduced any further by analysis even though it is known
// that the greatest number that cannot be expressed as the sum of two
// abundant numbers is less than this limit.
//
// Find the sum of all the positive integers which cannot be written as the
// sum of two abundant numbers.
func main() {
	fmt.Printf("P023A: %d\n", p023A())
	fmt.Printf("P023B: %d\n", p023B())
}

func isAbundant(n int) bool {
	if n < 1 {
		return false
	}

	divisors := common.Divisors(n)
	sum := 0
	for _, d := range divisors {
		if d == n {
			continue
		}
		sum += d
	}

	return sum > n
}

const max = 28123

// まともにやってみる
func p023A() int {
	// 過剰数かどうかはキャッシュしておく
	cache := make(map[int]bool, max)
	abundants := make([]int, 0, max)

	sum := 0
	for i := 1; i <= max; i++ {
		// まず対象が過剰数かどうかを判定してキャッシュに詰める
		cache[i] = isAbundant(i)
		if isA, _ := cache[i]; isA {
			abundants = append(abundants, i)
		}

		sumOfAbundants := false
		for _, a := range abundants {
			if 2*a > i {
				break
			}
			b := i - a
			isA, ok := cache[b]
			if !ok {
				isA = isAbundant(b)
				cache[b] = isA
			}
			if isA {
				sumOfAbundants = true
				break
			}
		}
		if !sumOfAbundants {
			sum += i
		}
	}

	return sum
}

// 過剰数を小さい順に列挙していく。
// まず最小の過剰数として12が挙がり、その和で表すことができるのは24のみである。
// よって、過剰数の和で表すことができる数は [24] となる。
// 次の過剰数は18なので、12と18の和で表すことができる数30と、18+18=36が追加さ
// れる。ここまでで [24,30,36] である。
// 次の過剰数は20で、[12,18]と20の和で表すことができるのは
//   12+20=32, 18+20=38
// である。この2つと20+20=40を加え、[24,30,36,32,38,40]となる。
// これを繰り返していき、1～maxの総和から過剰数の和で表せる数を全て引く。
func p023B() int {
	// 和で表すことができる数のキャッシュ
	cache := make(map[int]struct{}, max)

	sum := (1 + max) * max / 2
	abundants := make([]int, 0, max)
	for i := 1; i <= max; i++ {
		if !isAbundant(i) {
			continue
		}

		abundants = append(abundants, i)
		for _, a := range abundants {
			s := a + i
			if _, ok := cache[s]; ok {
				continue
			}

			if s <= max {
				sum -= s
			}
			cache[s] = struct{}{}
		}
	}

	return sum
}
