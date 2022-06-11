package main

import "fmt"

func missingNumber(nums []int) int {
	sum := 0
	n := len(nums)
	for _, i := range nums {
		sum += i
	}
	return n*(n+1)/2 - sum
}
func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}
