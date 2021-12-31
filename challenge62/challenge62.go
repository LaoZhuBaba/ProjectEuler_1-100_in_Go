package challenge62

import "fmt"

func int64ToOrderedString(i int64) string {
	s := fmt.Sprintf("%d", i)
	var a10 [10]int
	for _, v := range s {
		a10[v-'0']++
	}
	var retStr string
	for k, v := range a10 {
		for count := 0; count < v; count++ {
			retStr += fmt.Sprintf("%d", k)
		}
	}
	return retStr
}
func Challenge62() {
	m := make(map[string][]int64)
	for count := int64(1); count < 10000; count++ {
		cubed := count * count * count
		orderedStr := int64ToOrderedString(cubed)
		m[orderedStr] = append(m[orderedStr], cubed)
		if len(m[orderedStr]) > 4 {
			fmt.Printf("Found this set of cubes which are string permutations of one another: %v\n", m[orderedStr])
			var solution int64
			for _, v := range m[orderedStr] {
				if solution == 0 || v < solution {
					solution = v
				}
			}
			fmt.Printf("Challenge 62 solution is: %d\n", solution)
			return
		}
	}
}
