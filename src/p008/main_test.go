package main

import "errors"
import "reflect"
import "testing"

var p008Tests = []struct {
	input    uint
	expected uint
}{
	{input: 0, expected: 0},
	{input: 1, expected: 9},
	{input: 4, expected: 5832},
}

func TestP008A(t *testing.T) {
	for _, tt := range p008Tests {
		actual := p008A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP008B(t *testing.T) {
	for _, tt := range p008Tests {
		actual := p008B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP008C(t *testing.T) {
	for _, tt := range p008Tests {
		actual := p008C(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP008D(t *testing.T) {
	for _, tt := range p008Tests {
		actual := p008D(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var toNumsTests = []struct {
	input    string
	expected []uint
	err      error
}{
	// 正常
	{input: "", expected: []uint{}},
	{input: "1", expected: []uint{1}},
	{input: "12", expected: []uint{1, 2}},
	{input: "012", expected: []uint{0, 1, 2}},

	// 異常
	{input: "a", err: errors.New("Not number: a")},
	{input: "1b", err: errors.New("Not number: b")},
	{input: "c2", err: errors.New("Not number: c")},
	{input: "1#2", err: errors.New("Not number: #")},
	{input: "-200", err: errors.New("Not number: -")},
}

func TestToNums(t *testing.T) {
	for _, tt := range toNumsTests {
		actual, err := splitNums(tt.input)
		if err != nil {
			if tt.err == nil {
				t.Errorf("%s: Error[%s]", tt.input, err.Error())
			} else if err.Error() != tt.err.Error() {
				t.Errorf("%s: Expected error [%s] but got [%s]", tt.input, tt.err.Error(), err.Error())
			}
		} else {
			if tt.err != nil {
				t.Errorf("%s: Expected error but not", tt.input)

			} else if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s: Expected %v but got %v", tt.input, tt.expected, actual)
			}
		}
	}
}

var productTests = []struct {
	input    []uint
	expected uint
}{
	{input: []uint{}, expected: 0},
	{input: []uint{2}, expected: 2},
	{input: []uint{2, 3}, expected: 6},
	{input: []uint{4, 0, 5}, expected: 0},
}

func TestProduct(t *testing.T) {
	for _, tt := range productTests {
		actual := product(tt.input)
		if actual != tt.expected {
			t.Errorf("%v: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP008A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p008A(10)
	}
}

func BenchmarkP008B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p008B(10)
	}
}

func BenchmarkP008C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p008C(10)
	}
}

func BenchmarkP008D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p008D(10)
	}
}

func ExampleP008() {
	main()

	// Output:
	// P008A: 23514624000
	// P008B: 23514624000
	// P008C: 23514624000
	// P008D: 23514624000
}
