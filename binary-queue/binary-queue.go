package binaryqueue

import "strconv"

// binaryQueue takes an number, then creates a queue
// of the numbers leading from 1 to n in the queue
// and the values in the queue are the binary versions of the nums
func binaryQueue(lastNum int) []string {
	queue := []string{}
	if lastNum <= 0 {
		return queue
	}

	for i := 1; i <= lastNum; i++ {
		// need to use int64 to get binary representation of int
		binaryOfInt := strconv.FormatInt(int64(i), 2)
		queue = append(queue, binaryOfInt)
	}
	return queue
}
