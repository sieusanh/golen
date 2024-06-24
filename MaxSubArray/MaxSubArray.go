package MaxSubArray

import (
	"fmt"
)


func Test() {
	testCases := [][]int{
		[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		[]int{1},
		[]int{5, 4, -1, 7, 8},
	}

	for _, testCase := range testCases {
		result := maxSubArray(testCase)
		fmt.Printf("result = %d", result)
		fmt.Println()
	}
}
