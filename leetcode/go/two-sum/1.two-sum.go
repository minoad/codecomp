//package twosum
/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

//  func twoSum(nums []int, target int) []int {
// 14%
// 	for i, j := range nums {
// 		for m, n := range nums {
// 			sum := j + n
// 			if sum == target && i != m {
// 				return []int{i, m}
// 			}

// 		}
// 	}

// 	return []int{0,0}
// }

// 39%
// func twoSum(nums []int, target int) []int {
// 	for x, y := range nums {
// 		for i := 0; i < len(nums) && i != x ; i++ {
// 			if y+nums[i] == target {
// 				return []int{x, i}
// 			}
// 		}
// 	}

// 	return []int{0, 0}
// }

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, ok := m[complement]; ok && m[complement] != i {
			return []int{i, m[complement]}
		}

	}
	return nil
}


