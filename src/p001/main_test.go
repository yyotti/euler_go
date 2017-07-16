package main

import "testing"

type p001Input struct {
	max int
	ds  []int
}

var p001Tests = []struct {
	input    p001Input
	expected int
}{
	{
		input:    p001Input{max: 10, ds: []int{3, 5}},
		expected: 3 + 5 + 6 + 9, // 23
	},
	{
		input:    p001Input{max: 20, ds: []int{3, 5}},
		expected: 3 + 5 + 6 + 9 + 10 + 12 + 15 + 18, // 78
	},
	{
		input:    p001Input{max: 10, ds: []int{2, 4}},
		expected: 2 + 4 + 6 + 8, // 20
	},
	{
		input:    p001Input{max: 20, ds: []int{4, 2, 3}},
		expected: 2 + 3 + 4 + 6 + 8 + 9 + 10 + 12 + 14 + 15 + 16 + 18, // 117
	},
	{
		input:    p001Input{max: 20, ds: []int{4, 6, 9}},
		expected: 4 + 6 + 8 + 9 + 12 + 16 + 18, // 73
	},
}

func TestP001A(t *testing.T) {
	for _, tt := range p001Tests {
		actual := p001A(tt.input.max, tt.input.ds)
		if tt.expected != actual {
			t.Errorf("Expected %d but got %d", tt.expected, actual)
		}
	}
}

func TestP001B(t *testing.T) {
	for _, tt := range p001Tests {
		actual := p001B(tt.input.max, tt.input.ds)
		if tt.expected != actual {
			t.Errorf("Expected %d but got %d", tt.expected, actual)
		}
	}
}

func TestP001Z(t *testing.T) {
	for _, tt := range p001Tests[0:2] {
		actual := p001Z(tt.input.max)
		if tt.expected != actual {
			t.Errorf("Expected %d but got %d", tt.expected, actual)
		}
	}
}

func BenchmarkP001A(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p001A(i, []int{3, 5})
	}
}

func BenchmarkP001B(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p001B(i, []int{3, 5})
	}
}

func BenchmarkP001Z(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p001Z(i)
	}
}
