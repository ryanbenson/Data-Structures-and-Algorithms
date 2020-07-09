package rotatearray

func rotateArray(arr []int, nthNum int) []int {
	index := nthNum + 1 // use array based, so +1 the nth num
	firstChunk := arr[:index]
	secondChunk := arr[index:]

	var rotatedArr []int
	rotatedArr = append(rotatedArr, secondChunk...)
	rotatedArr = append(rotatedArr, firstChunk...)

	return rotatedArr
}
