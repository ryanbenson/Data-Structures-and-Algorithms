# Binary Queue

Reference: Issue #122 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

**Given a positive number N, generate binary numbers between 1 to N using a queue-type structure in linear time.**

```console
> binaryQueue(10)
> [1,10,11,100,101,110,111,1000,1001,1010,1011,1100,1101,1110,1111,10000]
```

## Thoughts

Was a fun exercise in making a fake queue system (FIFO) and working with binary representations of our numbers again. The benchmark came out really well too. So I'm pleased with how it came out. I  may come up with another flavor later to compare the benchmarks, but we'll see.

```console
BenchmarkBinaryQueue-12     1649887       705 ns/op
```