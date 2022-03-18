# Generate Arrow

Reference: Issue #201 of [cassidoo newsletter](https://buttondown.email/cassidoo/archive/there-will-always-be-enemies-time-to-stop-being/) 🎉 (amazing newsletter by the way)

**Given a direction and a number of columns, write a function that outputs an arrow of asterisks (see the pattern in the examples below)!**

```console
$ printArrow('right', 3)
Output:
*
 *
  *
 *
*

$ printArrow('left', 5)
Output:
    *
   *
  *
 *
*
 *
  *
   *
    *
```

## Thoughts

I had fun with this one! It was a mix of understanding how to build the output efficiently, but also understanding the rules and logic to build the spacing in different directions. I went with a single string builder using a loop. But you could alternatively use a recursive function, or build an array and compress it into a string. But I felt building it in cycle would be better than going over arrays multiple times. Overall it performs pretty well. I'm sure I could optimize it more, like using a different way of building the string, etc.

```console
BenchmarkCanPlantLong-12    	 2460577	       478 ns/op
```
