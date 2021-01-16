func findMaxConsecutiveOnes(nums []int) int {
	Count := 0
	newCount := 0
	for i, v := range nums {
		if v == 1 {
			newCount++
			if i == len(nums)-1 {
				if newCount > Count {
					return newCount
				}
			}
		}
		if v != 1 {
			if newCount > Count {
				Count = newCount

			}
			newCount = 0
		}
	}
	return Count
}