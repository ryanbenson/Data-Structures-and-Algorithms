# Convert Number to Roman Numeral

Reference: Issue #7 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

> Write a function that takes in a string and returns the first duplicate character it finds.

Example:

```console
getFirstDuplicateLetter("boop") => "o"
getFirstDuplicateLetter("hellothere") => "l"
getFirstDuplicateLetter("nohip") => ""
```

## Thoughts

Was a fun and quick exercise! Was a quick exercise on thinking of different ways to manage the letters. A first round was to just make a quick map, which worked out well and easy to understand. I decided to try an even more simple approach where you loop through each letter and just find the count of each letter occurance and break if it's over one. Seems to be even more readable, and performs better.

```console
BenchmarkGetFirstDuplicateLetter_Map-12       	 3705500	       316 ns/op
BenchmarkGetFirstDuplicateLetter_Search-12    	 9802736	       121 ns/op
```