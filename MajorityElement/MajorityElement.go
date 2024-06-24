package MajorityElement

import ("fmt")

func Test() {
	testCases := [][]int{
		[]int{3, 2, 3},
		[]int{2, 2, 1, 1, 1, 2, 2},
	}

	for _, testCase := range testCases {
		result := majorityElement2(testCase)
		fmt.Printf("result = %d", result)
		fmt.Println()
	}
}

func majorityElement(nums []int) int {
    mapp := make(map[int]int)

	for i := 0 ; i < len(nums); i++ {
		ele := nums[i]
		mapp[ele]++
		if mapp[ele] > len(nums) / 2 {
			return ele
		}
	}

	return nums[0]
}

func majorityElement2(nums []int) int {
	for i := 0 ; i < len(nums) - 1; i++ {
		count := 0
		ele := nums[i]
		for j := 0; j < len(nums); j++ {
			if nums[j] == ele {
				count++
			}
			if count > len(nums) / 2 {
				return ele
			}
		}
	}

	return nums[0]
}


