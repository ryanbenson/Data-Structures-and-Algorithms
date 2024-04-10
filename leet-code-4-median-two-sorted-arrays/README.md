# Median Two Sorted Arrays

[Reference](https://leetcode.com/problems/median-of-two-sorted-arrays/)

Given two sorted arrays, return the median of both arrays

- Median = the middle number. If there's an even amount of numbers, then it's the average of the two

## Example

- `Given nums = [2, 7, 11, 15], [1, 2, 5, 7, 9]`
- `Returns: 7`

## Result

Runtime: Beats 93.49% of users with Go

## Thoughts

A simple solution that is easy to read and understand is to simply merge the two arrays, re-sort the arrays and then get the median from the final array

While the simple solution is pretty easy to understand, it comes at the cost of merging the two arrays, then re-sorting, then iterating on the new array. That isn't quite necessary. The cost effective way to figure this out is to:

- Find the median index number using both arrays
- Looping through each sorted array finding the smallest number of the two, using a starting index
- Once we find a number, but we're not at the median, move the starting index of the array we used
- Keep going until we find the median, keeping track of the "previous" number so we can use that if it's an even length of numbers
- Return the median or the (median + previous number) / 2 if it's an even number of elements

The runtime and cost is dramatically less:

```console
BenchmarkGetMedianTwoArrays-12              4663264       246 ns/op
BenchmarkGetMedianTwoArraysOptimized-12    86629614        12.8 ns/op
```
