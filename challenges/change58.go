package challenges

import "fmt"

func Challenge58() {

	list := []int{1}
	var i int
	primeMap := make(map[int]bool)
	for count := 3; count < 1_000_000; count += 2 {
		cxc := count * count
		i = cxc - 3*(count-1)
		list = append(list, i)
		primeMap[i] = isPrime(i)
		i = cxc - 2*(count-1)
		list = append(list, i)
		primeMap[i] = isPrime(i)
		i = cxc - 1*(count-1)
		list = append(list, i)
		primeMap[i] = isPrime(i)
		i = cxc
		list = append(list, i)
		primeMap[i] = isPrime(i)
		var countPrime int
		for _, n := range list {
			if primeMap[n] {
				// fmt.Printf("%d is prime\n", n)
				countPrime++
			}

		}
		// fmt.Printf("countPrime is: %d\n", countPrime)
		// fmt.Printf("len(list) is: %d\n", len(list))
		pc := 100 * countPrime / len(list)
		//fmt.Printf("For count %d, %d%% are prime\n", count, pc)
		if pc < 10 {
			fmt.Printf("solution: %d\n", count)
			break
		}

	}
	//fmt.Printf("%v\n", list)
}
