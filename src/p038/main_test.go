package main

import "testing"

func BenchmarkP038A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p038A()
	}
}

func BenchmarkP038B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p038B()
	}
}

func ExampleP038() {
	main()

	// Output:
	// P038A: 932718654
	// P038B: 932718654
}
