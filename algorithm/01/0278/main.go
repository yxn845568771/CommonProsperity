package main

import (
	"fmt"
)

func main() {
	var a struct {
		A []int
	}
	a.A = nil
	fmt.Println(len(a.A))
}

func findMin(nums []int) int {
	l,r := 0,len(nums) -1
	if l < r {
		m := l + (r - l) / 2
		if nums[r] > nums[m]  {
			r = m
		}else {
			l = m + 1
		}
	}
	return nums[l]
}