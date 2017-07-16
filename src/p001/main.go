package main

import (
	"fmt"
)

// P001
// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
func main() {
	fmt.Printf("P001A: %d\n", p001A(1000, []int{3, 5}))
	fmt.Printf("P001B: %d\n", p001B(1000, []int{3, 5}))
	fmt.Printf("P001Z: %d\n", p001Z(1000))
}

// 素直にやる
func p001A(max int, ds []int) int {
	sum := 0
	for i := 1; i < max; i++ {
		for _, d := range ds {
			if i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}

// nずつ増やして加算していく
// dsの要素数が多い場合はこちらの方が高速だが、問題にあるように3と5だけでやる
// ならAの方が速い
func p001B(max int, ds []int) int {
	sum := 0
	for i, n := range ds {
		for j := 1; j*n < max; j++ {
			k := j * n
			found := false
			for _, n := range ds[i+1:] {
				if k%n == 0 {
					found = true
					break
				}
			}
			if !found {
				sum += k
			}
		}
	}

	return sum
}

// 問題のケースでしか使えない代わりに最高速。
func p001Z(max int) int {
	n3 := (max - 1) / 3
	n5 := (max - 1) / 5
	n15 := (max - 1) / 15

	s3 := (1 + n3) * 3 * n3 / 2
	s5 := (1 + n5) * 5 * n5 / 2
	s15 := (1 + n15) * 15 * n15 / 2

	return s3 + s5 - s15
}
