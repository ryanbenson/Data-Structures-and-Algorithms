package rotatearray

import (
	"errors"
)

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
