package main

import "fmt"

func main() {
	a := "123456"
	for i, i2 := range a {
	
	}
	fmt.Println(a[:0])
	
	// fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for {
		if left >= right {
			if nums[left] != target {
				return -1
			}
			return left
		}
		i := (right-left)/2 + left
		if nums[i] == target {
			return i
		}
		if nums[i] < target {
			left = i + 1
			continue
		}
		if nums[i] > target {
			right = i
			continue
		}
	}
}
