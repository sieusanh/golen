package HouseRobber

import (
	"fmt"
)

func Test() {
	arr1 := []int{1, 2, 3, 1}
	r1 := rob(arr1)
	fmt.Println(r1)

	arr2 := []int{2, 7, 9, 3, 1}
	r2 := rob(arr2)
	fmt.Println(r2)

	arr3 := []int{2, 1, 1, 2}
	r3 := rob(arr3)
	fmt.Println(r3)
}


func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

    oddSum := 0
	evenSum := 0
	for i := 0; i < len(nums); {
		evenSum += nums[i]
		if i + 1 >= len(nums) {
			break
		}
		oddSum += nums[i+1]
		i += 2
	}

	if oddSum > evenSum {
		return oddSum
	} else {
		return evenSum
	}
}