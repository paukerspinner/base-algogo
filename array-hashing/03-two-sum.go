// https://leetcode.com/problems/two-sum/description/
package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	pos := map[int][]int{}

	for idx, num := range nums {
		pos[num] = append(pos[num], idx)
	}

	sort.Ints(nums)

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			break
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	firstIdx := pos[nums[left]][0]
	pos[nums[left]] = pos[nums[left]][1:]
	secondIdx := pos[nums[right]][0]

	return []int{firstIdx, secondIdx}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{3, 0}, 3))
	fmt.Println(twoSum([]int{1, 2, 1, 2}, 3))
	fmt.Println(twoSum([]int{1, 2, 1, 2, 1, 2}, 2))
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(twoSum([]int{5, 4, 3, 2, 1}, 4))
}
