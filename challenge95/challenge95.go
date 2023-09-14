package challenge95

import (
	"euler/shared"
	"fmt"
)

const max = 1_000_000

// returnOrderedChain takes a map[int]int and checks that every
// map key is equal to exactly one map value and vice versa.  If
// not, return nil.  Where true, return as a list of numbers
// starting with the lowest value key and each member in the list
// is followed by its mapped value.  For example:
// {2:33, 33:4, 4:2} returns [2, 33, 4]
func returnOrderedChain(m map[int]int) (orderedChain []int) {
	lowest := max + 1 // initialise to an impossibly high number
	reverse := make(map[int]int)
	for k, v := range m {
		if v < lowest {
			lowest = v
		}
		// return nil if a key doesn't exist which is equal to every value
		if _, ok := m[v]; !ok {
			return nil
		}
		// Construct a reverse map so that we can confirm that every value
		// is unique.  To ensure that {2:3, 33:4, 4:33} returns nil.  We
		// test this by confirming that the reverse map is the same length as m
		reverse[v] = k
	}
	if len(reverse) != len(m) {
		return nil
	}
	orderedChain = []int{lowest}
	for v := lowest; m[v] != lowest; v = m[v] {
		orderedChain = append(orderedChain, m[v])
	}
	return orderedChain
}

func Challenge95() {
	var factors []int
	var maxChainLen, maxChainStarter int

	factorSum := make(map[int]int)
	finalList := make(map[string][]int)

	// Construct a map of numbers against the sum of their factors
	// ignore numbers which can't form chains: prime and perfect.
	for count := 0; count < max; count++ {
		factors = []int{}
		shared.Factorise(count, &factors)
		if len(factors) <= 1 {
			// ignore prime numbers
			continue
		}
		thisSum := shared.SumOfList(&factors)
		if thisSum == count {
			// ignore perfect numbers
			continue
		}
		factorSum[count] = thisSum
	}

	for k, v := range factorSum {
		n := v

		// Initialise the chain
		chain := map[int]int{k: v}
		for {
			var next int
			var ok bool
			// If factorSum[n] exists in map then it becomes the next
			// n.  Otherwise the chain ends, so break
			if next, ok = factorSum[n]; !ok {
				break
			}
			// If a chain doesn't already exists then create one
			if _, ok := chain[n]; !ok {
				chain[n] = factorSum[n]
			} else {
				orderedChain := returnOrderedChain(chain)
				if orderedChain != nil {
					// use orderedChainStr as a key to prevent duplicates
					orderedChainStr := fmt.Sprintf("%v", orderedChain)
					finalList[orderedChainStr] = orderedChain
				}
				break
			}
			n = next
		}
	}
	for _, l := range finalList {
		if len(l) > maxChainLen {
			maxChainLen = len(l)
			maxChainStarter = l[0]
		}
	}
	fmt.Printf("challenge 95 solution is: %d with a chain length of %d\n", maxChainStarter, maxChainLen)
}
