package common

import (
	"reflect"
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

func TestPrimeGeneratorB_Reset(t *testing.T) {
	gen := newPrimeGeneratorB()
	for i := 0; i < 10; i++ {
		gen.Next()
	}

	gen.Reset()
	if gen.idx != 0 {
		t.Errorf("Expected 0 but got %d", gen.idx)
	}
	if gen.primeMax != 29 {
		t.Errorf("Expected 29 but got %d", gen.primeMax)
	}
	if len(gen.primes) != 10 {
		t.Errorf("Expected 10 but got %d", len(gen.primes))
	}

	n := gen.Next()
	if n != 2 {
		t.Errorf("Expected 2 but got %d", n)
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
	gen := newPrimeGeneratorB()
	for i := 0; i < b.N; i++ {
		gen.Next()
	}

	gen.Reset()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen.Next()
	}
}

var primeFactorsTests = []struct {
	inN      int64
	inGen    []PrimeGenerator
	expected map[int64]int
}{
	{inN: 0, expected: map[int64]int{}},
	{inN: 1, expected: map[int64]int{}},
	{inN: 2, expected: map[int64]int{2: 1}},
	{inN: 3, expected: map[int64]int{3: 1}},
	{inN: 4, expected: map[int64]int{2: 2}},
	{inN: 5, expected: map[int64]int{5: 1}},
	{inN: 6, expected: map[int64]int{2: 1, 3: 1}},
	{inN: 7, expected: map[int64]int{7: 1}},
	{inN: 8, expected: map[int64]int{2: 3}},
	{inN: 9, expected: map[int64]int{3: 2}},
	{inN: 13195, expected: map[int64]int{5: 1, 7: 1, 13: 1, 29: 1}},
	{inN: 13195, inGen: []PrimeGenerator{}, expected: map[int64]int{5: 1, 7: 1, 13: 1, 29: 1}},
	// {{{
	{
		inN: 13195,
		inGen: []PrimeGenerator{
			func() PrimeGenerator {
				g := newPrimeGeneratorB()
				for g.Next() < 30 {
				}
				return g
			}(),
		},
		expected: map[int64]int{5: 1, 7: 1, 13: 1, 29: 1}},
	// }}}
}

func TestPrimeFactors(t *testing.T) {
	for _, tt := range primeFactorsTests {
		var actual map[int64]int
		if len(tt.inGen) != 0 {
			actual = PrimeFactors(tt.inN, tt.inGen...)
		} else {
			actual = PrimeFactors(tt.inN)
		}
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.inN, tt.expected, actual)
		}
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrimeFactors(3628800) // 1 x 2 x ... x 10
	}
}

func BenchmarkPrimeFactors_Cached(b *testing.B) {
	g := newPrimeGeneratorB()
	for g.Next() < 20 {
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrimeFactors(3628800, g) // 1 x 2 x ... x 10
	}
}

var isPrimeTests = []struct {
	input    uint
	expected bool
}{
	{input: 0, expected: false},
	{input: 1, expected: false},
	{input: 2, expected: true},
	{input: 3, expected: true},
	{input: 4, expected: false},
	{input: 5, expected: true},
	{input: 6, expected: false},
	{input: 7, expected: true},
	{input: 8, expected: false},
	{input: 9, expected: false},
	{input: 10, expected: false},
	{input: 11, expected: true},
}

func TestIsPrime(t *testing.T) {
	for _, tt := range isPrimeTests {
		actual := IsPrime(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var sieveTests = []struct {
	input    uint
	expected []uint
}{
	{input: 0, expected: []uint{}},
	{input: 1, expected: []uint{}},
	{input: 2, expected: []uint{2}},
	{input: 3, expected: []uint{2, 3}},
	{input: 4, expected: []uint{2, 3}},
	{input: 5, expected: []uint{2, 3, 5}},
	{input: 6, expected: []uint{2, 3, 5}},
	{input: 7, expected: []uint{2, 3, 5, 7}},
	{input: 50, expected: []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}},
}

func TestSieveA(t *testing.T) {
	for _, tt := range sieveTests {
		actual := sieveA(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func TestSieveB(t *testing.T) {
	for _, tt := range sieveTests {
		actual := sieveB(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkSieveA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieveA(10000)
	}
}

func BenchmarkSieveB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieveB(10000)
	}
}
