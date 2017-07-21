package main

import "testing"

var p008Tests = []struct {
	input    uint
	expected uint
}{
	{input: 0, expected: 0},
	{input: 1, expected: 9},
	{input: 4, expected: 5832},
}

func TestP008A(t *testing.T) {
	for _, tt := range p008Tests {
		actual := p008A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP008B(t *testing.T) {
	for _, tt := range p008Tests {
		actual := p008B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP008C(t *testing.T) {
	for _, tt := range p008Tests {
		actual := p008C(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP008D(t *testing.T) {
	for _, tt := range p008Tests {
		actual := p008D(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var productTests = []struct {
	input    []uint
	expected uint
}{
	{input: []uint{}, expected: 0},
	{input: []uint{2}, expected: 2},
	{input: []uint{2, 3}, expected: 6},
	{input: []uint{4, 0, 5}, expected: 0},
}

func TestProduct(t *testing.T) {
	for _, tt := range productTests {
		actual := product(tt.input)
		if actual != tt.expected {
			t.Errorf("%v: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP008A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p008A(10)
	}
}

func BenchmarkP008B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p008B(10)
	}
}

func BenchmarkP008C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p008C(10)
	}
}

func BenchmarkP008D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p008D(10)
	}
}

func ExampleP008() {
	main()

	// Output:
	// P008A: 23514624000
	// P008B: 23514624000
	// P008C: 23514624000
	// P008D: 23514624000
}
