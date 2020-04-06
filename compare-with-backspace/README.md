# Compare With Backspace
Reference: Issue #138 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

> Given two strings n and m, return true if they are equal when both are typed into empty text editors. The twist: # means a backspace character.

## Example
*   `compareWithBackspace("a##c", "#a#c") => true (they both become 'c')`
*   `compareWithBackspace("xy##", "z#w#") => true (they both come '')`

## Thoughts
This was fun ✨and it was interesting to learn more about the string <=> rune relationship when processing the strings in Go. And to be careful about trying to trim an empty string.