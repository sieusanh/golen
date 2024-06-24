package FactorialTrailingZero

import ("fmt")

func Test() {
	// testCases := []int{3, 5, 0}
	testCases := []int{30}

	for _, testCase := range testCases {
		result := trailingZeroes(testCase)
		fmt.Printf("result = %d", result)
		fmt.Println()
	}
}

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	var num uint32 = 1
    for n != 1 {
		num = num * uint32(n)
		if num < 0 {
			fmt.Println(num)
		}
        n = n - 1
		fmt.Printf("n %d", n)
		fmt.Println()
    }    

    count := 0
	fmt.Println("num ", num)
	for num != 1 {
		if num % 10 != 0 {
			break	
		} 

		num /= 10
		count++
	}

    return count
}

// func trailingZeroes2(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

//     count := 0
// 	l := n
// 	r := 1
//     for l > r {

// 		median := (l - r) / 2

//         if n * (n - 1) % 10 == 0 {
//             count++
//         }

// 		if l * r 

// 		if l * median > 10 {
// 			r = median
// 		} 
// 		if r * median > 10 {
// 			l = median
// 		}
//     }    
//     return count
// }