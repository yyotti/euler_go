package main

import "errors"
import "reflect"
import "testing"

var loadNamesTests = []struct {
	input    string
	expected []string
	err      error
}{
	{input: ""},
	{input: "testdata/not_exists.txt"},
	{input: "testdata/empty.txt", expected: []string{}},
	{input: "testdata/names_err1.txt", err: errors.New("Invalid format")},
	{input: "testdata/names_err2.txt", err: errors.New("Invalid format")},
	{input: "testdata/names_err3.txt", err: errors.New("Invalid format")},
	{input: "testdata/names_err4.txt", err: errors.New("Invalid format")},
	{input: "testdata/names_err5.txt", err: errors.New("Invalid format")},
	{input: "testdata/names1.txt", expected: []string{"MICHEL"}},
	{input: "testdata/names2.txt", expected: []string{"MICHEL", "PETER"}},
}

func TestLoadNames(t *testing.T) {
	for _, tt := range loadNamesTests {
		actual, err := loadNames(tt.input)
		if err != nil {
			if tt.expected != nil {
				t.Errorf("%s: Expected no error but got [%v]", tt.input, err)
			} else if tt.err != nil && err.Error() != tt.err.Error() {
				t.Errorf("%s: Expected error [%v] but got [%v]", tt.input, tt.err, err)
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

var scoreTests = []struct {
	input    string
	expected int
}{
	{input: "", expected: 0},
	{input: "A", expected: 1},
	{input: "b", expected: 0},
	{input: "aBc", expected: 2},
	{input: "Ab-C", expected: 4},
}

func TestScore(t *testing.T) {
	for _, tt := range scoreTests {
		actual := score(tt.input)
		if actual != tt.expected {
			t.Errorf("%s: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var p022Tests = []struct {
	input    string
	expected int
}{
	{input: "", expected: -1},
	{input: "testdata/not_exists.txt", expected: -1},
	{input: "testdata/empty.txt", expected: 0},
	{input: "testdata/names_err1.txt", expected: -1},
	{input: "testdata/names1.txt", expected: 50},
	{input: "testdata/names2.txt", expected: 178},
}

func TestP022A(t *testing.T) {
	for _, tt := range p022Tests {
		actual := p022A(tt.input)
		if actual != tt.expected {
			t.Errorf("%s: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func ExampleP022() {
	main()

	// Output:
	// P022A: 871198282
}
