package main

import (
	"reflect"
	"testing"
)

var uint64pIsPrimeTests = []struct {
	input    intp
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

func TestUint64p_isPrime(t *testing.T) {
	for _, tt := range uint64pIsPrimeTests {
		actual := tt.input.isPrime()
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
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
		actual := primeFactors(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var p003Tests = []struct {
	input    int64
	expected int64
}{
	{input: 0, expected: 1},
	{input: 1, expected: 1},
	{input: 2, expected: 2},
	{input: 3, expected: 3},
	{input: 4, expected: 2},
	{input: 5, expected: 5},
	{input: 6, expected: 3},
	{input: 13195, expected: 29},
}

func TestP003A(t *testing.T) {
	for _, tt := range p003Tests {
		actual := p003A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP003B(t *testing.T) {
	for _, tt := range p003Tests {
		actual := p003B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP003C(t *testing.T) {
	for _, tt := range p003Tests {
		actual := p003C(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP003D(t *testing.T) {
	for _, tt := range p003Tests {
		actual := p003D(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP003A(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p003A(997009) // 997*997
	}
}

func BenchmarkP003B(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p003B(997009) // 997*997
	}
}

func BenchmarkP003C(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p003C(997009) // 997*997
	}
}

func BenchmarkP003D(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p003D(997009) // 997*997
	}
}

func ExampleP003() {
	main()

	// Output:
	// P003A: 6857
	// P003B: 6857
	// P003C: 6857
	// P003D: 6857
}
