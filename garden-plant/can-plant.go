package canplant

import "math"

func canPlant(garden []int, plants int) bool {
	totalAvailableSpots := 0
	currentAvailableSpotsLen := 0
	for _, spot := range garden {
		if spot == 0 {
			currentAvailableSpotsLen++
		}
		if spot == 1 {
			// easier if the spots are even since each plant needs an open space around it
			// if it's even, take away one
			if currentAvailableSpotsLen %2 == 0 {
				currentAvailableSpotsLen = currentAvailableSpotsLen - 1
			}
			plantableSpotsFound := math.Ceil(float64(currentAvailableSpotsLen) / float64(3))
			totalAvailableSpots = totalAvailableSpots + int(plantableSpotsFound)
			// reset so we can keep going
			currentAvailableSpotsLen = 0
		}
	}
	if plants <= totalAvailableSpots {
		return true
	}

	return false
}