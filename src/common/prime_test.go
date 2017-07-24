package common

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewPrimeGenerator(t *testing.T) {
	actual := NewPrimeGenerator()
	if actual == nil {
		t.Errorf("Fibonacci generator is nil")
	}
	switch actual.(type) {
	case *primeGeneratorB:
		if actual.(*primeGeneratorB).idx != 0 {
			t.Errorf("Expected 0 but got %d", actual.(*primeGeneratorB).idx)
		}
	default:
		t.Errorf("Not a `primeGeneratorB` instance")
	}
}

func TestPrimeGeneratorA_start(t *testing.T) {
	gen := &primeGeneratorA{ch: make(chan uint)}
	go gen.start()
	for _, expected := range []uint{2, 3, 5, 7} {
		actual := <-gen.ch
		if actual != expected {
			t.Errorf("Expected %d but got %d", expected, actual)
			break
		}
	}
}

func newPrimeGeneratorB() *primeGeneratorB {
	gen := &primeGeneratorB{}
	gen.start()

	// キャッシュをクリアする
	primes = make([]uint, 0, 1000)
	primes = append(primes, 2, 3)
	primeMax = 3

	mutex = &sync.Mutex{}

	return gen
}

func TestPrimeGeneratorB_start(t *testing.T) {
	gen := newPrimeGeneratorB()

	if gen.idx != 0 {
		t.Errorf("Expected %d but got %d", 0, gen.idx)
	}
}

func TestPrimeGeneratorA_Next(t *testing.T) {
	gen := &primeGeneratorA{ch: make(chan uint)}
	go gen.start()
	for _, expected := range []uint{2, 3, 5, 7} {
		actual := gen.Next()
		if actual != expected {
			t.Errorf("Expected %d but got %d", expected, actual)
			break
		}
	}
}

func TestPrimeGeneratorB_Next(t *testing.T) {
	gen := newPrimeGeneratorB()
	for _, expected := range []uint{2, 3, 5, 7} {
		actual := gen.Next()
		if actual != expected {
			t.Errorf("Expected %d but got %d", expected, actual)
			break
		}
	}
}

func BenchmarkPrimeGeneratorA(b *testing.B) {
	gen := &primeGeneratorA{ch: make(chan uint)}
	go gen.start()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen.Next()
	}
}

func BenchmarkPrimeGeneratorB(b *testing.B) {
	gen := newPrimeGeneratorB()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen.Next()
	}
}

func BenchmarkPrimeGeneratorB_cached(b *testing.B) {
	// 10000未満の素数をあらかじめキャッシュしておく
	b.N = 10000
	genPre := newPrimeGeneratorB()
	for i := 0; i < b.N; i++ {
		genPre.Next()
	}

	gen := &primeGeneratorB{}
	gen.start()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen.Next()
	}
}

var primeFactorsTests = []struct {
	input    int64
	expected map[int64]int
}{
	{input: 0, expected: map[int64]int{}},
	{input: 1, expected: map[int64]int{}},
	{input: 2, expected: map[int64]int{2: 1}},
	{input: 3, expected: map[int64]int{3: 1}},
	{input: 4, expected: map[int64]int{2: 2}},
	{input: 5, expected: map[int64]int{5: 1}},
	{input: 6, expected: map[int64]int{2: 1, 3: 1}},
	{input: 7, expected: map[int64]int{7: 1}},
	{input: 8, expected: map[int64]int{2: 3}},
	{input: 9, expected: map[int64]int{3: 2}},
	{input: 13195, expected: map[int64]int{5: 1, 7: 1, 13: 1, 29: 1}},
}

func TestPrimeFactors(t *testing.T) {
	for _, tt := range primeFactorsTests {
		actual := PrimeFactors(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	// キャッシュをリセット
	newPrimeGeneratorB()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrimeFactors(3628800) // 1 x 2 x ... x 10
	}
}
