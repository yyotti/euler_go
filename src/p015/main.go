package main

import (
	"fmt"

	"github.com/yyotti/euler_go/src/common"
)

// Starting in the top left corner of a 2×2 grid, and only being able to move
// to the right and down, there are exactly 6 routes to the bottom right
// corner.
//
//    (図があるが省略)
//
// How many such routes are there through a 20×20 grid?

const (
	x = 20
	y = 20
)

func main() {
	fmt.Printf("P015A(size 10): %d\n", p015A(10, 10))
	fmt.Printf("P015B: %d\n", p015B(x, y))
	fmt.Printf("P015C(size 15): %d\n", p015C(15, 15))
	fmt.Printf("P015D: %d\n", p015D(x, y))
}

// 全パターンを調べ上げる
//
// NOTE: 遅すぎるので問題のサイズでの実行はしない
func p015A(x, y int) int64 {
	if x < 0 || y < 0 {
		return 0
	}
	if x == 0 || y == 0 {
		// 一直線しかない
		return 1
	}

	total := int64(0)

	if x > 0 {
		total += p015A(x-1, y)
	}

	if y > 0 {
		total += p015A(x, y-1)
	}

	return total
}

// 順列・組み合わせでやる(1)
//
// h x w のマス目を左上から右下まで最短で行くには、
//   →→↓→↓→→↓↓→...
// のように表現した場合、「(h+w)個の中にw個の→を置くのは何通りあるか」という
// 問題になる。これは nCr で計算できる。
//
// CombinationBだと多少遅い
func p015B(x, y int) int64 {
	return common.CombinationB(x+y, x)
}

// 順列・組み合わせでやる(2)
//
// CombinationCだとさらに遅いが、ライブラリ不要
// NOTE: 遅すぎるので問題のサイズでの実行はしない
func p015C(x, y int) int64 {
	return common.CombinationC(x+y, x)
}

// 順列・組み合わせでやる(3)
//
// CombinationDを使用(最速)
func p015D(x, y int) int64 {
	return common.CombinationD(x+y, x)
}
