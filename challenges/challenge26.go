// A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

// 1/2	= 	0.5
// 1/3	= 	0.(3)
// 1/4	= 	0.25
// 1/5	= 	0.2
// 1/6	= 	0.1(6)
// 1/7	= 	0.(142857)
// 1/8	= 	0.125
// 1/9	= 	0.(1)
// 1/10	= 	0.1
// Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

// Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.

package challenges

import "fmt"

var maxCount int = 0
var maxCountNum int = 0

func div(n, d, precision int) string {
	//var dividend, remainder int
	s := "0."
	flag := false
	count := 0
	m := make(map[int]int)
	for precision > 0 {
		if m[n] == 1 && !flag {
			s += "."
			flag = true
		}
		if m[n] == 2 {
			s += fmt.Sprintf(" %d", count)
			break
		}
		if flag {
			count++
		}
		m[n]++
		s += fmt.Sprintf("%d", n/d)
		n = 10 * (n % d)
		if n == 0 {
			return s
		}
		precision--
	}
	if count > maxCount {
		maxCount = count
		maxCountNum = d
	}
	return s
	// if precision == 0 {
	// 	return ""
	// }
	// dividend := n / d
	// remainder := n % d
	// s := fmt.Sprintf("%d", dividend)
	// if remainder == 0 {
	// 	return s
	// } else {
	// 	return s + div(remainder*10, d, precision-1)
	// }
}

func Challenge26() {
	for n := 2; n < 900; n++ {
		fmt.Printf("%s\n", div(10, n, 2200))
	}
	fmt.Printf("Answer is %d with a repeat count of %d\n", maxCountNum, maxCount)
}
