package main

import "fmt"

// Starting with the number 1 and moving to the right in a clockwise direction
// a 5 by 5 spiral is formed as follows:
//
//     21 22 23 24 25
//     20  7  8  9 10
//     19  6  1  2 11
//     18  5  4  3 12
//     17 16 15 14 13
//
// It can be verified that the sum of the numbers on the diagonals is 101.
//
// What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral
// formed in the same way?

const size = 1001

func main() {
	fmt.Printf("P028A: %d\n", p028A(size))
	fmt.Printf("P028B: %d\n", p028B(size))
	fmt.Printf("P028C: %d\n", p028C(size))
}

// 問題の通りに1001x1001の数列(？)を作り、対角線を足す
//
// 問題の性質上、サイズは奇数に限定する。
func p028A(size int) int {
	if size < 0 || size%2 == 0 {
		return -1
	}

	// 指定されたサイズの配列を確保
	nums := make([][]int, size)
	for i := 0; i < size; i++ {
		nums[i] = make([]int, size)
	}

	// 数字を配置する
	directions := []int{1, 1, -1, -1}
	c := size / 2
	i := 0
	j := 0
	nums[i+c][j+c] = 1

	n := 2
	maxN := size * size
	for n < maxN {
		for k, d := range directions {
			var m *int
			if k%2 == 0 {
				m = &j
			} else {
				m = &i
			}

			if d > 0 {
				for *m < d && n <= maxN {
					*m++
					nums[i+c][j+c] = n
					n++
				}
				directions[k]++
			} else {
				for *m > d && n <= maxN {
					*m--
					nums[i+c][j+c] = n
					n++
				}
				directions[k]--
			}
		}
	}

	// 対角線を足す
	sum := 0
	for k := 0; k < size; k++ {
		i := k
		j1 := i
		j2 := size - 1 - j1

		sum += nums[i][j1]
		if j1 != j2 {
			sum += nums[i][j2]
		}
	}

	return sum
}

// 再帰で考える
//
// サイズnの四角形の対角線の和をS(n)とし、サイズnの四角形の各頂点の数字の和を
// T(n)とすると、
//   S(n) = 1  (n = 1)
//        = T(n) + S(n-2)  (n >= 3)
// である。
// また、
//   T(n) = 1  (n = 1)
//        = n^2 + (n^2 - n + 1) + (n^2 - 2n + 2) + (n^2 - 3n + 3)  (n >= 3)
//        = 4n^2 - 6n + 6
// である。
func p028B(size int) int {
	if size < 0 || size%2 == 0 {
		return -1
	}

	if size == 1 {
		return 1
	}

	return t(size) + p028B(size-2)
}

func t(n int) int {
	if n == 1 {
		return 1
	}

	return 4*n*n - 6*n + 6
}

// P028Bの式をさらに計算してダイレクトに求める
//
// n = 1の場合は1である。
//
// n >= 3 の場合
// kを自然数とすれば、n = 2k+1 と書ける。その場合、T(n)をkで表すと
//   T(k) = 4n^2 - 6n + 6
//        = 4(2k+1)^2 - 6(2k+1) + 6
//        = 4(4k^2 + 4k + 1) - 6(2k+1) + 6
//        = 16k^2 + 16k + 4 - 12k - 6 + 6
//        = 16k^2 + 4k + 4
//        = 4(4k^2 + k + 1)
// S(n) = 1 + Σ(k=1 -> (n-1)/2)T(k)
//      = 1 + 4{4((n-1)/2)(((n-1)/2)+1)(2((n-1)/2)+1)/6 + ((n-1)/2)(((n-1)/2)+1)/2 + ((n-1)/2)}
//      = 1 + 4{4*((n-1)/2)*((n+1)/2)*(n)/6 + ((n-1)/2)*((n+1)/2)/2 + (n-1)/2}
//      = 1 + 4(n-1)*(n+1)*(n)/6 + (n-1)*(n+1)/2 + 2(n-1)
//      = 6/6 + 4(n-1)*(n+1)*(n)/6 + 3(n-1)*(n+1)/6 + 12(n-1)/6
//      = {6 + 4(n-1)*(n+1)*(n) + 3(n-1)*(n+1) + 12(n-1)}/6
//      = {6 + 4n^3-4n + 3n^2-3 + 12n - 12}/6
//      = (4n^3 + 3n^2 + 8n - 9)/6
// これはS(1)の場合も成立する。
func p028C(size int) int {
	if size < 0 || size%2 == 0 {
		return -1
	}

	return (4*size*size*size + 3*size*size + 8*size - 9) / 6
}
