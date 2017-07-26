package common

import (
	"math"
	"sync"
)

// PrimeGenerator : 無限素数ジェネレータ
type PrimeGenerator interface {
	start()
	Next() uint
	Reset()
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
func PrimeFactors(n int64, gen ...PrimeGenerator) map[int64]int {
	pCounts := map[int64]int{}
	var g PrimeGenerator
	if len(gen) == 0 {
		g = NewPrimeGenerator()
	} else {
		g = gen[0]
		g.Reset()
	}
	k := n
	for i := int64(g.Next()); i < math.MaxUint32 && i*i <= n; i = int64(g.Next()) {
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

// キャッシュを使う実装
type primeGeneratorB struct {
	primes   []uint
	primeMax uint
	mutex    *sync.Mutex
	idx      int
}

func (g *primeGeneratorB) start() {
	g.idx = 0

	g.primes = make([]uint, 0, 1000)
	g.primes = append(g.primes, 2, 3)
	g.primeMax = 3

	g.mutex = &sync.Mutex{}
}

// Next : 次の素数を取得する
func (g *primeGeneratorB) Next() uint {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	var p uint
	if g.idx < len(g.primes) {
		p = g.primes[g.idx]
	} else {
		i := g.primeMax + 2
		for ; ; i += 2 {
			isPrime := true
			for _, p := range g.primes {
				if p*p > i {
					break
				}
				if i%p == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				g.primes = append(g.primes, i)
				g.primeMax = i
				break
			}
		}
		p = i
	}
	g.idx++
	return p
}

// Reset : 生成する素数を最初に戻す
func (g *primeGeneratorB) Reset() {
	g.idx = 0
}

// IsPrime : 素数判定
func IsPrime(n uint, gen ...PrimeGenerator) bool {
	if n < 2 {
		return false
	}

	var g PrimeGenerator
	if len(gen) == 0 {
		g = NewPrimeGenerator()
	} else {
		g = gen[0]
		g.Reset()
	}

	for p := g.Next(); p*p <= n; p = g.Next() {
		if n%p == 0 {
			return false
		}
	}

	return true
}

// Primes : max以下の素数を得る
var Primes = sieveA

func sieveA(max uint) []uint {
	if max < 2 {
		return []uint{}
	}

	// 初期値でtrueを設定するのが無駄なので、trueなら素数「ではない」とする
	sieve := make([]bool, max+1)

	primes := make([]uint, 0, max)
	for i := uint(2); i <= max; i++ {
		if sieve[i] {
			continue
		}

		primes = append(primes, i)

		for j := 2 * i; j < max+1; j += i {
			sieve[j] = true
		}
	}

	return primes
}

func sieveB(max uint) []uint {
	if max < 2 {
		return []uint{}
	}

	primes := make([]uint, 0, max)
	primes = append(primes, 2)
	ms := map[uint]uint{}
	for n := uint(3); n <= max; n += 2 {
		f, ok := ms[n]
		if ok {
			// fがあるなら、nはf/2の倍数なので素数ではない
			delete(ms, n)
		} else {
			// 後に追加する倍数を奇数に限定するための調整
			f = n * 2
		}

		// 新たな倍数を追加
		for newN := n + f; ; newN += f {
			if _, ok := ms[newN]; !ok {
				ms[newN] = f
				break
			}
		}

		if !ok {
			primes = append(primes, n)
		}
	}

	return primes
}
