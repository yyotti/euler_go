package main

import (
	"fmt"
	"sync"

	"github.com/yyotti/euler_go/src/common"
)

// We shall say that an n-digit number is pandigital if it makes use of all
// the digits 1 to n exactly once; for example, the 5-digit number, 15234, is
// 1 through 5 pandigital.
//
// The product 7254 is unusual, as the identity, 39 x 186 = 7254, containing
// multiplicand, multiplier, and product is 1 through 9 pandigital.
//
// Find the sum of all products whose multiplicand/multiplier/product identity
// can be written as a 1 through 9 pandigital.
//
// HINT: Some products can be obtained in more than one way so be sure to only
// include it once in your sum.

var digits = 9

func main() {
	fmt.Printf("P032A: %d\n", p032A(digits))
	fmt.Printf("P032B: %d\n", p032B(digits))
}

// 与えられた数の順列を作り、それぞれで掛け算が成立するかをチェックする
//
// a x b = c をチェックするわけだが、a < b の制限のもとでやる。
func p032A(digits int) int {
	if digits < 3 {
		return 0
	}

	nums := make([]int, digits)
	for i := 0; i < digits; i++ {
		nums[i] = i + 1
	}

	max := common.Permutation(digits, digits)
	sum := 0
	cache := map[int]struct{}{}
	for i := 1; i <= int(max); i++ {
		perm := common.NthPermutation(nums, i)

		for j := 0; j < len(perm)/3; j++ {
			a := 0
			for _, d := range perm[:j+1] {
				a = a*10 + d
			}
			for k := j + 1; k <= len(perm[j+1:])/2; k++ {
				b := 0
				for _, d := range perm[j+1 : j+k+1] {
					b = b*10 + d
				}

				c := 0
				for _, d := range perm[j+k+1:] {
					c = c*10 + d
				}

				if a*b == c {
					if _, ok := cache[c]; !ok {
						sum += c
						cache[c] = struct{}{}
					}
				}
			}
		}
	}

	return sum
}

// P032Aを非同期で
func p032B(digits int) int {
	if digits < 3 {
		return 0
	}

	nums := make([]int, digits)
	for i := 0; i < digits; i++ {
		nums[i] = i + 1
	}

	max := common.Permutation(digits, digits)
	sum := 0
	cache := map[int]struct{}{}
	wg := &sync.WaitGroup{}
	wg.Add(int(max))
	ch := make(chan int)
	for i := 1; i <= int(max); i++ {
		go func(i int, ch chan<- int) {
			defer wg.Done()
			perm := common.NthPermutation(nums, i)

			for j := 0; j < len(perm)/3; j++ {
				a := 0
				for _, d := range perm[:j+1] {
					a = a*10 + d
				}
				for k := j + 1; k <= len(perm[j+1:])/2; k++ {
					b := 0
					for _, d := range perm[j+1 : j+k+1] {
						b = b*10 + d
					}

					c := 0
					for _, d := range perm[j+k+1:] {
						c = c*10 + d
					}

					if a*b == c {
						ch <- c
					}
				}
			}
		}(i, ch)
	}
	go func(ch <-chan int) {
		for s := range ch {
			if _, ok := cache[s]; !ok {
				sum += s
				cache[s] = struct{}{}
			}
		}
	}(ch)
	wg.Wait()
	close(ch)

	return sum
}
