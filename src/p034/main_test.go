package main

import "testing"

func BenchmarkP034A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p034A()
	}
}

func BenchmarkP034B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p034B()
	}
}

func ExampleP034() {
	main()

	// Output:
	// P034A: 40730
	// P034B: 40730
}
