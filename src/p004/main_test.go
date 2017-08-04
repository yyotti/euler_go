package main

import "testing"

var p004Tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 0},
	{input: 1, expected: 9},
	{input: 2, expected: 9009},
}

func TestP004A(t *testing.T) {
	for _, tt := range p004Tests {
		actual := p004A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func TestP004B(t *testing.T) {
	for _, tt := range p004Tests {
		actual := p004B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func TestP004C(t *testing.T) {
	for _, tt := range p004Tests {
		actual := p004C(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func TestP004D(t *testing.T) {
	for _, tt := range p004Tests {
		actual := p004C(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP004A_3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p004A(3)
	}
}

func BenchmarkP004B_3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p004B(3)
	}
}

func BenchmarkP004C_3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p004C(3)
	}
}

func BenchmarkP004D_3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p004D(3)
	}
}

func BenchmarkP004E_3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p004E(3)
	}
}

func BenchmarkP004A_4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p004A(4)
	}
}

func BenchmarkP004B_4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p004B(4)
	}
}

func BenchmarkP004C_4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p004C(4)
	}
}

func BenchmarkP004D_4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p004D(4)
	}
}

func BenchmarkP004E_4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p004E(4)
	}
}

func ExampleP004() {
	main()

	// Output:
	// P004A: 906609
	// P004B: 906609
	// P004C: 906609
	// P004D: 906609
	// P004E: 906609
}
