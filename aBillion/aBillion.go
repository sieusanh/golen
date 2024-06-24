
var aBillion uint64 = 1000000000
	fmt.Println()
	var count uint64 = 1
	for i := 0; ; i++ {
		count *= 2
		if (count >= aBillion) {
			fmt.Println("Bingo !")
			fmt.Println(i)
			break
		}
	}

	