package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3, 2, 4}, 6))
}

func findIndex(arr []int, target int) int {
	for index := range arr {
		if arr[index] == target {
			return index
		}
	}
	return -1
}

func twoSum(arr []int, target int) [2]int {
	var resultIndex int
	for index, value := range arr {
		resultIndex = findIndex(arr, target-value)
		if resultIndex != -1 && resultIndex != index {
			return [2]int{index, resultIndex}
		}
	}
	return [2]int{}
}
