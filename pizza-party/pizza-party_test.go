package pizzaparty

import "testing"

func TestGetPizzaCount(t *testing.T) {
	people := []person{}
	people = append(people, person{
		name: "Joe",
		num:  9,
	})
	people = append(people, person{
		name: "Cami",
		num:  3,
	})
	people = append(people, person{
		name: "Cassidy",
		num:  4,
	})

	slicesPerPizza := 8
	result := getPizzaCount(people, slicesPerPizza)
	expected := 2

	if result != expected {
		t.Errorf("When getting count of pizzas to get, got %v, want, %v", result, expected)
	}
}

func BenchmarkGetPizzaCount(b *testing.B) {
	people := []person{}
	people = append(people, person{
		name: "Joe",
		num:  9,
	})
	people = append(people, person{
		name: "Cami",
		num:  3,
	})
	people = append(people, person{
		name: "Cassidy",
		num:  4,
	})
	slicesPerPizza := 8
	for i := 0; i < b.N; i++ {
		getPizzaCount(people, slicesPerPizza)
	}
}
