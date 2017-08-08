package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

// The number 3797 has an interesting property. Being prime itself, it is
// possible to continuously remove digits from left to right, and remain prime
// at each stage: 3797, 797, 97, and 7. Similarly we can work from right to
// left: 3797, 379, 37, and 3.
//
// Find the sum of the only eleven primes that are both truncatable from left
// to right and right to left.
//
// NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
func main() {
	fmt.Printf("P037A: %d\n", p037A())
	fmt.Printf("P037B: %d\n", p037B())
}

// 性質を満たす素数が11個あるようなので、順に調べていって11個になったらやめる。
func p037A() int {
	gen := common.NewPrimeGenerator()
	g := common.NewPrimeGenerator()
	cnt := 0
	sum := 0
	for cnt < 11 {
		p := gen.Next()
		if isTruncatablePrime(int(p), g) {
			sum += int(p)
			cnt++
		}
	}

	return sum
}

func rightTruncates(n int) []int {
	if n < 10 {
		return []int{}
	}

	m := n
	nums := []int{}
	for m >= 10 {
		nums = append(nums, m/10)
		m /= 10
	}

	return nums
}

func leftTruncates(n int) []int {
	if n < 10 {
		return []int{}
	}

	k := 10
	for k <= n {
		k *= 10
	}

	m := n
	nums := []int{}
	for k > 1 {
		if m%k != m {
			nums = append(nums, m%k)
			m %= k
		}
		k /= 10
	}

	return nums
}

func isLeftTruncatablePrime(p int, gen common.PrimeGenerator) bool {
	if p < 10 {
		return false
	}
	for _, n := range leftTruncates(p) {
		if !common.IsPrime(uint(n), gen) {
			return false
		}
	}
	return true
}

func isRightTruncatablePrime(p int, gen common.PrimeGenerator) bool {
	if p < 10 {
		return false
	}
	for _, n := range rightTruncates(p) {
		if !common.IsPrime(uint(n), gen) {
			return false
		}
	}
	return true
}

func isTruncatablePrime(p int, gen common.PrimeGenerator) bool {
	return isLeftTruncatablePrime(p, gen) && isRightTruncatablePrime(p, gen)
}

// 生成していく方向
func p037B() int {
	left := []int{}
	right := []int{}
	for i := 1; i < 10; i++ {
		left = append(left, i)
		if i%2 != 0 && i%5 != 0 {
			// 右に偶数や5を付与しても絶対に素数にはならないので無視する
			right = append(right, i)
		}
	}

	gen := common.NewPrimeGenerator()

	// ベースにする素数
	ps := set{}
	for _, n := range []int{2, 3, 5, 7} {
		ps.add(n)
	}
	sum := 0
	for d := 2; ; d++ {
		rp := set{}
		for _, r := range right {
			for n := range ps {
				k := n*10 + r
				if common.IsPrime(uint(k), gen) && isRightTruncatablePrime(k, gen) {
					rp.add(k)
				}
			}
		}
		if len(rp) == 0 {
			break
		}

		lp := set{}
		for _, l := range left {
			for i := 0; i < d-1; i++ {
				l *= 10
			}

			for n := range ps {
				k := l + n
				if common.IsPrime(uint(k), gen) && isLeftTruncatablePrime(k, gen) {
					lp.add(k)
				}
			}
		}
		if len(lp) == 0 {
			break
		}

		tps := lp.intersect(rp)
		for p := range tps {
			sum += p
		}

		ps = lp.union(rp)
	}

	return sum
}

type set map[int]struct{}

func (s set) add(n int) {
	s[n] = struct{}{}
}

func (s set) union(t set) set {
	u := set{}
	for k := range s {
		u.add(k)
	}
	for k := range t {
		u.add(k)
	}
	return u
}

func (s set) intersect(t set) set {
	u := set{}
	for k := range s {
		if _, ok := t[k]; ok {
			u.add(k)
		}
	}
	return u
}
