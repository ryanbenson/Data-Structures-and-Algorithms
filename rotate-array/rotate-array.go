package rotatearray

import (
	"errors"
)

// rotateArray rotates an array using array slicing
func rotateArray(arr []int, index int) ([]int, error) {
	if (index + 1) == len(arr) {
		return arr, nil
	}

	if index > len(arr) {
		return nil, errors.New("Number provided exceeds length of array")
	}

	firstChunk := arr[:index+1]
	secondChunk := arr[index+1:]

	var rotatedArr []int
	rotatedArr = append(rotatedArr, secondChunk...)
	rotatedArr = append(rotatedArr, firstChunk...)

	return rotatedArr, nil
}

// rotateArrayLoop rotates an array given an index using a loop
func rotateArrayLoop(arr []int, index int) ([]int, error) {
	arrLen := len(arr)
	if (index + 1) == arrLen {
		return arr, nil
	}

	if index > arrLen {
		return nil, errors.New("Number provided exceeds length of array")
	}

	// setup an array with the same length as the one given, just empty values
	rotatedArr := make([]int, arrLen)
	// keep track of where to insert the proper values in the right places
	currentItemIndex := index

	// go through each item in the array and move it to the right place
	for _, val := range arr {
		rotatedArr[currentItemIndex] = val
		currentItemIndex++
		// if we're at the end, hop back to the beginning of the array
		if currentItemIndex >= arrLen {
			currentItemIndex = 0
		}
	}

	return rotatedArr, nil
}
