package twosum

// https://leetcode.com/problems/two-sum/

/* Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order. */

// HashMap Solution O(n)
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // map with int value and key
	results := []int{}
	for i, num := range nums {
		compliment := target - num // if target is 10 and curr is 4 we need to find 6 in map
		if j, found := numMap[compliment]; found {
			results = []int{j, i} //
		}
		numMap[num] = i // set key and value
	}
	return results
}

// Brute Force Solution O(n^2)
/* func TwoSum(nums []int, target int) []int {
	results := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				results = []int{i, j}
			}
		}
	}
	return results
} */
