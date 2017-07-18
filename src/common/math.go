package common

import (
	"math"
)

// Gcd : 最大公約数
func Gcd(m, n int) int {
	if m < 0 || n < 0 {
		return Gcd(int(math.Abs(float64(m))), int(math.Abs(float64(n))))
	}
	if m < n {
		return Gcd(n, m)
	}
	if n == 0 {
		return m
	}

	return Gcd(n, m%n)
}

// Lcm : 最小公倍数
func Lcm(m, n uint) uint {
	return m * n / uint(Gcd(int(m), int(n)))
}
