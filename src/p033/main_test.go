package main

import "testing"

func BenchmarkP033A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p033A()
	}
}

func BenchmarkP033B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p033B()
	}
}

func ExampleP033() {
	main()

	// Output:
	// P033A: 100
	// P033B: 100
}
