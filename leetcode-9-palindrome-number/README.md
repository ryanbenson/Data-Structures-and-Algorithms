# 9. Palindrome Number

[9. Palindrome Number](https://leetcode.com/problems/palindrome-number)

Given an integer x, return true if x is a palindrome, and false otherwise.

```
Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

## Results

```
Runtime 123ms Beats 84.17% of users with JavaScript
Memory 56.46MB Beats 91.92% of users with JavaScript
```

## Thoughts

Overall this was easy, but it's an `Easy` one so that makes sense. You could in theoery just convert to a string, then reverse the string and compare to the two. Going this route ends up requiring `O(n)` space since you need to hold the reversed version of the original string.

All of that isn't quite necessary. You really only need to loop through half the array and compare the items one by one. If you ever find anything not matching, you can short circuit. So in reality, you really only need to loop through at most half of the size of the array at worst.
