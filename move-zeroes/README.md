# Move Zeroes

Reference: Issue #157 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

> Given an array of random integers, move all the zeros in the array to the end of the array

Example:

```console
moveZeros([1, 2, 0, 1, 0, 0, 3, 6]) => [1, 2, 1, 3, 6, 0, 0, 0]
```

## Bonus

Keep this in O(n) time (or better)!

## Thoughts

My initial pass on this seemed pretty simple and hits the bonus. It creates an array with the exact length of the one given, with all 0s as the default values. It then goes through each item in the array and if it's `> 0` then it updates the next item in the array, starting from index 0, and using sort of a sliding window index. That way all of the numbers get added, in the right order, and the remainder 0 values are already there. It benchmarks pretty well too.

Had a clever idea to make an array of just the values (numbers greater than zero), then make a second array that is the length of the difference between the original array and the numbers, and merge them together. Granted, I had a feeling that since it was working with arrays and merging, it'd be notably slower, and it was. But it was still fun try to try.

```console
BenchmarkMoveZeroes-12         	43027975	        27.5 ns/op
BenchmarkMoveZeroesMerge-12    	 9310798	       130 ns/op
```