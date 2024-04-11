# 26. Remove Duplicates from Sorted Array

[Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array)

Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

- Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
- Return k.

Custom Judge:

The judge will test your solution with the following code:

```
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```

If all assertions pass, then your solution will be accepted.

```
Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

## Results

```
Runtime: 68ms Beats 43.35% of users with JavaScript
Memory 52.36MB Beats 48.31% of users with JavaScript
```

## Thoughts

Overall it was pretty easy, even if it's another weird problem because it's not good practice to mutate the existing params you are given. But it was easy. You basically start on item 1 in the array, you keep going through the list checking to see if the previous number changes at all. If it does, you update the next item in the sequence. But you only bother updating the early parts of the array because the latter numbers are tossed and don't matter.

So you have a pointer for the items in the array to replace, and you only increment that pointer whenever a new number comes in with the main loop.

It's a weird problem, and likely nothing that would ever realistically come up in real life. But it works.
