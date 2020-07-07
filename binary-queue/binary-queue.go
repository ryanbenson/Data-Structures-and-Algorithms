package binaryqueue

import "strconv"

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
