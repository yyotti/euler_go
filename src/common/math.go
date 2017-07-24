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
func Lcm(m, n int) int {
	return int(math.Abs(float64(m))) * int(math.Abs(float64(n))) / Gcd(m, n)
}

// Permutation : 標準の順列関数
var Permutation = PermutationA

// PermutationA : 順列
func PermutationA(n, r int) int64 {
	if r < 0 || n < r {
		return 0
	}

	perm := int64(1)
	for i := 0; i < r && n > i; i++ {
		perm *= int64(n - i)
	}
	return perm
}

// PermutationB : 順列
//
// 遅い代わりにバカでかい数値も扱える
func PermutationB(n, r int) *big.Int {
	if r < 0 || n < r {
		return big.NewInt(0)
	}

	res := big.NewInt(1)
	for i := 0; i < r && n > i; i++ {
		res.Mul(res, big.NewInt(int64(n-i)))
	}
	return res
}

// Combination : 標準のCombination
var Combination = CombinationD

// CombinationA : 組み合わせ
// nCr = nPr/(r!)
func CombinationA(n, r int) int64 {
	if n < 0 || r < 0 {
		return 0
	}
	if n <= r {
		return 1
	}

	if r > n/2 {
		return CombinationA(n, n-r)
	}

	m := PermutationB(n, r)
	d := PermutationB(r, r)
	m.Div(m, d)
	return m.Int64()
}

// CombinationB : 組み合わせ
// 分数でやる
func CombinationB(n, r int) int64 {
	if n < 0 || r < 0 {
		return 0
	}
	if n <= r {
		return 1
	}

	if r > n/2 {
		return CombinationB(n, n-r)
	}

	combi := big.NewRat(1, 1)
	for i := 0; i < r; i++ {
		combi.Mul(combi, big.NewRat(int64(n-i), int64(r-i)))
	}

	return combi.Num().Int64()
}

// CombinationC : 組み合わせ
//
// nCrには下記の漸化式が成立する。
//   nCr = (n-1)C(r-1) + (n-1)Cr
func CombinationC(n, r int) int64 {
	if n < 0 || r < 0 {
		return 0
	}
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
func CombinationD(n, r int) int64 {
	if n < 0 || r < 0 {
		return 0
	}
	if r == 0 || n <= r {
		return 1
	}

	if r > n/2 {
		return CombinationD(n, n-r)
	}

	ns1 := make([]int, r)
	ns2 := make([]int, r)
	for i := 0; i < r; i++ {
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

			m := Gcd(k, d)
			if m != 1 {
				ns1[j] /= m
				d /= m
			}

			if d == 1 {
				break
			}
		}
	}

	combi := int64(1)
	for _, m := range ns1 {
		combi *= int64(m)
	}

	return combi
}

// Add : 数字を表す文字列同士の和をとる
func Add(a, b string) string {
	ns1, err := SplitNums(a)
	if err != nil {
		panic(err)
	}
	ns2, err := SplitNums(b)
	if err != nil {
		panic(err)
	}

	l := len(ns1)
	if l < len(ns2) {
		l = len(ns2)
	}

	if len(ns1) < l {
		ns1 = append(make([]int, l-len(ns1), l), ns1...)
	}

	if len(ns2) < l {
		ns2 = append(make([]int, l-len(ns2), l), ns2...)
	}

	// 下の桁から和をとっていき、逆順に詰める
	digits := make([]int, 0, l)
	c := 0 // 繰り上がり
	for i := l - 1; i >= 0; i-- {
		d := c
		d += ns1[i]
		d += ns2[i]
		d, c = d%10, d/10
		digits = append(digits, d)
	}
	if c > 0 {
		digits = append(digits, c)
	}

	// 結果を逆順にして文字列にする
	str := make([]byte, len(digits))
	for i, d := range digits {
		str[len(digits)-1-i] = '0' + byte(d)
	}

	return string(str)
}

// Mul : 数字を表す文字列同士の積をとる
func Mul(a, b string) string {
	ns1, err := SplitNums(a)
	if err != nil {
		panic(err)
	}
	ns2, err := SplitNums(b)
	if err != nil {
		panic(err)
	}

	// 桁数が小さい方を「掛ける数」にする
	if len(ns1) < len(ns2) {
		ns1, ns2 = ns2, ns1
	}

	// 下の桁から筆算をしていってそれぞれを文字列で加算する
	mul := "0"
	for i := len(ns2) - 1; i >= 0; i-- {
		d := ns2[i]
		if d == 0 {
			// 0をかけるなら何もしない
			ns1 = append(ns1, 0)
			continue
		}
		c := 0
		digits := make([]int, 0, len(ns1))
		for j := len(ns1) - 1; j >= 0; j-- {
			m := d*ns1[j] + c
			m, c = m%10, m/10
			digits = append(digits, m)
		}
		if c > 0 {
			digits = append(digits, c)
		}

		// 結果を逆順にして文字列にする
		str := make([]byte, len(digits))
		for i, d := range digits {
			str[len(digits)-1-i] = '0' + byte(d)
		}

		// 加算
		mul = Add(mul, string(str))
		ns1 = append(ns1, 0)
	}

	return mul
}

// Divisors : 約数
func Divisors(n int) []int {
	if n < 1 {
		return []int{1}
	}

	ds := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if n/i == i {
				ds = append(ds, i)
			} else {
				ds = append(ds, []int{i, n / i}...)
			}
		}
	}

	return ds
}
