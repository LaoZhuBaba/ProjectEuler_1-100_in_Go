// There exists exactly one Pythagorean triplet for which a + b + c = 1000
package challenge9

import "fmt"

var max = 1000

var n3 int

func Challenge9() {
out:
	for n1 := 1; n1 < max/3; n1 += 1 {
		for n2 := n1 + 1; n2 < max/2; n2 += 1 {
			if n2*2+n1 >= max {
				break
			}
			n3 = max - n2 - n1
			//fmt.Printf("%d %d %d\n", n1, n2, n3)
			if n1*n1+n2*n2 == n3*n3 {
				fmt.Printf("*** %d %d %d ***\n", n1, n2, n3)
				fmt.Printf("Challenge 9 solution is: %d\n", n1*n2*n3)
				break out
			}
		}
	}
}

// func Challenge9() {
// 	out:
// 		for n1 := 1; n1 < 500; n1 += 1 {
// 			for n2 := 2; n2 < 500; n2 += 1 {
// 				for n3 := 3; n3 < 500; n3 += 1 {
// 					if n1*n1+n2*n2 == n3*n3 && n1+n2+n3 == 1000 {
// 						fmt.Printf("%d %d %d\n", n1, n2, n3)
// 						break out
// 					}
// 				}
// 			}
// 		}
// 	}
