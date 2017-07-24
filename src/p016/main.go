package main

import "fmt"
import "github.com/yyotti/euler_go/src/common"
import "math/big"

// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//
// What is the sum of the digits of the number 2^1000?

const e = 1000

func main() {
	fmt.Printf("P016A: %d\n", p016A(e))
	fmt.Printf("P016B: %d\n", p016B(e))
	fmt.Printf("P016C: %d\n", p016C(e))
}

// 普通にやったらどう考えてもオーバーフローするのでライブラリ使う
func p016A(e int) int {
	b := big.NewInt(2)
	b.Exp(b, big.NewInt(int64(e)), nil)

	ns, _ := common.SplitNums(b.Text(10))
	sum := 0
	for _, d := range ns {
		sum += d
	}

	return sum
}

// ライブラリを使わずにやる(1)
func p016B(e int) int {
	mul := "1"
	for i := 1; i <= e; i++ {
		mul = common.Mul(mul, "2")
	}

	sum := 0
	ds, _ := common.SplitNums(mul)
	for _, d := range ds {
		sum += d
	}

	return sum
}

// ライブラリを使わずにやる(2)
//
// 文字列の掛け算を1より減らしてみる
func p016C(e int) int {
	ms := []int{}
	k := e
	for {
		i := uint(0)
		for ; (1 << i) <= k; i++ {
		}
		if i == 0 {
			break
		}
		ms = append(ms, int(i))
		k -= int(i)
	}

	ns := make([]string, 0, len(ms))
	for _, m := range ms {
		ns = append(ns, pow2(m))
	}

	mul := "1"
	for _, n := range ns {
		mul = common.Mul(mul, n)
	}

	ds, _ := common.SplitNums(mul)
	sum := 0
	for _, d := range ds {
		sum += d
	}

	return sum
}

func pow2(e int) string {
	if e == 0 {
		return "1"
	}

	if e%2 == 1 {
		p := pow2((e - 1) / 2)
		p = common.Mul(p, p)
		return common.Mul("2", p)
	}

	p := pow2(e / 2)
	return common.Mul(p, p)
}
