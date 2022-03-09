# Longest Subsequence

Reference: Issue #236 of [cassidoo newsletter](https://buttondown.email/cassidoo/archive/no-act-of-kindness-no-matter-how-small-is-ever/) 🎉 (amazing newsletter by the way)

**Given an array of integers, find the length of the longest sub-sequence such that elements in the sub-sequence are consecutive integers, the consecutive numbers can be in any order.**

```console
$ longestSubSeq([1, 9, 87, 3, 10, 4, 20, 2, 45])
$ 4 // 1, 3, 4, 2

$ longestSubSeq([36, 41, 56, 35, 91, 33, 34, 92, 43, 37, 42])
$ 5 // 36, 35, 33, 34, 37
```

## Thoughts

This was an interesting one. A quick way to make this easier was to sort the main sequence. That makes it where we can check the sibling value for the next possible sub sequence. Otherwise we would need to do something extra like re-looping through the values to find a sibling, which would be wasteful for time.

I also realized that I could shorten the processing time by short circuiting if we have a sequence that's longer than what's leftover in the array. Example, if the current longest subsequence stopped at 8, and we now need to find a new one, but there's only 3 items left in the array, then there's no reason to keep trying because it won't ever find a new longest subsequence, so might as well short circuit.

Overall the solution runs very fast:

```console
BenchmarkLargestRect-12          4700745               277.6 ns/op
```
