A repository of coding challenges/tasks I came across in various job interviews.
I mostly do backend programming in Go and PHP, so here you can find some solutions.

The solutions to the tasks do not fulfill to be the shortest, fastest or cleanest solutions.

## Tasks

Each task should come with a solution plus some tests to verify.

### Task 1

Create a function `maskify` to mask digits of a credit card number with `#`.

**Requirements:**

- Do not mask the first digit and the last four digits
- Do not mask non-digit chars
- Do not mask if the input is less than 6
- Return '' when input is empty

### Task 2

Create a function `number_to_ordinal` to create an ordinal number for a given input.
Ordinal numbers in English have something like `st`, `nd`, `rd`, etc.

**Requirements:**

- Apply for number 1 to 1001... if that works any given number will do ;-)

### Task 3

Create a calculator for [Reverse Polish Notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation).
Write a `calculate` function that accepts an input and returns the result of the operation.

**Requirements:**

- Support the mathematical operations for `+`, `-`, `*` and `/`
- Check for invalid syntax, like `2 3+`. There is a space missing.
- Return 0 (integer) when nothing is entered
- Return the numeric value when no operand is given, like `1 2 3.5` return `3.5`

### Task 4

Given you have an array of numbers, you move inside the array by the value of the current element.
Write a function `jump_out_of_array` that outputs

- the amount of jumps until you jump out of the array
- `-1` when you reach the end of the array but do not jump out

**Requirements:**

- Array size is indefinite
- Array elements are integers, positive and negative

**Example:**

Given an array of `A[2, 3, -1, 1, 6, 4]`.

![](./docs/t4/task4.png)

- Jump 1: `A[0]` + `2` = `A[2]`
- Jump 2: `A[2]` + `(-1)` = `A[1]`
- Jump 3: `A[1]` + `3` = `A[4]`
- Jump 4: `A[4]` + `6` = out of range

So the result is `4`, you need `4` jumps to jump out of the array.

### Task 5

Find the k-complement pairs in an array of a given number. Write a function `k_complement` that that outputs the amount
of pairs.

**Requirements:**

_Do not_ use nested loops to solve this problem, because of a time complexity of the loop solution.
[Check this thread](https://stackoverflow.com/questions/11032015/how-to-find-time-complexity-of-an-algorithm) to see what time complexity of an algorithm means.

**Example:**

- `A[0]` + `A[8]` = `1` + `5` = `6`
- `A[1]` + `A[6]` = `8` + `-2` = `6`
- `A[4]` + `A[8]` = `1` + `5` = `6`
- `A[5]` + `A[5]` = `3` + `3` = `6`
- `A[5]` + `A[5]` = `3` + `3` = `6`
- `A[6]` + `A[1]` = `-2` + `8` = `6`
- `A[8]` + `A[0]` = `5` + `1` = `6`

The result here is `7`.

### Task 6

Calculate the [Fibonacci number](https://en.wikipedia.org/wiki/Fibonacci_number) of a given number
and return the last `6` non-zero numbers.

**Requirements:**

* Throw an exception when passing in a negative number

**Example:**

* `F8` = `21`, return `21`
* `F38` = `39088169`, return `88169`

## Contribute

Feel free to contribute. Use the issue list to propose new tasks or open PRs. Just provide proper tests
and description and requirements for the tasks.
