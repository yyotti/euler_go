package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/yyotti/euler_go/src/common"
)

const digitCount = 3

// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
func main() {
	fmt.Printf("P004A: %d\n", p004A(digitCount))
	fmt.Printf("P004B: %d\n", p004B(digitCount))
	fmt.Printf("P004C: %d\n", p004C(digitCount))
	fmt.Printf("P004D: %d\n", p004D(digitCount))
	fmt.Printf("P004E: %d\n", p004E(digitCount))
}

// 指定された桁数の数同士を掛け合わせて、回文数になるものの最大値を得る
func p004A(digit int) int {
	if digit < 1 {
		return 0
	}

	// 1桁の場合は9で確定
	if digit == 1 {
		return 9
	}

	from := int(math.Pow10(digit - 1))
	to := int(math.Pow10(digit))

	max := 0
	for i := from; i < to; i++ {
		for j := i; j < to; j++ {
			k := i * j
			if isParindromeA(k) && max < k {
				max = k
			}
		}
	}

	return max
}

// 文字列にして判定する
func isParindromeA(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

// 掛け合わせた結果が2桁以上の回文数になるためには、少なくともどちらか一方が
// 11の倍数でなければならない
func p004B(digit int) int {
	if digit < 1 {
		return 0
	}

	// 1桁の場合は9で確定
	if digit == 1 {
		return 9
	}

	from := int(math.Pow10(digit - 1))
	to := int(math.Pow10(digit))

	max := 0
	// iが11の倍数になるように調整
	for i := from + (11 - from%11); i < to; i += 11 {
		for j := from; j < to; j++ {
			k := j * i
			if isParindromeA(k) && max < k {
				max = k
			}
		}
	}

	return max
}

// p004Bで、max判定を先にした方がisParindromeより速くなるか？
func p004C(digit int) int {
	if digit < 1 {
		return 0
	}

	// 1桁の場合は9で確定
	if digit == 1 {
		return 9
	}

	from := int(math.Pow10(digit - 1))
	to := int(math.Pow10(digit))

	max := 0
	// iが11の倍数になるように調整
	for i := from + (11 - from%11); i < to; i += 11 {
		for j := from; j < to; j++ {
			k := j * i
			if max < k && isParindromeA(k) {
				max = k
			}
		}
	}

	return max
}

// 大きい数から順番に計算していって判定する
func p004D(digit int) int {
	if digit < 1 {
		return 0
	}

	// 1桁の場合は9で確定
	if digit == 1 {
		return 9
	}

	from := int(math.Pow10(digit - 1))
	to := int(math.Pow10(digit)) - 1

	max := 0
	maxJ := 0
	// iが11の倍数になるように調整
	for i := to - to%11; i >= from; i -= 11 {
		// iより大きい範囲で調べればいい
		j := to
		for ; j >= i; j-- {
			k := j * i
			if max < k && isParindromeA(k) {
				// maxが変わった時点でそれ以下のjでは見つかるわけがないので切る
				max = k
				break
			}
		}

		if max > 0 {
			if maxJ == 0 || maxJ < j {
				maxJ = j
			} else {
				// 保持しているjよりも小さいjが見つかってもmaxより大きい結果に
				// はならないので、切る
				break
			}
		}
	}

	return max
}

// p004D + isParindromeB (最速)
func p004E(digit int) int {
	if digit < 1 {
		return 0
	}

	// 1桁の場合は9で確定
	if digit == 1 {
		return 9
	}

	from := int(math.Pow10(digit - 1))
	to := int(math.Pow10(digit)) - 1

	max := 0
	maxJ := 0
	// iが11の倍数になるように調整
	for i := to - to%11; i >= from; i -= 11 {
		// iより大きい範囲で調べればいい
		j := to
		for ; j >= i; j-- {
			k := j * i
			if max < k && common.IsParindromeNum(k) {
				// maxが変わった時点でそれ以下のjでは見つかるわけがないので切る
				max = k
				break
			}
		}

		if max > 0 {
			if maxJ == 0 || maxJ < j {
				maxJ = j
			} else {
				// 保持しているjよりも小さいjが見つかってもmaxより大きい結果に
				// はならないので、切る
				break
			}
		}
	}

	return max
}
