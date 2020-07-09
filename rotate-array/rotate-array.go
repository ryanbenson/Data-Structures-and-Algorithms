package rotatearray

func rotateArray(arr []int, index int) []int {
	firstChunk := arr[:index]
	secondChunk := arr[index:]

	var rotatedArr []int
	rotatedArr = append(rotatedArr, firstChunk...)
	rotatedArr = append(rotatedArr, secondChunk...)

	return rotatedArr
}
