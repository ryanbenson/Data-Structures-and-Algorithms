package interioranglesize

// provides the interior angle of a same-size polygon with the given size
func getInteriorAngle(size int) int {
	if size < 2 {
		return 0
	}
	return (size-2) * 180 / size
}
