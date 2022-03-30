# Product List

Reference: Issue #182 of [cassidoo newsletter](https://buttondown.email/cassidoo/archive/without-struggle-there-is-no-progress-frederick/) 🎉 (amazing newsletter by the way)

**Implement a ProductList class that has two methods, add(n) (which pushes the value n to the back of the list) and product(m) (which returns the product of the last m numbers in the list)**

```console
ProductList p = new ProductList();
p.add(7);         // [7]
p.add(0);         // [7,0]
p.add(2);         // [7,0,2]
p.add(5);         // [7,0,2,5]
p.add(4);         // [7,0,2,5,4]
p.product(3);     // return 40 because 2 * 5 * 4
```

## Thoughts

This was a pretty simple exercise. The notable thing was knowing how to modify the list of products within the struct, otherwise you end up resetting the array on every call. So, if you don't use a pointer, by time you do a `get` call, there's nothing in the array. It was fun though! And since it was so simple, the performance is high:

```console
BenchmarkProductListAddGet-12    	 8652214	       132 ns/op
```
