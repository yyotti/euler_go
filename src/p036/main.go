package main

import (
	"fmt"
	"strconv"

	"github.com/yyotti/euler_go/src/common"
)

// The decimal number, 585 = 1001001001_2 (binary), is palindromic in both
// bases.
//
// Find the sum of all numbers, less than one million, which are palindromic
// in base 10 and base 2.
//
// (Please note that the palindromic number, in either base, may not include
// leading zeros.)

const max = 1000000

func main() {
	fmt.Printf("P036A: %d\n", p036A(max))
	fmt.Printf("P036B: %d\n", p036B(max))
	fmt.Printf("P036C: %d\n", p036C(max))
}

// 全部調べる
func p036A(max int) int {
	sum := 0
	for i := 1; i < max; i++ {
		if common.IsParindrome(strconv.Itoa(i)) && common.IsParindrome(fmt.Sprintf("%b", i)) {
			sum += i
		}
	}
	return sum
}

// 2進数で回文数になるには末尾の桁が1でなければならないので、奇数だけ調べれば
// いい。
func p036B(max int) int {
	sum := 0
	for i := 1; i < max; i += 2 {
		if common.IsParindrome(strconv.Itoa(i)) && common.IsParindrome(fmt.Sprintf("%b", i)) {
			sum += i
		}
	}
	return sum
}

// 回文数となる2進数を生成していく。
//
// Bを回文数の2進数とする。
// Bが偶数桁の場合、中央に0もしくは1を挿入したものは回文数である。
// Bが奇数桁の場合、中央のビットの隣に同じ値を挿入したものは回文数である。
//
// 2進数の回文数を列挙したうえで、それの10進数表現も回文数になるものの和をとれ
// ばよい。
func p036C(max int) int {
	// 1から始める
	queue := []string{"1"}
	sum := 0
	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:]
		n, _ := strconv.ParseInt(b, 2, 0)
		if int(n) >= max {
			continue
		}
		if common.IsParindromeNum(int(n)) {
			sum += int(n)
		}

		mid := len(b) / 2
		if len(b)%2 == 0 {
			// 中央に0と1を挿入
			for _, s := range []string{"0", "1"} {
				bb := b[:mid] + s + b[mid:]
				queue = append(queue, bb)
			}
		} else {
			// 中央の隣に中央と同じ数を挿入
			var bb string
			if len(b) == 1 {
				bb = b + b
			} else {
				bb = b[:mid] + b[mid:mid+1] + b[mid:]
			}
			queue = append(queue, bb)
		}
	}
	return sum
}
