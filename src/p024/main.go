package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

// A permutation is an ordered arrangement of objects. For example, 3124 is
// one possible permutation of the digits 1, 2, 3 and 4. If all of the
// permutations are listed numerically or alphabetically, we call it
// lexicographic order. The lexicographic permutations of 0, 1 and 2 are:
//
// 012   021   102   120   201   210
//
// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3,
// 4, 5, 6, 7, 8 and 9?

var nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

const n = 1000000

func main() {
	// Aは遅いので問題と同じは避ける
	fmt.Printf("P024A([0-7],30000): %d\n", p024A([]int{0, 1, 2, 3, 4, 5, 6, 7}, 30000))
	fmt.Printf("P024B: %d\n", p024B(nums, n))
}

// 10!個の数を作ってその指定番数を求める
// NOTE: やはり遅い
func p024A(nums []int, n int) int {
	if n < 1 {
		return -1
	}

	perms := permutations(nums)
	if len(perms) < n {
		return -1
	}

	perm := perms[n-1]
	num := 0
	for _, d := range perm {
		num = num*10 + d
	}
	return num
}

func permutations(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	perms := [][]int{}
	for _, n := range nums {
		rest := make([]int, 0, len(nums)-1)
		for _, n2 := range nums {
			if n == n2 {
				continue
			}
			rest = append(rest, n2)
		}

		for _, perm := range permutations(rest) {
			perms = append(perms, append([]int{n}, perm...))
		}
	}

	return perms
}

// 階乗進数でやる
func p024B(nums []int, n int) int {
	perm := common.NthPermutation(nums, n)
	if len(perm) == 0 {
		return -1
	}

	p := 0
	for _, n := range perm {
		p = p*10 + n
	}

	return p
}
