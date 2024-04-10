# Valid Syntax

[Reference](https://leetcode.com/problems/valid-parentheses/)

Given a string containing just the characters `'(', ')', '{', '}', '[' and ']'`, determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

## Examples

Example 1:

```console
Input: "()"
Output: true
```

Example 2:

```console
Input: "()[]{}"
Output: true
```

Example 3:

```console
Input: "(]"
Output: false
```

Example 4:

```console
Input: "([)]"
Output: false
```

Example 5:

```console
Input: "{[]}"
Output: true
```

## Result

Beats 75.84% of users with Go

## Thoughts

This was fun! I explored a more robust way of checking it that avoids manually hard coding the expectations while iterating through the characters. It leverages a map of the open and close expectations and constantly checks that. I did an alternative version that hard codes the expected open and close handling, which proved to be more efficient by a fair amount, but also ends up being more brittle and includes more magic strings.

It's a little over double the performance. So it's substantial, but at the loss of maintainability.

```console
BenchmarkValidSyntax-12              5200888       217 ns/op
BenchmarkValidSyntaxMagic-12        14785834        80.8 ns/op
```

I then decided to make a variation of the original concept, that does that mapping and checking systematically via a dictionary versus manaully hard coding the magic strings. This time using the runes that you get when you iterate over each character in the sequence, and not converting them to a string. They pass all of the tests, and the benchmarks are quite interesting.

The rune version of the lookups via the dictionary are almost the same in terms of performance verses the less mantainable magic string flavor. In that case the very minor impact to the performance would be worthwhile so you can simply manage a dictionary.

```console
BenchmarkValidSyntax-12              5255598       217 ns/op
BenchmarkValidSyntaxRune-12         13280041        89.8 ns/op
BenchmarkValidSyntaxMagic-12        14479861        82.4 ns/op
```
