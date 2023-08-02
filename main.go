package main

import (
	"fmt"
	"sort"

	slice "github.com/VincentQDo/goutils"
)

func main() {
	testCases := []map[string][]int{
		{"input": []int{2, 7, 11, 15}, "target": []int{9}},
		{"input": []int{3, 2, 4}, "target": []int{6}},
		{"input": []int{3, 3}, "target": []int{6}},
		{"input": []int{-1, -2, -3, -4, -5}, "target": []int{-8}},
		{"input": []int{0, 4, 3, 0}, "target": []int{0}},
		{"input": []int{5, 7, 11, 15}, "target": []int{22}},
		{"input": []int{1, 3, 7, 9, 13}, "target": []int{22}},
		{"input": []int{1000000, 2000000, 3000000, 4000000}, "target": []int{5000000}},
	}

	for _, value := range testCases {
		fmt.Println(twoSumMap(value["input"], value["target"][0]))
	}
}

func twoSumBrute(arr []int, target int) []int {
	var resultIndex int
	for index, value := range arr {
		resultIndex = slice.FindIndexInt(arr, target-value)
		if resultIndex != -1 && resultIndex != index {
			result := []int{index, resultIndex}
			sort.Ints(result)
			return result
		}
	}
	return []int{}
}

func twoSumMap(arr []int, target int) []int {
	var compliment int
	complimentMap := map[int]int{}
	for i, v := range arr {
		compliment = target - v
		_, ok := complimentMap[compliment]
		if ok {
			return []int{complimentMap[compliment], i}
		} else {
			complimentMap[v] = i
		}
	}
	return []int{}
}
