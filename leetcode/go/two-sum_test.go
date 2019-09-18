package main

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	n := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(n)
}

func Test_twoSumhash(t *testing.T) {
	twoSumhash([]int{3, 2, 4}, 6)
}
