# 169. Majority Element

[169. Majority Element](https://leetcode.com/problems/majority-element)

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

```
Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2
```

Follow-Up: Follow-up: Could you solve the problem in linear time and in O(1) space?

## Thoughts

Keeping the `O(1) space` requirement was pretty easy. My initial implementation was `O(n log n) time` which is pretty good, but it's mostly because it sorts the array to make it easier.

Found a way to gete `O(n) time and space` so it's an improvement because we end up making a map to manage the information. But we can still get it down to `O(1) space and O(n) time`

Finally, I did find a way to do it in `O(n) time and O(1) space` using a voting algorithm using `Moore's Voting Algorithm`. It's a pretty clever way to approach the problem.
