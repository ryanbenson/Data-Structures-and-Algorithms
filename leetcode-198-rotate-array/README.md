# 189. Rotate Array

[189. Rotate Array](https://leetcode.com/problems/rotate-array)

Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

```
Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
```

Follow up:

Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with O(1) extra space?

## Results

```
Runtime 87ms Beats 32.66% of users with JavaScript
Memory 57.14MB Beats 64.90% of users with JavaScript
```

## Thoughts

I did manage to do it in `O(1)` space on the first attempt, and it's a quite legible version too.

> The time complexity of this solution is O(n) because we are performing a splice operation on the array, which has a time complexity of O(n), and then performing an unshift operation, which also has a time complexity of O(n).
>
> The space complexity is O(1) because we are not using any extra space that grows with the input size. We are modifying the input array in-place.

Overall my approach was simple. You just figure out what you need to swap positions, and understand that you may end up with a `k` with a lot higher number than the length of the array, so you use `%` to help you out, then chunk out the piece you need and move it in place.
