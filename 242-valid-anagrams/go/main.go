package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {

	if isLengthDiff := len(s) != len(t); isLengthDiff {
		return false
	}

	charInventory := make(map[rune]int, len(s))

	for _, currentChar := range s {
		charInventory[currentChar]++
	}

	for _, currentChar := range t {
		if _, exist := charInventory[currentChar]; exist {
			if isZero := charInventory[currentChar] == 0; isZero {
				return false
			}
			charInventory[currentChar] = charInventory[currentChar] - 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	result := isAnagram("zorra", "arroz")
	fmt.Println(result)
}
