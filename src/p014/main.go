package main

import "fmt"

// The following iterative sequence is defined for the set of positive
// integers:
//
//     n -> n/2 (n is even)
//     n -> 3n + 1 (n is odd)
//
// Using the rule above and starting with 13, we generate the following
// sequence:
//
//     13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1
//
// It can be seen that this sequence (starting at 13 and finishing at 1)
// contains 10 terms. Although it has not been proved yet (Collatz Problem),
// it is thought that all starting numbers finish at 1.
//
// Which starting number, under one million, produces the longest chain?
//
// NOTE: Once the chain starts the terms are allowed to go above one million.

const max = 1000000

func main() {
	fmt.Printf("P014A: %d\n", p014A(max))
	fmt.Printf("P014B: %d\n", p014B(max))
	fmt.Printf("P014C: %d\n", p014C(max))
}

func collatzChain(start int) []int {
	if start < 1 {
		return []int{}
	}

	chain := []int{}
	for n := start; n != 1; n = collatz(n) {
		chain = append(chain, n)
	}
	chain = append(chain, 1)

	return chain
}

func collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

// 一番素直にやる
func p014A(max int) int {
	if max < 1 {
		return 0
	}

	maxLen := 0
	m := 0
	for i := 1; i <= max; i++ {
		chain := collatzChain(i)
		if maxLen < len(chain) {
			maxLen = len(chain)
			m = i
		}
	}

	return m
}

func collatzCnt(start int) int {
	if start < 1 {
		return 0
	}

	cnt := 1
	for n := start; n != 1; n = collatz(n) {
		cnt++
	}

	return cnt
}

// キャッシュが十分に大きくないと遅いらしい
func collatzCntCached(start int, cache map[int]int) int {
	if start < 1 {
		return 0
	}

	cnt := 1
	for n := start; n != 1; n = collatz(n) {
		if c, ok := cache[n]; ok {
			cnt += c - 1
			break
		}
		cnt++
	}

	return cnt
}

// 数列を作る必要はないので、カウントだけする
func p014B(max int) int {
	if max < 1 {
		return 0
	}

	maxCnt := 0
	m := 0
	for i := 1; i <= max; i++ {
		cnt := collatzCnt(i)
		if maxCnt < cnt {
			maxCnt = cnt
			m = i
		}
	}

	return m
}

// 計算範囲を絞る
//
// n <= N/2 について、nに対するコラッツ数列の長さL(n)を求められれば、L(2n)は
//   L(2n) = L(n) + 1
// であるため、L(n)より長くなる。これは全てのnについて言えるので、計算する範囲
// は n >= N/2 に限定してよい。
func p014C(max int) int {
	if max < 1 {
		return 0
	}

	maxCnt := 0
	m := 0
	for i := max / 2; i <= max; i++ {
		cnt := collatzCnt(i)
		if maxCnt < cnt {
			maxCnt = cnt
			m = i
		}
	}

	return m
}
