package twosum

func twoSum(list []int, target int) []int {
	// hold our ints in a map so we can do a lookup to see if we have the number we need
	numberMap := make(map[int]int)

	// go through our list and try to find the number we need
	for i, num := range list {
		// what number do we need? look that up
		numberNeeded := target - num

		// do we have the number mapped already?
		numberNeededIndex, ok := numberMap[numberNeeded]
		// if so, we have our answer, so we can short circuit and return now
		if ok {
			return []int{numberNeededIndex, i}
		}

		// didn't get our number, add what number we had to our map
		// maybe it can be used in a subsequent number
		numberMap[num] = i
	}

	// didn't find anything?
	return nil
}
