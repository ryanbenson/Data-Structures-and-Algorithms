# Is Golden Ratio

Reference: Issue #242 of [cassidoo newsletter](https://buttondown.email/cassidoo/archive/5-time-moves-slowly-but-passes-quickly-alice/) 🎉 (amazing newsletter by the way)

Given two strings n and m, return true if they are equal when both are entered into text editors. But: a `#` means a backspace character (deleting backwards), and a `%` means a delete character (deleting forwards).

Example:

```console
> equalWithDeletions("a##x", "#a#x")
> true      // both strings become "x"

> equalWithDeletions("fi##f%%%th %%year #time###", "fifth year time")
> false     // the first string becomes "fart"
```

## Thoughts

TBD

I did have an idea about possibly evaluating each character at a time that's determined, but given you can backspace and delete, that won't work well because a produced character might be deleted later. So, I realized the best way is to simply understand hthe final string then compare them.
