package main

import (
	"fmt"
	"strconv"

	"github.com/yyotti/euler_go/src/common"
)

// The number, 197, is called a circular prime because all rotations of the
// digits: 197, 971, and 719, are themselves prime.
//
// There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37,
// 71, 73, 79, and 97.
//
// How many circular primes are there below one million?

const max = 1000000

func main() {
	fmt.Printf("P035A: %d\n", p035A(max))
	fmt.Printf("P035B: %d\n", p035B(max))
}

// max以下の素数に対して、循環素数であるかをチェックしていく。
//
// 循環素数197に対して、971と719も循環素数であるから、一度チェックしたものは
// キャッシュしていってよい。
// また、桁のどこかに偶数もしくは5が入っている場合、それを循環させた数は素数に
// ならないものが出てくるので、無視できる。
func p035A(max int) int {
	if max < 1 {
		return 0
	}

	gen := common.NewPrimeGenerator()
	cache := map[int]bool{}
	cnt := 0
	for p := int(gen.Next()); p < max; p = int(gen.Next()) {
		_, ok := cache[p]
		if ok {
			continue
		}

		ns, _ := common.SplitNums(strconv.Itoa(p))
		if containsInvalidNum(ns) {
			continue
		}

		rotates := make([]int, 0, len(ns))
		circular := true
		for i := 0; i < len(ns); i++ {
			q := joinNums(ns)
			if circular && !common.IsPrime(uint(q), gen) {
				circular = false
			}
			rotates = append(rotates, q)
			ns = rotate(ns)
		}

		for _, r := range rotates {
			_, ok := cache[r]
			if ok {
				continue
			}

			if circular && r < max {
				cnt++
			}
			cache[r] = true
		}
	}

	return cnt
}

func containsInvalidNum(ns []int) bool {
	if len(ns) < 2 {
		return false
	}
	for _, d := range ns {
		if d%2 == 0 || d == 5 {
			return true
		}
	}
	return false
}

func rotate(ns []int) []int {
	if len(ns) < 2 {
		return append([]int{}, ns...)
	}

	return append(ns[len(ns)-1:], ns[:len(ns)-1]...)
}

func joinNums(ns []int) int {
	n := 0
	for _, d := range ns {
		n = n*10 + d
	}
	return n
}

// 10未満の素数は全て循環素数である。
//
// 10以上の素数は、各桁に偶数と5の倍数は含んではならないので、
//   1,3,7,9
// だけで構成された数でなければならない。
func p035B(max int) int {
	if max < 1 {
		return 0
	}

	gen := common.NewPrimeGenerator()
	if max < 10 {
		cnt := 0
		for gen.Next() < uint(max) {
			cnt++
		}
		return cnt
	}

	digits := 0
	m := max
	for m > 0 {
		digits++
		m /= 10
	}

	cache := map[int]bool{}
	cnt := 4 // 10未満の素数の数
	nums := [][]int{[]int{1}, []int{3}, []int{7}, []int{9}}
	for i := 2; i <= digits; i++ {
		nums = addDigits(nums)
		for _, ns := range nums {
			s := 0
			for _, d := range ns {
				s += d
			}
			if s%3 == 0 {
				continue
			}

			n := joinNums(ns)
			_, ok := cache[n]
			if ok {
				continue
			}

			rotates := make([]int, 0, digits)
			circular := true
			for i := 0; i < digits; i++ {
				m := rotateN(n, i)
				if circular && !common.IsPrime(uint(m), gen) {
					circular = false
				}
				rotates = append(rotates, m)
			}

			for _, r := range rotates {
				_, ok := cache[r]
				if ok {
					continue
				}

				if circular && r < max {
					cnt++
				}
				cache[r] = true
			}
		}
	}

	return cnt
}

func addDigits(nums [][]int) [][]int {
	ds := []int{1, 3, 7, 9}
	ns := make([][]int, 0, len(nums)*len(ds))
	for _, d := range ds {
		for _, n := range nums {
			ns = append(ns, append(n, d))
		}
	}

	return ns
}

func rotateN(n int, c int) int {
	if c <= 0 {
		return n
	}

	d := 0
	for i := 1; i <= n; i *= 10 {
		d++
	}

	k := 1
	for i := 0; i < c; i++ {
		k *= 10
		if k > n {
			k = 1
		}
	}

	m := n % k
	for i := 0; i < d-c%d; i++ {
		m *= 10
	}

	return m + n/k
}
