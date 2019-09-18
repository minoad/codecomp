package main

import "fmt"

// func twoSum(nums []int, target int) []int {
// 	n := []int{0, 5}

// 	for i, j := range nums {
// 		for m, n := range nums {
// 			sum := j + n
// 			if sum == target && i != m {
// 				fmt.Printf("%d-%d: %d + %d = %d\n", i, m, j, n, sum)
// 				return []int{i, m}
// 			}

// 		}
// 	}

// 	return n
// }

func twoSum(nums []int, target int) []int {
	for x, y := range nums {
		for i := 0; i < len(nums); i++ {
			if i != x {
				sums := y + nums[i]
				if sums == target {
					fmt.Println([]int{x, i})
					return []int{x, i}
					break
					//fmt.Printf("%d + %d = %d\n", y, nums[i], sums)
				}

			}

		}
	}

	return []int{0, 0}
}

func twoSumhash(nums []int, target int) []int {

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


func main() {
	fmt.Println("a ...interface{}")
	twoSum([]int{2, 7, 11, 15}, 9)
}
