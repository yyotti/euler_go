package main

import "fmt"

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

// 階乗進数を使うと解けるぽい
//
// 0からn-1の順列とそれを辞書順に並べた場合の番号は、階乗進数を使って相互変換
// できる。
// 参考：http://wasan.hatenablog.com/entry/2016/05/23/130403
//
// 階乗進数をF(x)と表すことにすると、0から5の順列のF(31021)番目を求めたければ
//   [012345] の左から3番目を取り出す  -> [3] [01245]
//   [01245] の左から1番目を取り出す -> [31] [0245]
//   [0245] の左から0番目を取り出す -> [310] [245]
//   [245] の左から2番目を取り出す -> [3105] [24]
//   [24] の左から1番目を取り出す -> [31054] [2]
//   [2] が余るので末尾へ -> [31054]
// となる。
// よって、n番目の順列を求めたければ、nを階乗進数に変換し、上記の手順で対応す
// る順列を求めてやればよい。
func p024B(nums []int, n int) int {
	if n < 1 {
		return -1
	}

	fs := factoradic(n - 1)
	if len(fs) >= len(nums) {
		return -1
	}
	for _, f := range fs {
		if f >= len(nums) {
			return -1
		}
	}

	if len(fs) < len(nums)-1 {
		fs = append(make([]int, len(nums)-len(fs)-1), fs...)
	}

	ns := make([]int, len(nums))
	copy(ns, nums)
	perm := 0
	for _, f := range fs {
		d := ns[f]
		perm = perm*10 + d
		rest := make([]int, 0, len(ns)-1)
		for i := range ns {
			if i == f {
				continue
			}
			rest = append(rest, ns[i])
		}
		ns = rest
	}

	return perm*10 + ns[0]
}

// 階乗進数へ変換する
//
// 2進数に変換する時と同じように割って余りをとっていくが、割る数を1から順に
// 増やしていく。
func factoradic(n int) []int {
	if n < 0 {
		return nil
	}

	i := 2
	k := n
	ds := []int{}
	for k > 0 {
		ds = append(ds, k%i)
		k /= i
		i++
	}

	for i := 0; i < len(ds)/2; i++ {
		j := len(ds) - i - 1
		ds[i], ds[j] = ds[j], ds[i]
	}

	return ds
}
