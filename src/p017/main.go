package main

import "fmt"
import "strings"

// If the numbers 1 to 5 are written out in words: one, two, three, four,
// five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out
// in words, how many letters would be used?
//
//
// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
// forty-two) contains 23 letters and 115 (one hundred and fifteen) contains
// 20 letters. The use of "and" when writing out numbers is in compliance
// with British usage.

const max = 1000

func main() {
	fmt.Printf("P017A: %d\n", p017A(max))
}

// つまらんので1つだけ回答を出して終わる
func p017A(max uint) uint {
	sum := uint(0)
	for i := uint(1); i <= max; i++ {
		s := Int(i).String()
		for _, t := range strings.Split(s, " ") {
			sum += uint(len(t))
		}
	}

	return sum
}

// Int : 専用int
type Int uint

// 面倒なので上限1000。速度も無視。
func (i Int) String() string {
	if i == 1000 {
		return "one thousand"
	}

	k := uint(i)
	str := ""
	if k >= 100 {
		str += Int(k/100).String() + " hundred"
		k %= 100
		if k == 0 {
			return str
		}
		str += " and "
	}
	if k >= 20 {
		l := k / 10
		switch l {
		case 2:
			str += "twenty"
		case 3:
			str += "thirty"
		case 4:
			str += "forty"
		case 5:
			str += "fifty"
		case 6:
			str += "sixty"
		case 7:
			str += "seventy"
		case 8:
			str += "eighty"
		case 9:
			str += "ninety"
		}
		k %= 10
		if k == 0 {
			return str
		}
		str += " "
	}

	if k > 0 {
		switch k {
		case 1:
			str += "one"
		case 2:
			str += "two"
		case 3:
			str += "three"
		case 4:
			str += "four"
		case 5:
			str += "five"
		case 6:
			str += "six"
		case 7:
			str += "seven"
		case 8:
			str += "eight"
		case 9:
			str += "nine"
		case 10:
			str += "ten"
		case 11:
			str += "eleven"
		case 12:
			str += "twelve"
		case 13:
			str += "thirteen"
		case 14:
			str += "fourteen"
		case 15:
			str += "fifteen"
		case 16:
			str += "sixteen"
		case 17:
			str += "seventeen"
		case 18:
			str += "eighteen"
		case 19:
			str += "nineteen"
		}
	}

	return str
}
