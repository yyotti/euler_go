package main

import "testing"

var p036Tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 0},
	{input: 1, expected: 0},
	{input: 2, expected: 1},
	{input: 10, expected: 25},
	{input: 50, expected: 58},
	{input: 100, expected: 157},
}

func TestP036A(t *testing.T) {
	for _, tt := range p036Tests {
		actual := p036A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP036B(t *testing.T) {
	for _, tt := range p036Tests {
		actual := p036B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP036C(t *testing.T) {
	for _, tt := range p036Tests {
		actual := p036C(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP036A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p036A(10000)
	}
}

func BenchmarkP036B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p036B(10000)
	}
}

func BenchmarkP036C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p036C(10000)
	}
}

func ExampleP036() {
	main()

	// Output:
	// P036A: 872187
	// P036B: 872187
	// P036C: 872187
}
