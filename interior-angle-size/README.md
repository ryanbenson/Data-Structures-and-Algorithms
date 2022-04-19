# Interior Angle Size

Reference: Issue #243 of [cassidoo newsletter](https://buttondown.email/cassidoo/archive/you-become-what-you-understand-soren-kierkegaard/) 🎉 (amazing newsletter by the way)

Given an integer n representing the number of sides of a [regular polygon](https://en.wikipedia.org/wiki/Regular_polygon), return the measure of each interior angle. Bonus points: implement some of the other functions listed in the linked Wikipedia page!

Example:

```console
getInteriorAngle(3)
$ 60 // Each angle in a triangle that is a regular polygon is 60 degrees

getInteriorAngle(8)
$ 135
```

## Thoughts

This was neat, but really simple. It's a pretty standardized algorithm to figure out the angle. Since it was so simple, I'll keep adding more methods!

And since it's so simple, the performance is high.

```console
BenchmarkIsEqualSimple-12    	1000000000	         0.524 ns/op
```
