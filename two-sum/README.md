# Two Sum
[Reference](https://leetcode.com/problems/two-sum/)

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

## Example
```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## Thoughts
Initial pass is to just loop through each number, and then for each number loop through again to find its needed match to get the target. But that requires a lot of extra looping, but does solve the answer.

I optimized it to make a map of every number it makes `[]{index : value}`. So as it goes, it has a record of every number and its index. And we know the number we need by subtracting the target from the current number. If it exists, we have our answer. If not, add it to our map and keep looking.