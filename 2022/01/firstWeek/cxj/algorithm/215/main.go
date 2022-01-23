package main

import (
	"fmt"
	"sort"
)

/*
title: 数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
eg1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
eg2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
tip:
1 <= k <= nums.length <= 104
-104 <= nums[i] <= 104
*/

func main() {
	eg1 := []int{3,2,1,5,6,4}
	eg1k := 2
	fmt.Println(indexKMax1(eg1k, eg1)) // output : 5
	eg2 := []int{3,2,3,1,2,4,5,5,6}
	eg2k := 4
	fmt.Println(indexKMax1(eg2k, eg2)) // output : 4
}

// go package sort IntSlice
func indexKMax1(k int, sli []int) int {
	sort.Ints(sli)
	return sli[len(sli)-k]
}