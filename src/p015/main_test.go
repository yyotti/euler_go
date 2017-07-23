package main

import "testing"

var p015Tests = []struct {
	inputX   uint
	inputY   uint
	expected uint64
}{
	{inputX: 0, inputY: 0, expected: 1},
	{inputX: 1, inputY: 0, expected: 1},
	{inputX: 0, inputY: 1, expected: 1},
	{inputX: 1, inputY: 1, expected: 2},
	{inputX: 0, inputY: 2, expected: 1},
	{inputX: 2, inputY: 0, expected: 1},
	{inputX: 1, inputY: 2, expected: 3},
	{inputX: 2, inputY: 1, expected: 3},
	{inputX: 2, inputY: 2, expected: 6},
	{inputX: 0, inputY: 3, expected: 1},
	{inputX: 3, inputY: 0, expected: 1},
	{inputX: 1, inputY: 3, expected: 4},
	{inputX: 3, inputY: 1, expected: 4},
	{inputX: 2, inputY: 3, expected: 10},
	{inputX: 3, inputY: 2, expected: 10},
	{inputX: 3, inputY: 3, expected: 20},
}

func TestP015A(t *testing.T) {
	for _, tt := range p015Tests {
		actual := p015A(tt.inputX, tt.inputY)
		if actual != tt.expected {
			t.Errorf("(%d, %d): Expected %d but got %d", tt.inputX, tt.inputY, tt.expected, actual)
		}
	}
}

func TestP015B(t *testing.T) {
	for _, tt := range p015Tests {
		actual := p015B(tt.inputX, tt.inputY)
		if actual != tt.expected {
			t.Errorf("(%d, %d): Expected %d but got %d", tt.inputX, tt.inputY, tt.expected, actual)
		}
	}
}

func TestP015C(t *testing.T) {
	for _, tt := range p015Tests {
		actual := p015C(tt.inputX, tt.inputY)
		if actual != tt.expected {
			t.Errorf("(%d, %d): Expected %d but got %d", tt.inputX, tt.inputY, tt.expected, actual)
		}
	}
}

func TestP015D(t *testing.T) {
	for _, tt := range p015Tests {
		actual := p015D(tt.inputX, tt.inputY)
		if actual != tt.expected {
			t.Errorf("(%d, %d): Expected %d but got %d", tt.inputX, tt.inputY, tt.expected, actual)
		}
	}
}

func BenchmarkP015A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p015A(10, 10)
	}
}

func BenchmarkP015B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p015B(10, 10)
	}
}

func BenchmarkP015C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p015C(10, 10)
	}
}

func BenchmarkP015D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p015D(10, 10)
	}
}

func ExampleP015() {
	main()

	// Output:
	// P015A(size 10): 184756
	// P015B: 137846528820
	// P015C(size 15): 155117520
	// P015D: 137846528820
}
