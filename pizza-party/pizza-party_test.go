package pizzaparty

import "testing"

func TestGetPizzaCount(t *testing.T) {
	people := []Person{}
	slicesPerPizza := 8
	result := getPizzaCount(people, slicesPerPizza)
	expected := 3

	if result != expected {
		t.Errorf("When getting count of pizzas to get, got %v, want, %v", result, expected)
	}
}
