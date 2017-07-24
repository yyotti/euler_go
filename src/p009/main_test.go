package main

import (
	"testing"
)

var p009Tests = []struct {
	input    int
	expected []int
}{
	{input: -1, expected: []int{}},
	{input: 0, expected: []int{}},
	{input: 1, expected: []int{}},
	{input: 3 + 4 + 5, expected: []int{3 * 4 * 5}},
	{input: 13, expected: []int{}},
	{input: 5 + 12 + 13, expected: []int{5 * 12 * 13}},
	{input: 8 + 15 + 17, expected: []int{8 * 15 * 17}},
	{input: 6 + 8 + 10, expected: []int{6 * 8 * 10}},
	{input: 15 + 36 + 39, expected: []int{9 * 40 * 41, 15 * 36 * 39}},
}

func TestP009A(t *testing.T) {
	for _, tt := range p009Tests {
		actual := p009A(tt.input)
		if len(actual) != len(tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		} else {
			for _, n := range actual {
				found := false
				for _, e := range tt.expected {
					if e == n {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
				}
			}
		}
	}
}

func TestP009B(t *testing.T) {
	for _, tt := range p009Tests {
		actual := p009B(tt.input)
		if len(actual) != len(tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		} else {
			for _, n := range actual {
				found := false
				for _, e := range tt.expected {
					if e == n {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
				}
			}
		}
	}
}

func BenchmarkP009A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p009A(1000)
	}
}

func BenchmarkP009B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p009B(1000)
	}
}

func ExampleP009() {
	main()

	// Output:
	// P009A: 31875000
	// P009B: 31875000
}
