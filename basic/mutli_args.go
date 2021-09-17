package main

import "fmt"

func sumInt(nums ...int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return sum
}

func sumInt2(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func multiArgs() {
	sum := sumInt(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(sum)

	nums := []int{1, 3, 5, 7, 9, 11, 13, 15}
	sum2 := sumInt(nums...)
	fmt.Println(sum2)

	sum3 := sumInt2(nums)
	fmt.Println(sum3)
}
