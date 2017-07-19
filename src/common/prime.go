package common

import (
	"math"
)

// PrimeGenerator : 無限素数ジェネレータ
// http://qiita.com/cia_rana/items/2a878181da41033ec1d8 から拝借
type PrimeGenerator struct {
	ch chan uint
}

// NewPrimeGenerator : PrimeGeneratorコンストラクタ
func NewPrimeGenerator() *PrimeGenerator {
	gen := PrimeGenerator{
		ch: make(chan uint),
	}

	go gen.start()

	return &gen
}

func (g *PrimeGenerator) start() {
	// Key   元の素数の奇数倍
	// Value 元の素数の2倍
	multiples := map[uint]uint{}

	// 2は例外
	g.ch <- 2

	// 3以上の奇数で素数を探す
	for n := uint(3); ; n += 2 {
		factor, ok := multiples[n]

		// factorがあるなら、nはfactor/2の倍数なので素数ではない
		if ok {
			delete(multiples, n)
		} else {
			// 後に追加する倍数を奇数に限定するための調整
			factor = n * 2
		}

		// 新たな倍数を追加
		for newN := n + factor; ; newN += factor {
			if _, ok := multiples[newN]; !ok {
				multiples[newN] = factor
				break
			}
		}

		if !ok {
			g.ch <- n
		}
	}
}

// Next : 次の素数を取得する
func (g *PrimeGenerator) Next() uint {
	return <-g.ch
}

// PrimeFactors : 素因数分解
func PrimeFactors(n uint) map[uint]uint {
	pCounts := map[uint]uint{}
	if n < 2 {
		return pCounts
	}

	gen := NewPrimeGenerator()
	k := n
	for i := gen.Next(); float64(i) <= math.Sqrt(float64(n)); i = gen.Next() {
		for k%i == 0 {
			cnt, ok := pCounts[i]
			if ok {
				pCounts[i] = cnt + 1
			} else {
				pCounts[i] = 1
			}

			k /= i
		}
	}

	if k != 1 {
		pCounts[k] = 1
	}

	return pCounts
}
