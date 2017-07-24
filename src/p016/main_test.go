package main

import "testing"

var pow2Tests = []struct {
	input    int
	expected string
}{
	{input: 0, expected: "1"},
	{input: 1, expected: "2"},
	{input: 2, expected: "4"},
	{input: 3, expected: "8"},
	{input: 4, expected: "16"},
	{input: 5, expected: "32"},
	{input: 6, expected: "64"},
	{input: 7, expected: "128"},
	{input: 8, expected: "256"},
	{input: 9, expected: "512"},
	{input: 10, expected: "1024"},
}

func TestPow2(t *testing.T) {
	for _, tt := range pow2Tests {
		actual := pow2(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %s but got %s", tt.input, tt.expected, actual)
		}
	}
}

var p016Tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 1},
	{input: 1, expected: 2},
	{input: 2, expected: 4},
	{input: 3, expected: 8},
	{input: 4, expected: 7},
	{input: 5, expected: 5},
	{input: 6, expected: 10},
	{input: 7, expected: 11},
	{input: 15, expected: 26},
}

func TestP016A(t *testing.T) {
	for _, tt := range p016Tests {
		actual := p016A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP016B(t *testing.T) {
	for _, tt := range p016Tests {
		actual := p016B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP016C(t *testing.T) {
	for _, tt := range p016Tests {
		actual := p016C(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP016A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p016A(100)
	}
}

func BenchmarkP016B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p016B(100)
	}
}

func BenchmarkP016C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p016C(100)
	}
}

func ExampleP016() {
	main()

	// Output:
	// P016A: 1366
	// P016B: 1366
	// P016C: 1366
}
