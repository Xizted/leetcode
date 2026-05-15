package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if isEmpty := len(nums) == 0; isEmpty {
		return nil
	}

	numsSaved := make(map[int]int, len(nums))

	for currentIndex, currentValue := range nums {
		value_searched := target - currentValue

		if index, exist := numsSaved[value_searched]; exist {
			return []int{index, currentIndex}
		}
		numsSaved[currentValue] = currentIndex
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
