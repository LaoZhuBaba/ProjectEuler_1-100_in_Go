// A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example,
// the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

// A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this
// sum exceeds n.

// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum
// of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be
// written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even
// though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than
// this limit.

// Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

package challenges

import "fmt"

const p23_max = 28123

// const p23_max = 3000

func checkIfSumOf2InList(list *[]int, sum int) (bool, int, int) {
	for n1 := 0; (*list)[n1] <= sum; n1++ {
		for n2 := 0; (*list)[n1]+(*list)[n2] <= sum; n2++ {
			if (*list)[n1]+(*list)[n2] == sum {
				return true, (*list)[n1], (*list)[n2]
			}
		}
	}
	return false, 0, 0
}
func Challenge23() {

	var factorSum int
	abundantList := new([]int)
	var p23Total int

	for c := 1; c <= p23_max; c++ {
		list := new([]int)
		factorise(c, list)
		factorSum = sumOfList(list)
		if factorSum == c {
			fmt.Printf("%d is a perfect number!\n", c)
		} else {
			if factorSum > c {
				fmt.Printf("%d is an abundant number!\n", c)
				*abundantList = append(*abundantList, c)
			}
		}
		fmt.Printf("For c = %d, factor list is: %v\n", c, *list)
		*list = nil
	}
	for c := 1; c <= p23_max; c++ {
		tf, _, _ := checkIfSumOf2InList(abundantList, c)
		if tf {
			continue
			//			fmt.Printf("%d is the sum of %d & %d\n", c, n1, n2)
		}
		if !tf {
			fmt.Printf("%d is not the sum of two abundant numbers!\n", c)
			p23Total += c
		}
	}
	fmt.Printf("Problem 23 total is: %d\n", p23Total)
}
