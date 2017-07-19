package main

import "testing"

var p007Tests = []struct {
	input    uint
	expected uint
}{
	{input: 0, expected: 0},
	{input: 1, expected: 2},
	{input: 2, expected: 3},
	{input: 3, expected: 5},
	{input: 4, expected: 7},
	{input: 5, expected: 11},
	{input: 6, expected: 13},
}

func TestP007A(t *testing.T) {
	for _, tt := range p007Tests {
		actual := p007A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP007B(t *testing.T) {
	for _, tt := range p007Tests {
		actual := p007B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP007A_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p007A(1000)
	}
}

func BenchmarkP007A_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p007A(5000)
	}
}

func BenchmarkP007B_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p007B(1000)
	}
}

func BenchmarkP007B_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p007B(5000)
	}
}

func ExampleP007() {
	main()

	// Output:
	// P007A: 104743
	// P007B: 104743
}
