package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

// The fraction 49/98 is a curious fraction, as an inexperienced mathematician
// in attempting to simplify it may incorrectly believe that 49/98 = 4/8,
// which is correct, is obtained by cancelling the 9s.
//
// We shall consider fractions like, 30/50 = 3/5, to be trivial examples.
//
// There are exactly four non-trivial examples of this type of fraction, less
// than one in value, and containing two digits in the numerator and
// denominator.
//
// If the product of these four fractions is given in its lowest common terms,
// find the value of the denominator.
func main() {
	fmt.Printf("P033A: %d\n", p033A())
	fmt.Printf("P033B: %d\n", p033B())
}

// 全部調べる
//
// 1より小さい分数なので、分母が分子より大きい。
// また、分母・分子から0を消去するのは自明な例で、分子が「02」のようになること
// はなく、「19/90」で9を消去してしまうと0除算になり、「20/32」で2を消去すると
// 0になりもとと異なる数になるので、分母・分子ともに10の倍数は無視してよい。
func p033A() int {
	prod := []int{1, 1}
	for d := 11; d < 100; d++ {
		if d%10 == 0 {
			continue
		}

		ds := []int{d / 10, d % 10}
		for n := 11; n < d; n++ {
			if n%10 == 0 {
				continue
			}

			g := common.Gcd(n, d)
			rat1 := []int{n / g, d / g}

			ns := []int{n / 10, n % 10}
			var rat2 []int
			if ds[0] == ns[0] {
				rat2 = []int{ns[1], ds[1]}
			} else if ds[0] == ns[1] {
				rat2 = []int{ns[0], ds[1]}
			}
			if len(rat2) == 0 && ds[0] != ds[1] {
				if ds[1] == ns[0] {
					rat2 = []int{ns[1], ds[0]}
				} else if ds[1] == ns[1] {
					rat2 = []int{ns[0], ds[0]}
				}
			}
			if len(rat2) == 0 {
				continue
			}

			g = common.Gcd(rat2[0], rat2[1])
			rat2[0] /= g
			rat2[1] /= g

			if rat1[0] == rat2[0] && rat1[1] == rat2[1] {
				prod[0] *= rat1[0]
				prod[1] *= rat1[1]
			}
		}
	}

	g := common.Gcd(prod[0], prod[1])
	return prod[1] / g
}

// 逆に考える
//
// 分母・分子から数字を1つずつ消すので、分母・分子ともに1桁の数になる。
// なので、分母・分子ともに1桁の分数（n/1を含む）に対して、分母・分子に0でない
// 同じ数字を加えた数がもとと同じになるか調べる。
func p033B() int {
	prod := []int{1, 1}
	for d := 1; d < 10; d++ {
		for n := 1; n < d; n++ {
			g := common.Gcd(n, d)
			rat1 := []int{n / g, d / g}

			for k := 1; k < 10; k++ {
				ds := []int{d*10 + k, k*10 + d}
				ns := []int{n*10 + k, k*10 + n}
				if ds[0] > ns[0] {
					g := common.Gcd(ds[0], ns[0])
					if rat1[0] == ns[0]/g && rat1[1] == ds[0]/g {
						prod[0] *= ns[0]
						prod[1] *= ds[0]
						continue
					}
				}
				if ds[0] > ns[1] {
					g := common.Gcd(ds[0], ns[1])
					if rat1[0] == ns[1]/g && rat1[1] == ds[0]/g {
						prod[0] *= ns[1]
						prod[1] *= ds[0]
						continue
					}
				}
				if ds[1] != ds[0] {
					if ds[1] > ns[0] {
						g := common.Gcd(ds[1], ns[0])
						if rat1[0] == ns[0]/g && rat1[1] == ds[1]/g {
							prod[0] *= ns[0]
							prod[1] *= ds[1]
							continue
						}
					}
					if ds[1] > ns[1] {
						g := common.Gcd(ds[1], ns[1])
						if rat1[0] == ns[1]/g && rat1[1] == ds[1]/g {
							prod[0] *= ns[1]
							prod[1] *= ds[1]
							continue
						}
					}
				}
			}
		}
	}

	g := common.Gcd(prod[0], prod[1])
	return prod[1] / g
}
