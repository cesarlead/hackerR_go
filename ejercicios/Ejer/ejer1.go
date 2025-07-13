package Ejer

// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
func SumSubArrayK(nums []int, k int) []int {

	if k <= 0 {
		return []int{}
	}

	var result []int
	var currentSum, count int

	for _, num := range nums {
		currentSum += num
		count++

		if currentSum == k {
			result = append(result, currentSum)
			currentSum = 0
			count = 0
		}
	}

	if count > 0 {
		result = append(result, currentSum)
	}

	return result

}
