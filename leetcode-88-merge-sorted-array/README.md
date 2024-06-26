## 88. Merge Sorted Array

[Leetcode 88 Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array)

You are given two integer arrays `nums1` and `nums2`, sorted in non-decreasing order, and two integers `m` and `n`, representing the number of elements in `nums1` and `nums2` respectively.

Merge `nums1` and `nums2` into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array `nums1`. To accommodate this, `nums1` has a length of `m` + `n`, where the first m elements denote the elements that should be merged, and the last `n` elements are set to 0 and should be ignored. `nums2` has a length of `n`.

## Results

```
Runtime: 28ms
Beats 99.98% of users with JavaScript
```

That runtime win!

## Thoughts

Overall quite easy. Mostly just use splicing to do the magic for you, and overall it's a really simple way to approach it.

There's a bonus to do it in O(m+n), and I did manage to do it. Ironically though, the O(m+n) solution performed worse:

```
Runtime: 49ms
Beats 79.00% of users with JavaScript
```

You basically take the 2 pre-sorted arrays, look at the ends, take the higher number, and then move that into the latter position, and move your pointer down. You keep doing that until you're at the end of your pointer.
