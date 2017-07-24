package main

import "fmt"

// By starting at the top of the triangle below and moving to adjacent numbers
// on the row below, the maximum total from top to bottom is 23.
//
// 3
// 7 4
// 2 4 6
// 8 5 9 3
//
// That is, 3 + 7 + 4 + 9 = 23.
//
// Find the maximum total from top to bottom of the triangle below:
//
// 75
// 95 64
// 17 47 82
// 18 35 87 10
// 20 04 82 47 65
// 19 01 23 75 03 34
// 88 02 77 73 07 63 67
// 99 65 04 28 06 16 70 92
// 41 41 26 56 83 40 80 70 33
// 41 48 72 33 47 32 37 16 94 29
// 53 71 44 65 25 43 91 52 97 51 14
// 70 11 33 28 77 73 17 78 39 68 17 57
// 91 71 52 38 17 14 91 43 58 50 27 29 48
// 63 66 04 68 89 53 67 30 73 16 69 87 40 31
// 04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
//
// NOTE: As there are only 16384 routes, it is possible to solve this problem
// by trying every route. However, Problem 67, is the same challenge with a
// triangle containing one-hundred rows; it cannot be solved by brute force,
// and requires a clever method! ;o)

var nums = [][]int{
	{75},
	{95, 64},
	{17, 47, 82},
	{18, 35, 87, 10},
	{20, 4, 82, 47, 65},
	{19, 1, 23, 75, 3, 34},
	{88, 2, 77, 73, 7, 63, 67},
	{99, 65, 4, 28, 6, 16, 70, 92},
	{41, 41, 26, 56, 83, 40, 80, 70, 33},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
}

func main() {
	fmt.Printf("P018A: %d\n", p018A(nums))
	fmt.Printf("P018B: %d\n", p018B(nums))
	fmt.Printf("P018C: %d\n", p018C(nums))
}

// 全パターンで探索する
//
// 1行目と2行目の和をとると2個の和ができる。そのそれぞれについて、3行目の数字
// と和をとると、2個のそれぞれに対して2個の足す数があるので、4個の和ができる。
// これを繰り返すことになるため、全体の行数をNとすると、最後の1行を除くので
// 2^(N-1)個の和ができることになる。問題に 16384 とあるが、これは 2^(15-1) で
// ある。
// P067では100行の山を相手にすることになり、その全パターンを出そうとすると
// 2^(100-1) 個になるので、別のアプローチが必要というのも頷ける。
//
// NOTE: 上記の理由から、この解法に対して15行より大きい行数の問題を入れること
// は避ける。
func p018A(nums [][]int) int {
	var sums [][]int
	for i, ns := range nums {
		if len(ns) != i+1 {
			return -1
		}
		if i == 0 {
			sums = [][]int{[]int{ns[0], 0}}
			continue
		}

		newSums := make([][]int, 2*len(sums))
		for j, ss := range sums {
			s := ss[0]
			k := ss[1]
			newSums[2*j] = []int{s + ns[k], k}
			newSums[2*j+1] = []int{s + ns[k+1], k + 1}
		}

		sums = newSums
	}

	max := 0
	for _, ss := range sums {
		if max < ss[0] {
			max = ss[0]
		}
	}

	return max
}

// 再帰的にやる
//
// 頂点に着目する。和を最大にするには、そこから左に進んだ場合に得られる和の
// 最大値と右に進んだ場合に得られる和の最大値を比べ、大きい方に進めばよい。
// 左の最大値を出すには、頂点を2行目の左に移し、そこから進める左側と右側の山の
// 最大値をそれぞれ出して、大きい方に進めばよい。
// つまり、再帰的に最大値を出して、大きい方を選んでいくようにする。
func p018B(nums [][]int) int {
	if len(nums) == 0 {
		return 0
	}

	return p018BSub(nums, 0, 0)
}

// 実質的にこっちがP018Bの回答
//
// numsのi行j列から下に辿った場合の最大値を求める
func p018BSub(nums [][]int, i, j int) int {
	if i < 0 || i >= len(nums) {
		return -1
	}
	ns := nums[i]
	if j < 0 || j >= len(ns) {
		return -1
	}

	if i == len(nums)-1 {
		return ns[j]
	}

	left := p018BSub(nums, i+1, j)
	right := p018BSub(nums, i+1, j+1)
	if left < 0 || right < 0 {
		return -1
	}
	if left > right {
		return left + ns[j]
	}
	return right + ns[j]
}

// 高速化のための工夫
//
// 対象の山を下記とする。（問題にあるもの）
//
// 3
// 7 4
// 2 4 6
// 8 5 9 3
//
// P018Bでやった再帰処理で、頂点から最大値を調べていく過程で3行目までいった
// 時点を考える。
// 3行目の2に対しては、次が最終行なので[8]と[5]になり、その大きい方と2を足す。
// 3行目の4に対しては、[5]と[9]になり、その大きい方と4を足す。
// 3行目の6に対しては、[9]と[3]になり、その大きい方と6を足す。
// 最終的には、4行目が消えた下記の山を対象としているのと同じ状況になる。
//
// 3
// 7 4
// 10 13 15
//
// この後も同じ処理を続けていくことになる。
// つまり、頂点から調べるのではなく底辺から調べていき、下記の操作を繰り返すこ
// とで最大値を得ることができる。
//   0. 全体の行数N行とする。
//   1. 底辺(N行目)に着目する。
//   2. 隣り合う数字を比べて大きい方を残す
//   3. 2で残した数列はN-1個の数値からなる。それらとN-1行目の数字を足す。
//   4. 3で足した結果を底辺として1からの処理を繰り返す
func p018C(nums [][]int) int {
	if len(nums) == 0 {
		return 0
	}

	var sums []int
	for i := len(nums) - 1; i >= 0; i-- {
		bottom := nums[i]
		if len(bottom) != i+1 {
			return -1
		}
		if i == len(nums)-1 {
			sums = bottom
			continue
		}

		newSums := make([]int, len(sums)-1)
		for j := 0; j < len(sums)-1; j++ {
			if sums[j] > sums[j+1] {
				newSums[j] = sums[j] + bottom[j]
			} else {
				newSums[j] = sums[j+1] + bottom[j]
			}
		}
		sums = newSums
	}

	return sums[0]
}
