package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculateBuses(n int, families []int) int {
	sort.Ints(families)

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

	fmt.Println("Input the number of members in the family (separated by a space) : ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	familyStrings := strings.Fields(input)

	if len(familyStrings) != n {
		fmt.Println("Input must be equal with count of family")
		return
	}

	families := make([]int, n)
	for i := 0; i < n; i++ {
		num, err := strconv.Atoi(familyStrings[i])
		if err != nil {
			fmt.Println("Invalid number:", familyStrings[i])
			return
		}
		families[i] = num
	}

	result := calculateBuses(n, families)
	fmt.Println("Minimum bus required is :", result)
}
