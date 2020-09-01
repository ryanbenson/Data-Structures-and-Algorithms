# Rotate Array

Reference: Issue #9 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

**Given a list of menu items and their prices, and the amount in your wallet, return all combinations of items that you can buy.**

```console
menu = { 'taco' : 1, 'burger' : 2, 'shawarma' : 3 }
wallet = 3
> combo(menu, wallet)
['taco', 'taco', 'taco']
['taco', 'burger']
['shawarma']
```

## Questions

As I'm working through this, I'm noting the I/O from the example and I'm wondering, should combinations that don't consume all of the wallet be acceptable? Example:

```console
menu = { 'taco' : 1, 'burger' : 2, 'shawarma' : 3 }
wallet = 3
> combo(menu, wallet)
['taco', 'taco', 'taco']
['taco', 'burger']
['burger']
['shawarma']
```

Noting the addition of just a `burger`, leaving the wallet with a value of `1`
