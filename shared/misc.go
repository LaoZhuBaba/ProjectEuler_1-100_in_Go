package shared

import "math"

func IntToIntSlice(n int) []int {
	slice := make([]int, 0)

	for n != 0 {
		slice = append([]int{n % 10}, slice...)
		n = n / 10
	}
	return slice
}

func Uint64ToUint64Slice(n uint64) []uint64 {
	slice := make([]uint64, 0)

	for n != 0 {
		slice = append([]uint64{n % 10}, slice...)
		n = n / 10
	}
	return slice
}

func IntSliceToInt(s []int) int {
	var ret int
	for _, v := range s {
		ret = ret*10 + v
	}
	return ret
}
func ContainsDuplicates(s []int) bool {
	for index := 0; index < len(s)-1; index++ {
		n1 := s[index]
		//for index, n1 in range s {
		for _, n2 := range s[index+1:] {
			if n2 == n1 {
				return true
			}
		}
	}
	return false
}

func IntContainsDuplicateDigits(n int) bool {
	return ContainsDuplicates(IntToIntSlice(n))
}

func IntContainsDigit0(n int) bool {
	nSlice := IntToIntSlice(n)
	for _, digit := range nSlice {
		if digit == 0 {
			return true
		}
	}
	return false
}

func GetDigitNumberRange(n int) (int, int) {
	ret := 1
	for count := 1; count < n; count++ {
		ret = ret * 10
	}
	return ret, (ret*10 - 1)
}

func Factorise(n int, list *[]int) {
	sr := int(math.Sqrt(float64(n)))
	*list = append(*list, 1)
	for c := 2; c <= sr; c++ {
		if n%c == 0 {
			if c != n/c {
				*list = append(*list, c, n/c)
			} else {
				*list = append(*list, c)
			}
		}
	}
}

func SumOfList(l *[]int) int {
	sum := 0
	for _, n := range *l {
		sum += n
	}
	return sum
}
