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

Was a fun and quick exercise! Was a quick exercise on thinking of different ways to manage the letters. A first round was to just make a quick map, which worked out well and easy to understand.

```console
BenchmarkGetFirstDuplicateLetter_Map-12    	 3720231	       312 ns/op
```