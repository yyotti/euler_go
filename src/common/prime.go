package common

import (
	"math"
	"sync"
)

// PrimeGenerator : 無限素数ジェネレータ
type PrimeGenerator interface {
	start()
	Next() uint
}

// NewPrimeGenerator : PrimeGeneratorコンストラクタ
func NewPrimeGenerator() PrimeGenerator {
	gen := primeGeneratorB{}
	gen.start()

	return &gen
}

// http://qiita.com/cia_rana/items/2a878181da41033ec1d8 から拝借
// NOTE: リークする可能性がある？
type primeGeneratorA struct {
	ch chan uint
}

func (g *primeGeneratorA) start() {
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
func (g *primeGeneratorA) Next() uint {
	return <-g.ch
}

// PrimeFactors : 素因数分解
func PrimeFactors(n int64) map[int64]int {
	pCounts := map[int64]int{}
	gen := NewPrimeGenerator()
	k := n
	for i := int64(gen.Next()); i < math.MaxUint32 && i*i <= n; i = int64(gen.Next()) {
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

	if k > 1 {
		pCounts[k] = 1
	}

	return pCounts
}

var primes []uint
var primeMax uint
var mutex *sync.Mutex

func init() {
	primes = make([]uint, 0, 1000)
	primes = append(primes, 2, 3)
	primeMax = 3

	mutex = &sync.Mutex{}
}

// キャッシュを使う実装
type primeGeneratorB struct {
	idx int
}

func (g *primeGeneratorB) start() {
	g.idx = 0
}

// Next : 次の素数を取得する
func (g *primeGeneratorB) Next() uint {
	mutex.Lock()
	defer mutex.Unlock()

	var p uint
	if g.idx < len(primes) {
		p = primes[g.idx]
	} else {
		i := primeMax + 2
		for ; ; i += 2 {
			isPrime := true
			for _, p := range primes {
				if p*p > i {
					break
				}
				if i%p == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				primes = append(primes, i)
				primeMax = i
				break
			}
		}
		p = i
	}
	g.idx++
	return p
}
