package main

import (
	"fmt"
	"sort"
)

func calculateBuses(n int, families []int) int {
	if len(families) != n {
		fmt.Println("The number of families does not match the input.")
		return -1
	}

	sort.Ints(families)
	fmt.Println(families)

	buses := 0
	left, right := 0, n-1
	for left <= right {
		if families[left]+families[right] <= 4 {
			left++
			right--
		} else {
			right--
		}
		buses++
	}

	return buses
}

func main() {
	var n int
	fmt.Print("Input the number of families : ")
	fmt.Scan(&n)

	families := make([]int, n)
	fmt.Println("Input the number of members in the family ( separated by a space) : ")
	for i := 0; i < n; i++ {
		fmt.Scan(&families[i])
	}

	result := calculateBuses(n, families)
	if result != -1 {
		fmt.Println("Minimum bus required is :", result)
	}
}
