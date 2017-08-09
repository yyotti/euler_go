package common

import (
	"sort"
)

// DigitCount : nの桁数
func DigitCount(n uint64) int {
	return len(Digits(n))
}

// IsPandigital : パンデジタル数判定
func IsPandigital(n int, m int) bool {
	if n < 1 || m < 1 {
		return false
	}

	digits := Digits(uint64(n))
	if len(digits) != m {
		return false
	}

	sort.Ints(digits)
	for i, d := range digits {
		if i+1 != d {
			return false
		}
	}

	return true
}

// Digits : 1桁ごとに分割
func Digits(n uint64) []int {
	if n < 0 {
		return []int{}
	} else if n == 0 {
		return []int{0}
	}

	digits := []int{}
	for n > 0 {
		digits = append(digits, int(n%10))
		n /= 10
	}
	for i := 0; i < len(digits)/2; i++ {
		j := len(digits) - i - 1
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}
