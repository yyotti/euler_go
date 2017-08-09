package main

import (
	"fmt"
	"math"

	"github.com/yyotti/euler_go/src/common"
)

// Take the number 192 and multiply it by each of 1, 2, and 3:
//
//     192 x 1 = 192
//     192 x 2 = 384
//     192 x 3 = 576
//
// By concatenating each product we get the 1 to 9 pandigital, 192384576.
// We will call 192384576 the concatenated product of 192 and (1,2,3)
//
// The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4,
// and 5, giving the pandigital, 918273645, which is the concatenated product
// of 9 and (1,2,3,4,5).
//
// What is the largest 1 to 9 pandigital 9-digit number that can be formed as
// the concatenated product of an integer with (1,2, ... , n) where n > 1?
func main() {
	fmt.Printf("P038A: %d\n", p038A())
	fmt.Printf("P038B: %d\n", p038B())
}

// 1,2,...,nをかけていく値をkとする。
// 結果が1から9のパンデジタル数になるためには、kに1をかけた値の桁数と2をかけた
// 値の桁数を足したら9にならなければならないが、n>1の制約があるので k < 10000
// でなければならない。
// そこで、kを大きい方から調べていって、最初に条件を満たす数を探す。
func p038A() int {
	var p uint
	for k := 9999; k > 0; k-- {
		d := 9
		p = 0
		for n := 1; ; n++ {
			l := n * k
			dcnt := common.DigitCount(uint64(l))
			p = uint(float64(p)*math.Pow10(dcnt)) + uint(l)
			d -= dcnt
			if d <= 0 {
				break
			}
		}
		if common.IsPandigital(int(p), 9) {
			break
		}
	}

	return int(p)
}

// 1から9の順列を大きい方から順に調べていって最初に見つかったものを解とする。
// (むしろ遅くなった)
func p038B() int {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	var p int
	for i := 1; ; i++ {
		perm := common.NthPermutation(nums, i)
		for d := 4; d > 0; d-- {
			ns := []int{}
			for len(perm) > 0 {
				var ds []int
				if len(perm) >= 2*d {
					ds = perm[:d]
					perm = perm[d:]
				} else {
					ds = perm
					perm = []int{}
				}
				n := 0
				for _, k := range ds {
					n = n*10 + k
				}
				ns = append(ns, n)
			}

			found := true
			for i := 1; i < len(ns); i++ {
				n := ns[i]
				if n%(i+1) != 0 || n/(i+1) != ns[0] {
					found = false
					break
				}
			}
			if len(ns) > 0 && found {
				for _, k := range ns {
					dcnt := common.DigitCount(uint64(k))
					p = int(float64(p)*math.Pow10(dcnt)) + k
				}
				break
			}
		}
		if p > 0 {
			break
		}
	}
	return p
}
