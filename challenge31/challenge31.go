package challenge31

import "fmt"

func fitCoins(n int64, coins []int64, total *int64) {
	for count, coin := range coins {
		if coin == n {
			(*total)++
			continue
		}
		if coin > n {
			continue
		}
		fitCoins(n-coin, coins[count:], total)
	}
}

func Challenge31() {

	var total int64
	coins := []int64{200, 100, 50, 20, 10, 5, 2, 1}

	fitCoins(200, coins, &total)
	fmt.Printf("Challenge 31 solution is: %d\n", total)
}
