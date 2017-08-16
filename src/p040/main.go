package main

import "fmt"
import "github.com/yyotti/euler_go/src/common"
import "math"
import "strconv"

// An irrational decimal fraction is created by concatenating the positive
// integers:
//
//     0.123456789101112131415161718192021...
//
// It can be seen that the 12th digit of the fractional part is 1.
//
// If dn represents the nth digit of the fractional part, find the value of
// the following expression.
//
// d1 x d10 x d100 x d1000 x d10000 x d100000 x d1000000

const k = 7

func main() {
	fmt.Printf("P040A: %d\n", p040A(k))
	fmt.Printf("P040B: %d\n", p040B(k))
}

// そのままやる
//
// dnを毎回計算するのは無駄だしキャッシュを使うのも面倒なので、関数にはせずに
// やる。
func p040A(k int) int {
	if k < 0 {
		return -1
	}

	maxN := int(math.Pow10(k - 1))

	buf := make([]int, 0, maxN)
	buf = append(buf, 0) // 桁合わせのダミー
	i := 1
	for len(buf) <= maxN {
		ns, _ := common.SplitNums(strconv.Itoa(i))
		buf = append(buf, ns...)
		i++
	}

	p := 1
	for n := 1; n <= maxN; n *= 10 {
		p *= buf[n]
	}

	return p
}

// 一般項から求める
//
// チャンパーノウン定数を数列とみなす。
// 自然数nの桁数をkとすれば、k桁で最小の自然数は 10^(k-1) である。
//   1桁の数は9個ある。
//   2桁の数は、99から9(1桁の数の個数)を引いた90個ある。
//   3桁の数は、999から2桁以下の数の個数を引いた900個ある。
//   ...
//   (k-1)桁の数は、10^(k-1) - 1から(k-2)桁以下の数の個数を引いた
//     10^(k-1) - 1 - (10^(k-2) - 1) = 10^(k-1) - 10^(k-2)
//                                   = 10^(k-2) * (10 - 1)
//                                   = 9 * 10^(k-2)
//   個ある。
// それぞれの個数に対して、桁数を掛け合わせた数だけの項数をとるので、まず小数
// 第m位の数が何桁の自然数にあたるかを求められる。
// 次に、nは 10^(k-1) から数えて n - 10^(k-1) 番目の数なので、そこからmが入っ
// ている n を求めることができる。
func p040B(k int) int {
	if k < 0 {
		return -1
	}

	m := 1
	p := 1
	for i := 1; i <= k; i++ {
		p *= d(m)
		m *= 10
	}

	return p
}

func d(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	}

	k := 1
	d := 1 // k桁の自然数の最小値
	s := 9
	p := 9
	for s <= n {
		p = p / k * (k + 1) * 10
		s += p
		d *= 10
		k++
	}
	s -= p

	e := 0
	m := 0
	for ; e < n-s; e += k {
		m++
	}
	s += e - k
	d += m - 1

	for i := 0; i < k-(n-s); i++ {
		d /= 10
	}

	return d % 10
}
