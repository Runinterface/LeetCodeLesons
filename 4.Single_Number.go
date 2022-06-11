package main

import "fmt"

// XOR ^
// How this work?
// 0 1 = 1, 0 0 = 0, 1 1 = 0
// ^ - bitwise XOR integers

func singleNumber(nums []int) int {
	var res int
	for _, i := range nums {
		res ^= i
	}
	return res
}
func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}
