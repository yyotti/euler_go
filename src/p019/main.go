package main

import "fmt"
import "time"

// You are given the following information, but you may prefer to do some research for yourself.
//
// ・1 Jan 1900 was a Monday.
// ・Thirty days has September,
//   April, June and November.
//   All the rest have thirty-one,
//   Saving February alone,
//   Which has twenty-eight, rain or shine.
//   And on leap years, twenty-nine.
// ・A leap year occurs on any year evenly divisible by 4, but not on
//   a century unless it is divisible by 400.
//
// How many Sundays fell on the first of the month during the twentieth
// century (1 Jan 1901 to 31 Dec 2000)?
func main() {
	fmt.Printf("P019A: %d\n", p019A())
}

func isLeapYear(y int) bool {
	return y%400 == 0 || y%4 == 0 && y%100 != 0
}

func dates(y, m int) int {
	switch m {
	case 2:
		if isLeapYear(y) {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

func next(y, m int) (int, int) {
	if m == 12 {
		return y + 1, 1
	}

	return y, m + 1
}

// まともにやる
//
// 0～6をそれぞれ日曜、月曜、...、土曜に割り当てる。
// 1900/01/01が月曜らしいので1を初期値とし、月の日数や閏年に注意しながら日を
// 進めていき、1901/01/01～2000/12/31の中で1日が日曜になる月をカウントする。
//
// NOTE: ライブラリ使うよりこっちのが速い
func p019A() int {
	// 1900/01/01は月曜(0)
	y := 1900
	m := 1
	w := 1

	// まず1901/01までもっていく
	for y == 1900 {
		w = (w + dates(y, m)) % 7
		y, m = next(y, m)
	}

	// 2001/01/01になるまでループ
	cnt := 0
	for y < 2001 {
		if w == 0 {
			cnt++
		}
		w = (w + dates(y, m)) % 7
		y, m = next(y, m)
	}

	return cnt
}

// 日付ライブラリを使う
//
// NOTE: P019Aより遅いが楽勝
func p019B() int {
	// 1901/01/01～2000/12/31を普通に調べればいい
	cnt := 0
	for d := time.Date(1901, 1, 1, 0, 0, 0, 0, time.Local); d.Year() < 2001; d = d.AddDate(0, 1, 0) {
		if d.Weekday() == time.Sunday {
			cnt++
		}
	}
	return cnt
}
