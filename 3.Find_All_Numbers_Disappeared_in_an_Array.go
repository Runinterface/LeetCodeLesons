package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	seen := make(map[int]bool)
	res := make([]int, 0)
	for _, i := range nums {
		seen[i] = true
	}
	for i := 1; i <= n; i++ {
		ok, _ := seen[i]
		if !ok {
			res = append(res, i)
		}
	}
	return res
}
func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
