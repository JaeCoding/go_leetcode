package main

import "fmt"

//Given an unsorted integer array, find the smallest missing positive integer.
//
// Example 1:
//
//
//Input: [1,2,0]
//Output: 3
//
//
// Example 2:
//
//
//Input: [3,4,-1,1]
//Output: 2
//
//
// Example 3:
//
//
//Input: [7,8,9,11,12]
//Output: 1
//
//
// Note:
//
// Your algorithm should run in O(n) time and uses constant extra space.
// Related Topics Array

//leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	// find the algo in O(n) time and constant space
	// key: smallest missing positive integer x, must in [1, len + 1]
	//          - when 1 - n is entire, x is len + 1
	//          - when 1 - n is missing one, x is in [1,n]
	// think： hash the n[i] to index of n[i] - 1, ig: n[i] = 4, swap with n[3],
	//        target is swap all num to index of num - 1

	for i := 0; i < len(nums); i++ {

		// 在交换的[1-n]范围内， 且 不等于 nums[i] - 1 索引下的数
		for 1 <= nums[i] && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// find first ele that nums[i] != i + 1
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	s := firstMissingPositive([]int{7, 8, 9, 11, 12})
	fmt.Print(s)
}

//leetcode submit region end(Prohibit modification and deletion)
