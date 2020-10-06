package pizzaparty

import "math"

type Person struct {
	name string
	num  int
}

func getPizzaCount(people []Person, slices int) int {
	totalSlicesNeeded := 0
	for _, person := range people {
		totalSlicesNeeded += person.num
	}
	totalPizzas := math.Ceil(float64(totalSlicesNeeded / slices))
	return int(totalPizzas)
}
