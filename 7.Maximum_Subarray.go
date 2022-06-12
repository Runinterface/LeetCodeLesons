package main

import "fmt"

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]
	for _, i := range nums[1:] {
		if sum < 0 {
			sum = i
		} else {
			sum += i
		}
		if max < sum {
			fmt.Println(sum)
			max = sum
			fmt.Println(max)
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSubArray(nums)
}
