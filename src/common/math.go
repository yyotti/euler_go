package common

import (
	"math"
	"math/big"
)

// Gcd : 最大公約数
func Gcd(m, n int) int {
	if m < 0 || n < 0 {
		return Gcd(int(math.Abs(float64(m))), int(math.Abs(float64(n))))
	}
	if m < n {
		return Gcd(n, m)
	}
	if n == 0 {
		return m
	}

	return Gcd(n, m%n)
}

// Lcm : 最小公倍数
func Lcm(m, n uint) uint {
	return m * n / uint(Gcd(int(m), int(n)))
}

// Permutation : 標準の順列関数
var Permutation = PermutationA

// PermutationA : 順列
func PermutationA(n, r uint) uint64 {
	perm := uint64(1)
	for i := uint(0); i < r && n > i; i++ {
		perm *= uint64(n - i)
	}
	return perm
}

// PermutationB : 順列
//
// 遅い代わりにバカでかい数値も扱える
func PermutationB(n, r uint) *big.Int {
	res := big.NewInt(1)
	for i := uint(0); i < r && n > i; i++ {
		res.Mul(res, big.NewInt(int64(n-i)))
	}
	return res
}

// Combination : 標準のCombination
var Combination = CombinationD

// CombinationA : 組み合わせ
// nCr = nPr/(r!)
func CombinationA(n, r uint) uint64 {
	if n <= r {
		return 1
	}

	if r > n/2 {
		return CombinationA(n, n-r)
	}

	m := PermutationB(n, r)
	d := PermutationB(r, r)
	m.Div(m, d)
	return m.Uint64()
}

// CombinationB : 組み合わせ
// 分数でやる
func CombinationB(n, r uint) uint64 {
	if n <= r {
		return 1
	}

	if r > n/2 {
		return CombinationB(n, n-r)
	}

	combi := big.NewRat(1, 1)
	for i := uint(0); i < r; i++ {
		combi.Mul(combi, big.NewRat(int64(n-i), int64(r-i)))
	}

	return combi.Num().Uint64()
}

// CombinationC : 組み合わせ
//
// nCrには下記の漸化式が成立する。
//   nCr = (n-1)C(r-1) + (n-1)Cr
func CombinationC(n, r uint) uint64 {
	if r == 0 || n <= r {
		return 1
	}

	if r > n/2 {
		return CombinationC(n, n-r)
	}

	return CombinationC(n-1, r-1) + CombinationC(n-1, r)
}

// CombinationD : 組み合わせ
//
// 基本的にCombinationAと考え方は同じだが、先に約分できるところはする。
// （扱う数を小さくしておくことでオーバーフローを防ぐ）
func CombinationD(n, r uint) uint64 {
	if r == 0 || n <= r {
		return 1
	}

	if r > n/2 {
		return CombinationD(n, n-r)
	}

	ns1 := make([]uint, r)
	ns2 := make([]uint, r)
	for i := uint(0); i < r; i++ {
		ns1[i] = n - i
		ns2[i] = r - i
	}

	for i := 0; i < len(ns2); i++ {
		d := ns2[i]
		if d == 1 {
			continue
		}

		for j := 0; j < len(ns1); j++ {
			k := ns1[j]
			if k == 1 {
				continue
			}

			m := Gcd(int(k), int(d))
			if m != 1 {
				ns1[j] /= uint(m)
				d /= uint(m)
			}

			if d == 1 {
				break
			}
		}
	}

	combi := uint64(1)
	for _, m := range ns1 {
		combi *= uint64(m)
	}

	return combi
}
