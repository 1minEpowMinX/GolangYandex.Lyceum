package main

type Chest struct {
	val []int
	wt  []int
}

func Knapsack(chest *Chest, maxWeight int) (int, []int) {
	n := len(chest.val)
	newItem := false
	backpack := make([][]int, n+1)
	item_numbers := make([]int, 0, n+1)
	for i := range backpack {
		backpack[i] = make([]int, maxWeight+1)
	}

	for item := 1; item <= n; item++ {
		newItem = false
		for capacity := 1; capacity <= maxWeight; capacity++ {
			maxcostWithoutCurrent := backpack[item-1][capacity]
			maxcostWithCurrent := 0

			weightOfCurrent := chest.wt[item-1]
			if capacity >= weightOfCurrent {
				maxcostWithCurrent = chest.val[item-1]
				remainingCapacity := capacity - weightOfCurrent
				maxcostWithCurrent += backpack[item-1][remainingCapacity]
			}

			if maxcostWithCurrent > maxcostWithoutCurrent {
				backpack[item][capacity] = maxcostWithCurrent
				newItem = true
			} else {
				backpack[item][capacity] = maxcostWithoutCurrent
			}
		}
		if newItem {
			item_numbers = append(item_numbers, item-1)
			newItem = false
		}
	}

	return backpack[n][maxWeight], item_numbers
}
