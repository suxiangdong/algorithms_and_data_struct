package slice

func append1() [2]int {
	slice := make([]int, 3)

	slice1 := append(slice, 1)

	//fmt.Println(len(slice1), cap(slice1))
	// [4,6]
	return [2]int{len(slice1), cap(slice1)}
}

func append2() [2]int {
	slice := make([]int, 3)

	slice1 := append(slice, 1, 1, 1, 1)

	//fmt.Println(len(slice1), cap(slice1))
	// [7,8]
	return [2]int{len(slice1), cap(slice1)}
}
