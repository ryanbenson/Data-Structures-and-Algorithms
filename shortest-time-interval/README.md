# Shortest Time Interval

Reference: Issue #240 of [cassidoo newsletter](https://buttondown.email/cassidoo/archive/the-best-time-to-make-friends-is-before-you-need/) 🎉 (amazing newsletter by the way)

**Given a list of times in a 24-hour period, return the smallest interval between two times in the list. Hint: you can do this in O(n) time! Assume the list is unsorted.**

```console
$ smallestTimeInterval(['01:00', '08:15', '11:30', '13:45', '14:10', '20:05'])
$ '25 minutes'
```

## Thoughts

### Initial Thoughts

Just getting started, but the interesting bit is the `O(n)` time given the list should be assume it is unsorted. If it's sorted, then it's trivial. A naive way to do it is to do a normal `sort()`, but most sorts are `(O(n log n))` (heap, merge, etc.). But I found [Count Sort](https://iq.opengenus.org/time-and-space-complexity-of-counting-sort/) which is `O(n)` which will help! [Another interesting article and break down on it too.](https://www.interviewcake.com/concept/java/counting-sort)

> In the worst-case scenario, the complexity of this algorithm is as follows:
> 
> * Time complexity: O(n + k), where n is the numbers of integers in your original array and k is the range of your min and max values.
> * Space complexity: O(n + k)
>
> When k is less or equal to n, both complexities reduce to O(n)
