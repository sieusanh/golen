package LCIS

import (
	"fmt"
)



func Test() {
	testCases := [][]int{
		[]int{1, 3, 5, 4, 7},
		[]int{2, 2, 2},
		[]int{3},
		[]int{4, 3},
		[]int{5, 6},
		[]int{1, 3, 5, 4, 2, 3, 4, 5},
		[]int{1, 3, 2},
	}

	for _, testCase := range testCases {
		max := findLengthOfLCIS(testCase)
		fmt.Printf("max = %d", max)
		fmt.Println()
	}
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	l := 0
	r := 0
	k := nums[1] - nums[0]
	max := 1

	for i := 0; i < len(nums) - 1; i++ {
		r = i + 1
		d := nums[i + 1] - nums[i]

		if d > 0 {
			if r - l + 1 > max {
				max = r - l + 1
			}
			if d != k {
				l = i + 1
			}
		} else {
			if d < 0 {
				l = i + 1
			}
			k = d
		}
	}

	return max
}

/*
dry run



[1, 3, 5, 4, 7]

l = 0, r = 0, k = 2, max = 1

i = 0, r = 1, d = 2, max = 2
i = 1, r = 2, d = 2, max = 3
i = 2, r = 3, d = -1
=> l = 3, k = -1

i = 3, r = 4, d = 3
=> l = 4, k = 3

[1, 3, 5, 4, 2, 3, 4, 5]
i = 2, r = 3, d = 1
=> l = 3, k = 1

i = 3, r = 4, d = -2
=> l = 4, k = -2

i = 4, r = 5, d = 1
=> l = 5, k = 1

i = 5, r = 6, d = 1
max = 3

i = 6, r = 7, d = 1
max = 7 - 5 + 1

*/

