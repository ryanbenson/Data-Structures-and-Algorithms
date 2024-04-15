# 55. Jump Game

[55. Jump Game](https://leetcode.com/problems/jump-game)

You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

```
Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
```

(This isn't in the official description, but helpful to understand)

```
Example 3:
Input: nums = [5,9,3,2,1,0,2,3,3,1,0,0]
Output: true
Explanation: You go from 5 -> 9, you then go from 9 to the last 3, skipping only 7 steps, then that 3 takes you to the end
```

Thoughts:

This is a [greedy algorithm](https://en.wikipedia.org/wiki/Greedy_algorithm). You essentially use the highest number you can find to meet your goal, and keep going until you run out of jumps or reach the end. The trick is to keep finding the biggest number, and seeing if that will take you to the end on each one you find.
