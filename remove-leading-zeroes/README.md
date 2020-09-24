# Remove Leading Zeroes

Reference: [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

> Given an int array, remove all leading zeros from the array.

Example:

```console
removeLeading({0, 0, 0, 1, 0, 2, 3})
> 1 0 2 3 
```

## Thoughts

I have a few ideas on implementing this, as it's a simple problem, but you can get creative in how you solve the problem. Some may be more readable than others, but it's interesting to see how they shake out balancing performance and legibility. My first thought was to find the first occurance of an int >= 1, then getting a slice of n to the end of the list. I opted to use the built-in `sort.Search` first, and maybe I'll try a simple for loop after. But the initial results are impressive.

Interestingly enough, the manual for loop is substantially faster, and it's arguably more readable, even if it's more lines of code.

```console
BenchmarkRemoveSearchList-12    	91502872	        13.0 ns/op
BenchmarkRemoveForLoop-12       	289410956	         4.13 ns/op
```