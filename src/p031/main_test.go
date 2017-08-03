package main

import "testing"

var p031Tests = []struct {
	inP      int
	inCoins  []int
	expected int
}{
	{inP: -1, inCoins: []int{1}, expected: 0},
	{inP: 0, inCoins: []int{1, 2}, expected: 0},
	{inP: 1, inCoins: []int{}, expected: 0},
	{inP: 1, inCoins: []int{2}, expected: 0},
	{inP: 1, inCoins: []int{2, 1}, expected: 1},
	{inP: 2, inCoins: []int{2, 1}, expected: 2},
	{inP: 3, inCoins: []int{2, 1}, expected: 2},
	{inP: 4, inCoins: []int{2, 1}, expected: 3},
	{inP: 4, inCoins: []int{3, 2, 1}, expected: 4},
}

func TestP031A(t *testing.T) {
	for _, tt := range p031Tests {
		actual := p031A(tt.inP, tt.inCoins)
		if actual != tt.expected {
			t.Errorf("%d,%v: Expected %d but got %d", tt.inP, tt.inCoins, tt.expected, actual)
		}
	}
}

func ExampleP031() {
	main()

	// Output:
	// P031A: 73682
}
