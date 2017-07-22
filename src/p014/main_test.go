package main

import (
	"reflect"
	"testing"
)

var collatzTests = []struct {
	input    uint
	expected uint
}{
	{input: 1, expected: 4},
	{input: 2, expected: 1},
	{input: 3, expected: 10},
	{input: 4, expected: 2},
	{input: 5, expected: 16},
	{input: 6, expected: 3},
}

func TestCollatz(t *testing.T) {
	for _, tt := range collatzTests {
		actual := collatz(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var collatzChainTests = []struct {
	input    uint
	expected []uint
}{
	{input: 0, expected: []uint{}},
	{input: 1, expected: []uint{1}},
	{input: 2, expected: []uint{2, 1}},
	{input: 3, expected: []uint{3, 10, 5, 16, 8, 4, 2, 1}},
	{input: 13, expected: []uint{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}},
}

func TestCollatzChain(t *testing.T) {
	for _, tt := range collatzChainTests {
		actual := collatzChain(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkCollatzChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		collatzChain(100000)
	}
}

var collatzCntTests = []struct {
	input    uint
	expected uint
}{
	{input: 0, expected: 0},
	{input: 1, expected: 1},
	{input: 2, expected: 2},
	{input: 3, expected: 8},
	{input: 13, expected: 10},
}

func TestCollatzCnt(t *testing.T) {
	for _, tt := range collatzCntTests {
		actual := collatzCnt(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkCollatzCnt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		collatzCnt(100000)
	}
}

var collatzCntCachedTests = []struct {
	inStart  uint
	inCache  map[uint]uint
	expected uint
}{
	{inStart: 0, inCache: map[uint]uint{}, expected: 0},
	{inStart: 0, inCache: map[uint]uint{1: 1}, expected: 0},
	{inStart: 0, inCache: map[uint]uint{0: 2}, expected: 0},
	{inStart: 1, inCache: map[uint]uint{}, expected: 1},
	{inStart: 1, inCache: map[uint]uint{0: 100}, expected: 1},
	{inStart: 1, inCache: map[uint]uint{1: 0}, expected: 1},
	{inStart: 2, inCache: map[uint]uint{}, expected: 2},
	{inStart: 2, inCache: map[uint]uint{0: 100}, expected: 2},
	{inStart: 2, inCache: map[uint]uint{1: 3}, expected: 2},
	{inStart: 2, inCache: map[uint]uint{2: 4}, expected: 4},
	{inStart: 3, inCache: map[uint]uint{}, expected: 8},
	{inStart: 3, inCache: map[uint]uint{0: 100}, expected: 8},
	{inStart: 3, inCache: map[uint]uint{1: 200}, expected: 8},
	{inStart: 3, inCache: map[uint]uint{3: 200}, expected: 200},
	{inStart: 3, inCache: map[uint]uint{5: 50}, expected: 52},
	{inStart: 13, inCache: map[uint]uint{}, expected: 10},
	{inStart: 13, inCache: map[uint]uint{16: 30}, expected: 35},
}

func TestCollatzCntCached(t *testing.T) {
	for _, tt := range collatzCntCachedTests {
		actual := collatzCntCached(tt.inStart, tt.inCache)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.inStart, tt.expected, actual)
		}
	}
}

func BenchmarkCollatzCntCached_Nocache(b *testing.B) {
	cache := map[uint]uint{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		collatzCntCached(100000, cache)
	}
}

func BenchmarkCollatzCntCached_Cache(b *testing.B) {
	cache := map[uint]uint{
		1: 1, 2: 2, 4: 3, 8: 4, 16: 5,
		5: 6, 10: 7, 3: 8, 20: 8,
		32: 6, 64: 7, 21: 8, 128: 8,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		collatzCntCached(100000, cache)
	}
}

var p014Tests = []struct {
	input    uint
	expected uint
}{
	{input: 0, expected: 0},
	{input: 1, expected: 1},
	{input: 2, expected: 2},
	{input: 3, expected: 3},
	{input: 4, expected: 3},
	{input: 10, expected: 9},
}

func TestP014A(t *testing.T) {
	for _, tt := range p014Tests {
		actual := p014A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func TestP014B(t *testing.T) {
	for _, tt := range p014Tests {
		actual := p014B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func TestP014C(t *testing.T) {
	for _, tt := range p014Tests {
		actual := p014C(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP014A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014A(10000)
	}
}

func BenchmarkP014B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014B(10000)
	}
}

func BenchmarkP014C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014C(10000)
	}
}

func ExampleP014() {
	main()

	// Output:
	// P014A: 837799
	// P014B: 837799
	// P014C: 837799
}
