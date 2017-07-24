package main

import (
	"fmt"
	"os"

	"github.com/yyotti/euler_go/src/common"
)

const total = 1000

// A Pythagorean triplet is a set of three natural numbers, a < b < c,
// for which,
//
//     a^2 + b^2 = c^2
//
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
func main() {
	printAnswer("P009A", p009A(total))
	printAnswer("P009B", p009B(total))
}

func printAnswer(name string, ans []int) {
	if len(ans) == 1 {
		fmt.Printf("%s: %d\n", name, ans[0])
	} else if len(ans) == 0 {
		fmt.Fprintf(os.Stderr, "%s: No answer\n", name)
	} else {
		fmt.Fprintf(os.Stderr, "%s: Got some answers: %v\n", name, ans)
	}
}

// 素直に全パターンを調べる
//
// a + b + c = N とすると、c = N - a - b。
// a < b < c = N - a - b より、
//   b < N - a - b
//   2b < N - a
//
// また、a = b = c の場合に 3a = N であるから、
//   3a < N
// でなければならない。
func p009A(total int) []int {
	if total < 0 {
		return []int{}
	}

	ans := []int{}
	for a := 1; 3*a < total; a++ {
		for b := a + 1; 2*b < total-a; b++ {
			c := total - a - b

			if a*a+b*b == c*c {
				ans = append(ans, a*b*c)
			}
		}
	}

	return ans
}

// ピタゴラス数の一般項を使って調べる
//
// m,nを互いに素な自然数とし、
//   m > n
//   m - n は奇数
// とする。
//
// このとき、自然数の組(a, b, c)が原始ピタゴラス数であるには
//   (a, b, c) = (m^2 - n^2, 2mn, m^2 + n^2)
//               or
//               (2mn, m^2 - n^2, m^2 + n^2)
// であればよい。
// どちらを選択するかは、計算した結果のa,bの大小関係次第ではあるが、今回は別に
// どちらであっても答えは変わらないので（a,b,cを表示するわけではないので）、
// どちらか一方だけを用いれば良い。
//
// m の範囲は、c < N/2 より
//   m^2 < m^2 + n^2 < N/2
// であるから、
//   m^2 < N/2
// であればよい。
//
// また、原始ピタゴラス数の最小の組は(3,4,5)なので、totalが12未満の場合は無視
// してよい。
func p009B(total int) []int {
	if total < 12 {
		return []int{}
	}

	ans := []int{}
	for m := 2; 2*m*m <= total; m++ {
		n := 1
		if m%2 != 0 {
			n = 2
		}
		for ; n < m; n += 2 {
			if (m-n)%2 != 1 || common.Gcd(m, n) != 1 {
				continue
			}

			a := m*m - n*n
			b := 2 * m * n
			if a >= b {
				a, b = b, a
			}
			c := m*m + n*n

			s := a + b + c
			if total%s == 0 {
				d := total / s
				ans = append(ans, d*a*d*b*d*c)
			}
		}
	}

	return ans
}
