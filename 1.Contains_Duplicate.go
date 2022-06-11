package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := map[int]bool{}
	for _, r := range nums {
		if found := seen[r]; found {
			return true
		}
		seen[r] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
