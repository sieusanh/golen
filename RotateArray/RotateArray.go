package RotateArray

import (
	"fmt"
)


// func Test() {
// 	testCases := [][]int{
// 		[]int{1, 2, 3, 4, 5, 6, 7},
// 		[]int{-1, -100, 3, 99},
// 		[]int{3},
// 		[]int{4, 3},
// 		[]int{5, 6},
// 		[]int{1, 3, 5, 4, 2, 3, 4, 5},
// 		[]int{1, 3, 2},
// 	}

// 	for _, testCase := range testCases {
// 		max := rotate(testCase)
// 		fmt.Printf("max = %d", max)
// 		fmt.Println()
// 	}
// }

func Test() {

	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	rotate(arr1, k1)

	arr2 := []int{-1, -100, 3, 99}
	k2 := 2
	rotate(arr2, k2)
}

func rotate(nums []int, k int) {
	len := len(nums)
	// temp := nums[len - k:len]
	// nums = nums[:len - k]
	result := append(nums[len - k:len], nums[:len - k]...)
	// for i := 0; i < k; i++ {
	// 	endEle := nums[len(nums) - 1]
	// 	// fmt.Printf("end %d", endEle)
	// 	fmt.Println()
	// 	nums = nums[:len(nums) - 1]
	// 	// fmt.Println("nums1 ", nums)
	// 	nums = append([]int{endEle}, nums...)
	// 	// fmt.Println("nums2 ", nums)
	// }

	fmt.Println(result)
}

