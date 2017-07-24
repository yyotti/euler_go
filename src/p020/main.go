package main

import (
	"fmt"
	"math/big"

	"github.com/yyotti/euler_go/src/common"
)

// n! means n x (n − 1) x ... x 3 x 2 x 1
//
// For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800,
// and the sum of the digits in the number 10! is
// 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
//
// Find the sum of the digits in the number 100!

const max = 100

func main() {
	fmt.Printf("P020A: %d\n", p020A(max))
	fmt.Printf("P020B: %d\n", p020B(max))
}

// ライブラリ使って言われたことをそのままやる
func p020A(max int) int {
	if max < 0 {
		return 0
	}

	f := big.NewInt(1)
	for i := max; i > 1; i-- {
		f.Mul(f, big.NewInt(int64(i)))
	}

	ds, _ := common.SplitNums(f.Text(10))
	sum := 0
	for _, d := range ds {
		sum += d
	}
	return sum
}

// ライブラリ使わずにやってみる
//
// 掛け算の筆算を実装する
func p020B(max int) int {
	if max < 0 {
		return 0
	}

	// 各桁の数字を保持する配列。ただし桁の並びは逆順。
	// 100!にどの程度の桁数が必要か分からないが、とりあえず100桁は最初に確保
	// しておく。
	digits := make([]int, 0, 100)
	digits = append(digits, 1)
	for i := 2; i <= max; i++ {
		c := 0
		for j := 0; j < len(digits); j++ {
			d := digits[j]*i + c
			digits[j] = d % 10
			c = d / 10
		}
		for c > 0 {
			digits = append(digits, c%10)
			c /= 10
		}
	}

	sum := 0
	for _, d := range digits {
		sum += d
	}
	return sum
}
