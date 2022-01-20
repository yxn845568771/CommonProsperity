package _215

func findKthLargest(nums []int, k int) int {
	return a(nums,k)
}

func b(nums []int,k int) int {
	sort.Ints(nums)
	return nums[len(nums) - k ]
}

func a(nums []int,k int) int {
	aa := make([]int,20001)
	for _,v := range nums {
		aa[v+10000] += 1
	}
	for i := 20000;  i >=0 ;i -- {
		if aa[i] == 0 {
			continue
		}
		k -= aa[i]
		if k <= 0 {
			return i - 10000
		}
	}
	return 0
}
