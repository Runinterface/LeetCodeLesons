package main

import "fmt"

// You gotta use Fibonacci digit
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	n1 := 1
	n2 := 2
	for i := 3; i < n+1; n-- {
		n1, n2 = n2, n1+n2
	}
	return n2
}
func main() {
	n := 3
	fmt.Println(climbStairs(n))
}
