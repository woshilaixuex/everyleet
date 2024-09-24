package leetcodego

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150 80. 删除有序数组中的重复项 II
 * 输入：nums = [1,1,1,2,2,3]
 * 输出：5, nums = [1,1,2,2,3] O(1)空间复杂度
 * @Date: 2024-09-24 20:37
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lennums := len(nums)
	temp := 0
	number := nums[0]
	index := 0
	reverArr := func(temp, index_ int) {
		for j := index + 2; index_ < lennums; {
			nums[j] = nums[index_]
			j++
			index_++
		}
	}
	for i := 0; i < lennums; i++ {
		if number != nums[i] {
			if temp > 2 {
				reverArr(temp, i)
				lennums = lennums - (temp - 2)
				i = index + 2
			}
			temp = 1
			index = i
			number = nums[i]
		} else {
			temp++
		}
	}
	if temp > 2 {
		reverArr(temp, lennums-1)
		lennums = lennums - (temp - 2)
	}
	return lennums
}
func removeDuplicates_(nums []int) int { // 官方题解
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
