# Largest Island Rectange

Reference: Issue #237 of [cassidoo newsletter](https://buttondown.email/cassidoo/archive/change-your-thoughts-and-you-change-your-world/) 🎉 (amazing newsletter by the way)

**Given a 2D array of 0s and 1s, return the size of the largest rectangle of 1s in the array.**

```console
let islands = [
    [0,0,0,1,0]
    [1,1,0,0,1]
    [1,1,0,0,0]
    [0,0,1,0,0]
]

$ largestRect(islands)
$ '2x2'
```

## Thoughts

This was fun! And it was more complicated than it appears on the surface. And there were a few ways to approach it. I opted to go for checking for any possible rectangles and then just keeping the biggest one around. Alternatively, you could hold them all as objects then find the biggest one later, but the aforementioned way seemed both faster and memory efficient given you don't need to hold unnecessary data, or bother having to clean them up as they become unneeded.

Overall, it was a fun thought experiment and had me thinking different ways to approach the problem. There are some edge cases that seem interesting too that I might pursue later (e.g. unfinished rectangles and donuts)

```
// unfinished island
let islands = [
    [0,0,0,1,0]
    [1,1,1,0,1]
    [1,1,0,0,0]
    [1,1,1,0,0]
]

// donut island
let islands = [
    [0,0,0,1,0]
    [1,1,1,0,1]
    [1,0,1,0,0]
    [1,1,1,0,0]
]
```

But the acceptance criteria was around filled rectangles, so I'll leave it at that for the time being.

Overall it runs pretty fast though, though the time scale seems to increase as the map size expands, which is something to consider in terms of optimization.

```
BenchmarkLargestRect-12                 15627930                75.24 ns/op
BenchmarkLargestRectLargeMap-12          3670162               332.8 ns/op
```

Some ideas to improve performance:

* After finding the island size, essentially block off / delete those coordinates from the map so you don't evaluate them
* If you find a large island, and you reach the end of a part of the map where it's not possible to find a larger island, then it can short circuit