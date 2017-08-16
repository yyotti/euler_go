package main

import "testing"

func BenchmarkP041A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p041A()
	}
}

func BenchmarkP041B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p041B()
	}
}

func ExampleP041() {
	main()

	// Output:
	// P041A: 7652413
	// P041B: 7652413
}
