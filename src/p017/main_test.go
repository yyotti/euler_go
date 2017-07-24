package main

import "testing"

var intStringTests = []struct {
	input    Int
	expected string
}{
	{input: Int(1), expected: "one"},
	{input: Int(2), expected: "two"},
	{input: Int(3), expected: "three"},
	{input: Int(4), expected: "four"},
	{input: Int(5), expected: "five"},
	{input: Int(6), expected: "six"},
	{input: Int(7), expected: "seven"},
	{input: Int(8), expected: "eight"},
	{input: Int(9), expected: "nine"},
	{input: Int(10), expected: "ten"},
	{input: Int(11), expected: "eleven"},
	{input: Int(12), expected: "twelve"},
	{input: Int(13), expected: "thirteen"},
	{input: Int(14), expected: "fourteen"},
	{input: Int(15), expected: "fifteen"},
	{input: Int(16), expected: "sixteen"},
	{input: Int(17), expected: "seventeen"},
	{input: Int(18), expected: "eighteen"},
	{input: Int(19), expected: "nineteen"},
	{input: Int(20), expected: "twenty"},
	{input: Int(21), expected: "twenty one"},
	{input: Int(30), expected: "thirty"},
	{input: Int(32), expected: "thirty two"},
	{input: Int(40), expected: "forty"},
	{input: Int(43), expected: "forty three"},
	{input: Int(50), expected: "fifty"},
	{input: Int(54), expected: "fifty four"},
	{input: Int(60), expected: "sixty"},
	{input: Int(65), expected: "sixty five"},
	{input: Int(70), expected: "seventy"},
	{input: Int(76), expected: "seventy six"},
	{input: Int(80), expected: "eighty"},
	{input: Int(87), expected: "eighty seven"},
	{input: Int(90), expected: "ninety"},
	{input: Int(98), expected: "ninety eight"},
	{input: Int(100), expected: "one hundred"},
	{input: Int(111), expected: "one hundred and eleven"},
	{input: Int(200), expected: "two hundred"},
	{input: Int(221), expected: "two hundred and twenty one"},
	{input: Int(1000), expected: "one thousand"},
}

func TestInt_String(t *testing.T) {
	for _, tt := range intStringTests {
		actual := tt.input.String()
		if actual != tt.expected {
			t.Errorf("%d: Expected \"%s\" but got \"%s\"", int(tt.input), tt.expected, actual)
		}
	}
}

var p017Tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 0},
	{input: 1, expected: len("one")},
	{input: 2, expected: len("one") + len("two")},
	{input: 3, expected: len("one") + len("two") + len("three")},
}

func TestP017A(t *testing.T) {
	for _, tt := range p017Tests {
		actual := p017A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func ExampleP017() {
	main()

	// Output:
	// P017A: 21124
}
