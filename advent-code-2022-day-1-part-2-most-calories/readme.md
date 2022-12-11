# Advent Code 2022 - Day 1 - Part 1

--- Day 1: Calorie Counting - Part 2 ---

https://adventofcode.com/2022/day/1#part2

```bash
#getCalories x 754,125 ops/sec ±1.50% (90 runs sampled)
#getCaloriesRanked x 826,980 ops/sec ±1.75% (94 runs sampled)
#getCaloriesLargeDataset x 4,096 ops/sec ±1.65% (89 runs sampled)
#getCaloriesRankedLargeDataset x 4,819 ops/sec ±1.71% (88 runs sampled)
```

## Notes

Doing the ranking manually is fairly more efficient for time and memory compared to maintaining an array. It's about 15-17% more efficient with large data sets, and around 10% more efficient with smaller ones.
