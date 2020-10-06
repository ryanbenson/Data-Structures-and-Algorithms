# Pizza Party

Reference: Issue #163 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

> Given an array of people objects (where each person has a name and a number of pizza slices they’re hungry for) and a number for the number of slices that the pizza can be sliced into, return the number of pizzas you need to buy.

```console
arr = [{ name: Joe, num: 9 }, { name: Cami, num: 3 }, { name: Cassidy, num: 4 }]
gimmePizza(arr, 8)
2 // 16 slices needed, pizzas can be sliced into 8 pieces, so 2 pizzas should be ordered
```

## Thoughts

Was a fun exercise, and man can Joe really pack it away. I wanted to try to optimize it quickly so the operation can be done effectively, and the trick was a simple loop with some type conversions. Might be able to squeak out more, but could come at a loss of legibility. And it's pretty easy to understand as it is.

```console
BenchmarkGetPizzaCount-12    	91164912	        13.1 ns/op
```