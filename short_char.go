package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func processString(input string) (string, string) {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	order := make(map[rune]int)
	var vowelChar, consonantChar []rune

	for i, char := range strings.ToLower(input) {
		if !unicode.IsLetter(char) {
			continue
		}

		if _, exists := order[char]; !exists {
			order[char] = i
		}

		if vowels[char] {
			vowelChar = append(vowelChar, char)
		} else {
			consonantChar = append(consonantChar, char)
		}
	}

	// fmt.Println(order)

	sort.SliceStable(vowelChar, func(i, j int) bool {
		return order[vowelChar[i]] < order[vowelChar[j]]
	})

	sort.SliceStable(consonantChar, func(i, j int) bool {
		return order[consonantChar[i]] < order[consonantChar[j]]
	})

	return string(vowelChar), string(consonantChar)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input one line of words (S): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Error invalid input")
		return
	}
	vowels, consonants := processString(input)
	fmt.Println("Vowel Characters:", vowels)
	fmt.Println("Consonant Characters:", consonants)
}
