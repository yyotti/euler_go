package main

import "reflect"
import "testing"

var fibATests = []struct {
	input    []int
	expected int
}{
	{input: []int{1, 2, 1}, expected: 1},
	{input: []int{1, 2, 2}, expected: 2},
	{input: []int{1, 2, 3}, expected: 3},
	{input: []int{1, 2, 4}, expected: 5},
	{input: []int{1, 2, 5}, expected: 8},
	{input: []int{1, 2, 6}, expected: 13},
	{input: []int{1, 2, 7}, expected: 21},
	{input: []int{1, 2, 8}, expected: 34},
	{input: []int{1, 2, 9}, expected: 55},
	{input: []int{1, 2, 10}, expected: 89},
}

func TestFibA(t *testing.T) {
	for _, tt := range fibATests {
		actual := fibA(tt.input[0], tt.input[1], tt.input[2])
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
	input    []int
	expected []int
}{
	{input: []int{1, 2, 1}, expected: []int{0, 1}},
	{input: []int{1, 2, 2}, expected: []int{0, 1, 2}},
	{input: []int{1, 2, 3}, expected: []int{0, 1, 2, 3}},
	{input: []int{1, 2, 4}, expected: []int{0, 1, 2, 3}},
	{input: []int{1, 2, 5}, expected: []int{0, 1, 2, 3, 5}},
	{input: []int{1, 2, 6}, expected: []int{0, 1, 2, 3, 5}},
	{input: []int{1, 2, 7}, expected: []int{0, 1, 2, 3, 5}},
	{input: []int{1, 2, 8}, expected: []int{0, 1, 2, 3, 5, 8}},
	{input: []int{1, 2, 9}, expected: []int{0, 1, 2, 3, 5, 8}},
	{input: []int{1, 2, 10}, expected: []int{0, 1, 2, 3, 5, 8}},
}

func TestFibB(t *testing.T) {
	for _, tt := range fibBTests {
		actual := fibB(tt.input[0], tt.input[1], tt.input[2])
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

var newFibonacciGeneratorTests = []struct {
	input    []int
	expected *fibonacciGenerator
}{
	{[]int{0, 1}, &fibonacciGenerator{a1: 0, a2: 1}},
	{[]int{1, -2}, &fibonacciGenerator{a1: 1, a2: -2}},
}

func TestNewFibonacciGenerator(t *testing.T) {
	for _, tt := range newFibonacciGeneratorTests {
		actual := newFibonacciGenerator(tt.input[0], tt.input[1])
		if actual == nil {
			t.Errorf("Fibonacci generator is nil")
		}
		if actual.a1 != tt.expected.a1 {
			t.Errorf("a1: Expected %d but got %d", tt.expected.a1, actual.a1)
		}
		if actual.a2 != tt.expected.a2 {
			t.Errorf("a2: Expected %d but got %d", tt.expected.a2, actual.a2)
		}
		if actual.ch == nil {
			t.Errorf("ch is nil")
		}
	}
}

var fibonacciGeneratorStartTests = []struct {
	input    *fibonacciGenerator
	expected []int
}{
	{&fibonacciGenerator{a1: 0, a2: 1, ch: make(chan int)}, []int{0, 1, 1, 2}},
	{&fibonacciGenerator{a1: 1, a2: -2, ch: make(chan int)}, []int{1, -2, -1, -3}},
}

func TestFibonacciGenerator_start(t *testing.T) {
	for _, tt := range fibonacciGeneratorStartTests {
		gen := tt.input
		go gen.start()
		for _, expected := range tt.expected {
			actual := <-gen.ch
			if actual != expected {
				t.Errorf("Expected %d but got %d", expected, actual)
				break
			}
		}
	}
}

var fibonacciGeneratorNextTests = []struct {
	input    *fibonacciGenerator
	expected []int
}{
	{&fibonacciGenerator{a1: 1, a2: 1, ch: make(chan int)}, []int{1, 1, 2, 3}},
	{&fibonacciGenerator{a1: 2, a2: 4, ch: make(chan int)}, []int{2, 4, 6, 10}},
}

func TestFibonacciGenerator_Next(t *testing.T) {
	for _, tt := range fibonacciGeneratorNextTests {
		gen := tt.input
		go gen.start()
		for _, expected := range tt.expected {
			actual := gen.Next()
			if actual != expected {
				t.Errorf("Expected %d but got %d", expected, actual)
				break
			}
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
