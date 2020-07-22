# Convert Number to Roman Numeral

Reference: Issue #1 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

> Write a function that takes in a number from 1 to 1000 and returns that number in Roman Numerals.

Example:

```console
fromInt(13) => XIII
```

## Thoughts

Was a fun exercise. The conversion table was simple to figure out, then it was a simple exercise of how to best process a number into the proper format of a roman numeral given how they're created and their rules. Once the table and decreasing the number logic was thought through, the rest was easy. The simple and readable solution leverages a table, but there are probably more efficient ways to process it, even if there's a cost to readability that I might explore.

```console
BenchmarkFromInt-12    	12099063	        85.7 ns/op
```