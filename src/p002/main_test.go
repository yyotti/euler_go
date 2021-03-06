package main

import "reflect"
import "testing"

var fibATests = []struct {
	input    int
	expected uint
}{
	{input: 1, expected: 1},
	{input: 2, expected: 2},
	{input: 3, expected: 3},
	{input: 4, expected: 5},
	{input: 5, expected: 8},
	{input: 6, expected: 13},
	{input: 7, expected: 21},
	{input: 8, expected: 34},
	{input: 9, expected: 55},
	{input: 10, expected: 89},
}

func TestFibA(t *testing.T) {
	for _, tt := range fibATests {
		actual := fibA(tt.input)
		if actual != tt.expected {
			t.Errorf("Expected %d but got %d", tt.expected, actual)
		}
	}
}

var p002Tests = []struct {
	input    int
	expected int
}{
	{input: 1, expected: 0},
	{input: 2, expected: 2},
	{input: 3, expected: 2},
	{input: 4, expected: 2},
	{input: 5, expected: 2},
	{input: 6, expected: 2},
	{input: 7, expected: 2},
	{input: 8, expected: 2 + 8},
	{input: 9, expected: 2 + 8},
	{input: 10, expected: 2 + 8},
	{input: 100, expected: 2 + 8 + 34},
	{input: 1000, expected: 2 + 8 + 34 + 144 + 610},
}

func TestP002A(t *testing.T) {
	for _, tt := range p002Tests {
		actual := p002A(tt.input)
		if actual != tt.expected {
			t.Errorf("Expected %d but got %d", tt.expected, actual)
		}
	}
}

var fibBTests = []struct {
	input    int
	expected []int
}{
	{input: 1, expected: []int{0, 1}},
	{input: 2, expected: []int{0, 1, 2}},
	{input: 3, expected: []int{0, 1, 2, 3}},
	{input: 4, expected: []int{0, 1, 2, 3}},
	{input: 5, expected: []int{0, 1, 2, 3, 5}},
	{input: 6, expected: []int{0, 1, 2, 3, 5}},
	{input: 7, expected: []int{0, 1, 2, 3, 5}},
	{input: 8, expected: []int{0, 1, 2, 3, 5, 8}},
	{input: 9, expected: []int{0, 1, 2, 3, 5, 8}},
	{input: 10, expected: []int{0, 1, 2, 3, 5, 8}},
}

func TestFibB(t *testing.T) {
	for _, tt := range fibBTests {
		actual := fibB(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Expected %d but got %d", tt.expected, actual)
		}
	}
}

func TestP002B(t *testing.T) {
	for _, tt := range p002Tests {
		actual := p002B(tt.input)
		if actual != tt.expected {
			t.Errorf("Expected %d but got %d", tt.expected, actual)
		}
	}
}

func TestP002C(t *testing.T) {
	for _, tt := range p002Tests {
		actual := p002C(tt.input)
		if actual != tt.expected {
			t.Errorf("Expected %d but got %d", tt.expected, actual)
		}
	}
}

func TestNewFibonacciGenerator(t *testing.T) {
	actual := newFibonacciGenerator()
	if actual == nil {
		t.Errorf("Fibonacci generator is nil")
	}
	if actual.ch == nil {
		t.Errorf("ch is nil")
	}
}

func TestFibonacciGenerator_start(t *testing.T) {
	gen := &fibonacciGenerator{ch: make(chan uint)}
	go gen.start()
	for _, expected := range []uint{0, 1, 1, 2} {
		actual := <-gen.ch
		if actual != expected {
			t.Errorf("Expected %d but got %d", expected, actual)
			break
		}
	}
}

func TestFibonacciGenerator_Next(t *testing.T) {
	gen := &fibonacciGenerator{ch: make(chan uint)}
	go gen.start()
	for _, expected := range []uint{0, 1, 1, 2, 3, 5} {
		actual := gen.Next()
		if actual != expected {
			t.Errorf("Expected %d but got %d", expected, actual)
			break
		}
	}
}

func TestP002D(t *testing.T) {
	for _, tt := range p002Tests {
		actual := p002D(tt.input)
		if actual != tt.expected {
			t.Errorf("Expected %d but got %d", tt.expected, actual)
		}
	}
}

func BenchmarkP002A(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p002A(10000)
	}
}

func BenchmarkP002B(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p002B(10000)
	}
}

func BenchmarkP002C(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p002C(10000)
	}
}

func BenchmarkP002D(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p002D(10000)
	}
}

func ExampleP002() {
	main()

	// Output:
	// P002A: 4613732
	// P002B: 4613732
	// P002C: 4613732
	// P002D: 4613732
}
