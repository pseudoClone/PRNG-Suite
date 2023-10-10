# Linear Congruency Generator (LCG)

### Equation
```math
x_{n+1} \equiv (a x_n + c)\mod m
\\
\text{Where x} _{n}{= seed} \text{, c = constant, a = multiplier, n = iteration}
```
### Case:
- The multiplier value is taken from Thomas Clausen.
- The initial seed can be taken from a constant value too.

### Good Guidelines
- m and c => relatively prime as possible
- a - 1 is divisible by all prime factors of m
- if m is divisible by 4, a - 1 is also divisible by 4
