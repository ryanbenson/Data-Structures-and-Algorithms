# Median Two Sorted Arrays
[Reference](https://leetcode.com/problems/median-of-two-sorted-arrays/)

Given two sorted arrays, return the median of both arrays

*   Median = the middle number. If there's an even amount of numbers, then it's the average of the two

## Example
*   `Given nums = [2, 7, 11, 15], [1, 2, 5, 7, 9]` 
*   `Returns: 7`

## Bonus
An added bonus if it can be done in run time complexity should be O(log (m+n))

## Thoughts

A simple solution that is easy to read and understand is to simply merge the two arrays, re-sort the arrays and then get the median from the final array
