# Mean Median Mode
Reference: Issue #134 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

Given a set of numbers, return the mean, median, and mode.

*   Mean = average of all of the numbers
*   Median = the middle number. If there's an even amount of numbers, then it's the average of the two
*   Mode = the number that occurs the most

## Example
*   `Given nums = [2, 7, 11, 15, 7]`
*   `Returns: { mean: 8, media: 7, mode: 7 }`

## Edge Cases
*   If there is more than one mode value, the mode would be an array of modes
*   If there's no repeated numbers, then mode is actually `nil`

## Thoughts
The exercise is pretty simple. It could be optimized to do all of the processing in one go, but I feel that making the separate analysis into separate functions is better as they can be isolated, run as an individual call, and tested in isolation. It also makes it a lot easier to read. Simplicity is always better. ✨