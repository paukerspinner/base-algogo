// https://leetcode.com/problems/contains-duplicate/
package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}

	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
	fmt.Println(containsDuplicate([]int{-100000, 100000}))
	fmt.Println(containsDuplicate([]int{1}))
	fmt.Println(containsDuplicate([]int{1, 1}))
	fmt.Println(containsDuplicate([]int{0}))
	fmt.Println(containsDuplicate([]int{-1, 0, 1}))
}
