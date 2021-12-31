// You are given the following information, but you may prefer to do some research for yourself.

// 1 Jan 1900 was a Monday.
// Thirty days has September,
// April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.
// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

package challenge19

import "fmt"

type Date = struct {
	dom  int
	dow  int
	moy  int
	year int
}

func lengthOfMonth(mon int, year int) int {

	// Test for Jan, Mar, May, Jul
	if mon <= 7 && mon%2 == 1 {
		return 31
	}
	// Test for Aug, Oct, Dec
	if mon >= 8 && mon%2 == 0 {
		return 31
	}
	if mon == 2 {
		if year%4 == 0 {
			if year%100 == 0 {
				if year%400 == 0 {
					return 29
				}
				return 28
			}
			return 29
		}
		return 28
	}
	return 30
}
func nextDoW(d *Date) {
	d.dow = (d.dow + 1) % 7
}

func nextDoM(d *Date) {
	if d.dom < lengthOfMonth(d.moy, d.year) {
		d.dom++

	} else {
		if d.moy == 12 {
			d.moy = 1
			d.dom = 1
			d.year++
		} else {
			d.dom = 1
			d.moy++
		}
	}
}

func nextDay(day *Date) {
	nextDoW(day)
	nextDoM(day)
}

func Challenge19() {
	daysOfWeek := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	sundayCount := 0

	for day := (Date{1, 2, 1, 1901}); day.year != 2001; nextDay(&day) {
		fmt.Printf("%s %d-%d-%d\n", daysOfWeek[day.dow], day.year, day.moy, day.dom)
		if day.dow == 0 && day.dom == 1 {
			sundayCount++
		}
	}
	fmt.Printf("Challenge 19 solution: total number of first-of-month Sundays is %d\n", sundayCount)
}
