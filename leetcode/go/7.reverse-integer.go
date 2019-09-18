import (
	"fmt"
	"regexp"
	"strconv"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
func reverse(x int) int {
	//fmt.Println(x)
	r, _ := regexp.Compile("[0-9]")
	res := ""
	strInt := strconv.FormatInt(int64(x), 10)
	for i := len(strInt) - 1; i >= 0; i-- {
		if r.MatchString(string(strInt[i])) {
			res += string(strInt[i])
		}
	}
	n, err := strconv.Atoi(res)
	if err != nil {
		fmt.Println(err)
	}

	if n > 2147483648 || (n*-1) < -2147483648 {
		return 0
	}
	if x < 0 {
		return n * -1
	}
	return n
}

