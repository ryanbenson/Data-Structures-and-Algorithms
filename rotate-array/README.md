# Rotate Array

Reference: Issue #8 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

**Implement a function rotateArray(int[] arr, n) which rotates an array by n places.**

```console
> rotateArray([1, 2, 3, 4, 5], 2)
> [4, 5, 1, 2, 3]
```

## Thoughts

This was a fun one. There's a really simple and straight-forward way to approach the problem. There's a built-in way to split the array up. But I want to try other methods too, like rebuilding the array using index calculations instead of splitting into two arrays and then merging.

At least the benchmarks are pretty impressive.

```console
BenchmarkRotateArray-12    	24443248	        48.6 ns/op
```