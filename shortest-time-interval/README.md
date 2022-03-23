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

### Afterward

It was interesting building a count sort and getting it to work as expected. What is more interesting is that it performs worse by a large margin than letting Go do its own sorting. It was still a fascinating problem to solve and to do it in `O(n)`. Overall, I had a great time digging into this, and I learned something new!

```console
BenchmarkSmallestTimeInterval-12          	  118671	      9774 ns/op
BenchmarkSmallestTimeIntervalSimple-12    	 1000000	      1035 ns/op
```

The issue with this specific use case of a count sort is that there is a large number of possible minutes it has to look for. The timespan is 24 hours, so the potential max / min span is 24 hours, so 86K potential iterations. You could try to shrink that by imposing limits, like:

* The times have to be divisible by 5
* The times have to be 15 min increments

Doing so dramatically cuts down the min / max window, and you end up with a performant flavor of it. But it's still:

* Harder to read using the count sort due to the complexity of the algorithm
* Harder to maintain since you're running a custom sort algorithm
* Slower in this specific use case due to the paramters given of dealing with time

Doing benchmarks with more restrictive time spans does improve performance, but the `O(n)` flavor loses in the end compared to a more expensive time algorithm interestingly.

```console
BenchmarkSmallestTimeInterval-12                  	  116524	      9604 ns/op
BenchmarkSmallestTimeIntervalFifteenMinIncr-12    	  605389	      1931 ns/op
BenchmarkSmallestTimeIntervalFiveMinIncr-12       	  400156	      3032 ns/op
BenchmarkSmallestTimeIntervalSimple-12            	 1171320	      1020 ns/op
```

I also got curious if you threw a lot of possible times at it, and it's reflective of the issue still:

```console
BenchmarkSmallestTimeInterval-12                   	  112047	      9635 ns/op
BenchmarkSmallestTimeIntervalMany-12               	   48232	     24228 ns/op
BenchmarkSmallestTimeIntervalFifteenMinIncr-12     	  647314	      1843 ns/op
BenchmarkSmallestTimeIntervalFiveMinIncr-12        	  403718	      2927 ns/op
BenchmarkSmallestTimeIntervalFiveMinIncrMany-12    	  117955	     10108 ns/op
BenchmarkSmallestTimeIntervalSimple-12             	 1000000	      1015 ns/op
BenchmarkSmallestTimeIntervalSimpleMany-12         	  263340	      4516 ns/op
```

Overall, a count sort is a pretty helpful algorithm, but it feels like it's more suited for smaller number sequences in a more limited capacity. And possibly more so if you already know the boundaries, so you don't have to find them yourself. But that's also a limit because that's assuming you can know the outer bounds.
