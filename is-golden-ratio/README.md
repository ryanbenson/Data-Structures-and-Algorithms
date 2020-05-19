# Is Golden Ratio
Reference: Issue #143 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

Given two numbers X and Y, write a function to check that X and Y are in golden ratio.

Two numbers are in the golden ratio if their ratio is the same as the ratio of their sum to the larger of the two quantities. [Check out the formula here](https://en.wikipedia.org/wiki/Golden_ratio).

Example:

```console
isGolenRatio(1, 2)
> false

isGoldenRatio(19.416, 12)
> true

isGoldenRatio(2, 5)
> false

isGoldenRatio(1, 0.618)
> true

isGoldenRatio(61.77, 38.176)
> true
```

## Thoughts
This was fun! Was an exercise in figuring out the algorithm for the ratio. The only hiccup was finding that there isn't a native way to have a float round to a specific decimal level, so I had to roll my own and be mindful of how it rounds. My initial thought process was to assume `x` was larger and flip the values if that proved false. It seemed like the most effective approach and better than using an array and sorting, even if it's 2 elements. But I saw some chatter about it, so I thought I'd give it a go and see how they compare. It's interesting the impact it has, and it doesn't really bring a lot of legibility to the code either.

### Benchmarks

```console
BenchmarkIsGoldenRatio-12         	115293860	        10.4 ns/op
BenchmarkIsGoldenRatioArray-12    	16713176	        70.3 ns/op
```
