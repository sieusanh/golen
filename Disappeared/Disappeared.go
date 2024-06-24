package RotateArray

import (
	"fmt"
)

func Test() {
	arr1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	rotate(arr1)

	arr2 := []int{1, 1}
	rotate(arr2)
}

func findDisappearedNumbers(nums []int) []int {
	mapp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		mapp[nums[i]] = true
	} 	
	
	
}
