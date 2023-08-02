package main

import (
	"reflect"
	"testing"
)

func TestTwoSumBrute(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
		{[]int{5, 7, 11, 15}, 22, []int{1, 3}},
		{[]int{1, 3, 7, 9, 13}, 22, []int{3, 4}},
		{[]int{5000000, 2000000, 3000000, 4000000}, 5000000, []int{1, 2}},
	}

	for _, tc := range testCases {
		got := twoSumMap(tc.arr, tc.target)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("twoSumBrute(%v, %v) = %v; want %v", tc.arr, tc.target, got, tc.want)
		}
	}
}
