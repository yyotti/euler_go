package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
// containing over five-thousand first names, begin by sorting it into
// alphabetical order. Then working out the alphabetical value for each name,
// multiply this value by its alphabetical position in the list to obtain a
// name score.
//
// For example, when the list is sorted into alphabetical order, COLIN, which
// is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So,
// COLIN would obtain a score of 938 x 53 = 49714.
//
// What is the total of all the name scores in the file?

const file = "./data/names.txt"

func main() {
	fmt.Printf("P022A: %d\n", p022A(file))
}

func loadNames(file string) ([]string, error) {
	str, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if len(str) == 0 {
		return []string{}, nil
	}

	parts := strings.Split(strings.Trim(string(str), "\n"), ",")
	names := make([]string, len(parts))
	for i, s := range parts {
		if len(s) < 3 || s[0] != '"' || s[len(s)-1] != '"' {
			return nil, errors.New("Invalid format")
		}
		names[i] = s[1 : len(s)-1]
	}

	return names, nil
}

func score(name string) int {
	if len(name) == 0 {
		return 0
	}

	s := 0
	for _, c := range name {
		if c >= 'A' && c <= 'Z' {
			s += int(c-'A') + 1
		}
	}
	return s
}

// つまらんので1つだけ回答を出して終わる
func p022A(file string) int {
	ns, err := loadNames(file)
	if err != nil {
		return -1
	}

	names := make([]string, len(ns))
	for i, s := range ns {
		names[i] = strings.ToUpper(s)
	}
	sort.Strings(names)

	sum := 0
	for i, name := range names {
		sum += score(name) * (i + 1)
	}
	return sum
}
