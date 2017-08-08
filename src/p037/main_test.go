package main

import "github.com/yyotti/euler_go/src/common"
import "reflect"
import "testing"

func BenchmarkP037A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p037A()
	}
}

func BenchmarkP037B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p037B()
	}
}

var rightTruncatesTests = []struct {
	input    int
	expected []int
}{
	{input: -1, expected: []int{}},
	{input: 0, expected: []int{}},
	{input: 1, expected: []int{}},
	{input: 10, expected: []int{1}},
	{input: 23, expected: []int{2}},
	{input: 102, expected: []int{10, 1}},
	{input: 100, expected: []int{10, 1}},
	{input: 213, expected: []int{21, 2}},
	{input: 210, expected: []int{21, 2}},
}

func TestRightTruncates(t *testing.T) {
	for _, tt := range rightTruncatesTests {
		actual := rightTruncates(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var leftTruncatesTests = []struct {
	input    int
	expected []int
}{
	{input: -1, expected: []int{}},
	{input: 0, expected: []int{}},
	{input: 1, expected: []int{}},
	{input: 10, expected: []int{0}},
	{input: 23, expected: []int{3}},
	{input: 102, expected: []int{2}},
	{input: 100, expected: []int{0}},
	{input: 213, expected: []int{13, 3}},
	{input: 210, expected: []int{10, 0}},
}

func TestLeftTruncates(t *testing.T) {
	for _, tt := range leftTruncatesTests {
		actual := leftTruncates(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var isLeftTruncatablePrimeTests = []struct {
	input    int
	expected bool
}{
	{input: 2, expected: false},
	{input: 3, expected: false},
	{input: 5, expected: false},
	{input: 7, expected: false},
	{input: 11, expected: false},
	{input: 13, expected: true},
	{input: 23, expected: true},
	{input: 31, expected: false},
	{input: 37, expected: true},
	{input: 41, expected: false},
	{input: 43, expected: true},
	{input: 223, expected: true},
	{input: 737, expected: true},
}

func TestIsLeftTruncatablePrime(t *testing.T) {
	gen := common.NewPrimeGenerator()
	for _, tt := range isLeftTruncatablePrimeTests {
		actual := isLeftTruncatablePrime(tt.input, gen)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var isRightTruncatablePrimeTests = []struct {
	input    int
	expected bool
}{
	{input: 2, expected: false},
	{input: 3, expected: false},
	{input: 5, expected: false},
	{input: 7, expected: false},
	{input: 11, expected: false},
	{input: 13, expected: false},
	{input: 23, expected: true},
	{input: 31, expected: true},
	{input: 37, expected: true},
	{input: 41, expected: false},
	{input: 43, expected: false},
	{input: 233, expected: true},
	{input: 737, expected: true},
}

func TestIsRightTruncatablePrime(t *testing.T) {
	gen := common.NewPrimeGenerator()
	for _, tt := range isRightTruncatablePrimeTests {
		actual := isRightTruncatablePrime(tt.input, gen)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var isTruncatablePrimeTests = []struct {
	input    int
	expected bool
}{
	{input: 2, expected: false},
	{input: 3, expected: false},
	{input: 5, expected: false},
	{input: 7, expected: false},
	{input: 11, expected: false},
	{input: 13, expected: false},
	{input: 23, expected: true},
	{input: 31, expected: false},
	{input: 37, expected: true},
	{input: 41, expected: false},
	{input: 43, expected: false},
	{input: 127, expected: false},
	{input: 737, expected: true},
}

func TestIsTruncatablePrime(t *testing.T) {
	gen := common.NewPrimeGenerator()
	for _, tt := range isTruncatablePrimeTests {
		actual := isTruncatablePrime(tt.input, gen)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

type x struct{}

var setAddTests = []struct {
	s        set
	input    int
	expected set
}{
	{s: set{}, input: 1, expected: set{1: x{}}},
	{s: set{1: x{}}, input: 2, expected: set{2: x{}, 1: x{}}},
	{s: set{2: x{}}, input: 2, expected: set{2: x{}}},
}

func TestSet_Add(t *testing.T) {
	for i, tt := range setAddTests {
		tt.s.add(tt.input)
		actual := tt.s
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("TEST %d: Expected %v but got %v", i, tt.expected, actual)
		}
	}
}

var setUnionTests = []struct {
	s        set
	input    set
	expected set
}{
	{s: set{}, input: set{}, expected: set{}},
	{s: set{}, input: set{1: x{}}, expected: set{1: x{}}},
	{s: set{2: x{}}, input: set{2: x{}}, expected: set{2: x{}}},
	{s: set{2: x{}}, input: set{1: x{}}, expected: set{2: x{}, 1: x{}}},
	{s: set{2: x{}, 1: x{}}, input: set{2: x{}}, expected: set{2: x{}, 1: x{}}},
	{s: set{1: x{}}, input: set{1: x{}, 2: x{}}, expected: set{2: x{}, 1: x{}}},
	{s: set{1: x{}, 2: x{}}, input: set{2: x{}, 3: x{}}, expected: set{2: x{}, 1: x{}, 3: x{}}},
}

func TestSet_Union(t *testing.T) {
	for i, tt := range setUnionTests {
		actual := tt.s.union(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("TEST %d: Expected %v but got %v", i, tt.expected, actual)
		}
	}
}

var setIntersectTests = []struct {
	s        set
	input    set
	expected set
}{
	{s: set{}, input: set{}, expected: set{}},
	{s: set{}, input: set{1: x{}}, expected: set{}},
	{s: set{2: x{}}, input: set{2: x{}}, expected: set{2: x{}}},
	{s: set{2: x{}}, input: set{1: x{}}, expected: set{}},
	{s: set{2: x{}, 1: x{}}, input: set{2: x{}}, expected: set{2: x{}}},
	{s: set{1: x{}}, input: set{1: x{}, 2: x{}}, expected: set{1: x{}}},
	{s: set{2: x{}, 3: x{}}, input: set{2: x{}, 3: x{}}, expected: set{2: x{}, 3: x{}}},
	{s: set{1: x{}, 2: x{}, 3: x{}}, input: set{4: x{}, 2: x{}, 3: x{}}, expected: set{2: x{}, 3: x{}}},
}

func TestSet_Intersect(t *testing.T) {
	for i, tt := range setIntersectTests {
		actual := tt.s.intersect(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("TEST %d: Expected %v but got %v", i, tt.expected, actual)
		}
	}
}

func ExampleP037() {
	main()

	// Output:
	// P037A: 748317
	// P037B: 748317
}
