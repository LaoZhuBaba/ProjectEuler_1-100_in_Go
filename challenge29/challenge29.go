package challenge29

import (
	"fmt"
	"math/big"
)

func pow(a, b int64) *big.Int {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	var ret = new(big.Int)
	ret = ret.Exp(bigA, bigB, nil)
	return ret
}

// func pow(a, b int64) *big.Int {

// 	bigA := big.NewInt(a)
// 	exp := big.NewInt(a)

// 	for count := int64(1); count < b; count++ {
// 		exp.Mul(exp, bigA)
// 	}
// 	return exp
// }
func bigIntInList(n *big.Int, l []big.Int) bool {
	for count := 0; count < len(l); count++ {
		if l[count].Cmp(n) == 0 {
			return true
		}
	}
	return false
}
func Challenge29() {
	c29Max := 0
	list := make([]big.Int, 0)
	//m := make(map[string]bool)

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			e := pow(int64(a), int64(b))
			// fmt.Printf("%d to the power of %d is %d\n", a, b, e)
			if !bigIntInList(e, list) {
				list = append(list, *e)
				if len(list) > c29Max {
					c29Max = len(list)
				}
			}
		}
	}
	fmt.Printf("Challenge 29 solution is %d\n", c29Max)
}
