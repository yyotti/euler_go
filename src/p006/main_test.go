package main

import "testing"

var p006Tests = []struct {
	input    uint
	expected uint
}{
	{input: 0, expected: 0},
	{input: 1, expected: 0},
	{input: 2, expected: 4},
	{input: 3, expected: 22},
	{input: 4, expected: 70},
	{input: 5, expected: 170},
	{input: 10, expected: 2640},
}

func TestP006A(t *testing.T) {
	for _, tt := range p006Tests {
		actual := p006A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP006B(t *testing.T) {
	for _, tt := range p006Tests {
		actual := p006B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP006A_50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p006A(50)
	}
}

func BenchmarkP006B_50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p006B(50)
	}
}

func BenchmarkP006A_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p006A(100)
	}
}

func BenchmarkP006B_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p006B(100)
	}
}

func ExampleP006() {
	main()

	// Output:
	// P006A: 25164150
	// P006B: 25164150
}
