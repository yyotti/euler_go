package main

import (
	"fmt"
	"math"
	"math/big"
)

// The Fibonacci sequence is defined by the recurrence relation:
//
// F(n) = F(n−1) + F(n−2), where F(1) = 1 and F(2) = 1.
// Hence the first 12 terms will be:
//
// F(1) = 1
// F(2) = 1
// F(3) = 2
// F(4) = 3
// F(5) = 5
// F(6) = 8
// F(7) = 13
// F(8) = 21
// F(9) = 34
// F(10) = 55
// F(11) = 89
// F(12) = 144
// The 12th term, F(12), is the first term to contain three digits.
//
// What is the index of the first term in the Fibonacci sequence to contain
// 1000 digits?

const digits = 1000

func main() {
	fmt.Printf("P025A: %d\n", p025A(digits))
	fmt.Printf("P025B: %d\n", p025B(digits))
}

// 指定された桁数を超えるまでフィボナッチ数列を計算し続ける
func p025A(digits int) int {
	if digits < 1 {
		return -1
	}

	f := big.NewInt(0)
	fib := big.NewInt(1)
	n := 1
	for len(fib.Text(10)) < digits {
		s := fib.Text(10)
		fib.Add(fib, f)
		f.SetString(s, 10)
		n++
	}

	return n
}

// フィボナッチ数列の一般項は
//   F(n) = (P^n)/R5
// で近似できる(ただし0.5未満で誤差がある)。
// ここで、Pは黄金比、R5は5の平方根である。
//
// F(n)の桁数をDとすれば、logを10を底とする対数とした場合、
//   10^(D-1) <= F(n) < 10^D
// より
//   D-1 <= log10(F(n)) < D
// が成立する。よって、
//   D-1 <= log10((P^n)/R5) = n*log10(P) - log10(R5)
// より
//   n*log10(P) >= D-1+log10(R5)
//   n >= (D-1+log10(R5))/log10(P)
// であるから、これを満たす最初のnを求めればよい。
//
// NOTE: この計算はD=1の場合は成立しない。そのため、D=1の場合のみ固定値を返す。
func p025B(digits int) int {
	if digits < 1 {
		return -1
	}

	if digits == 1 {
		return 1
	}

	phi := (1 + math.Sqrt(5)) / 2
	rt5 := math.Sqrt(5)
	n := (float64(digits) - 1 + math.Log10(rt5)) / math.Log10(phi)
	return int(math.Ceil(n))
}
