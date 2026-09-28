# Guess the Number

A simple command-line number guessing game written in Go.

## About

The program generates a random number between 1 and 100. The player has 10 attempts to guess the number.

After each incorrect guess, the game provides a hint indicating whether the target number is higher or lower.

## Features

* Random number generation
* 10 attempts to guess the number
* Higher/lower hints after each incorrect guess
* Command-line interface

## Requirements

* Go 1.20 or later

## Getting Started

Clone the repository:

```bash
git clone https://github.com/vl-gith/guess-the-number-go
cd guess-the-number-go
```

Run the game:

```bash
go run guess-number.go
```

## How to Play

1. Run the program.
2. Enter a number between 1 and 100.
3. Follow the hints after each incorrect guess.
4. Guess the number within 10 attempts to win!

## Strategy

The most efficient way to play is to use a **binary search strategy**.

Instead of guessing random numbers, always choose the middle of the remaining range:

```text
1–100   → guess 50
51–100  → guess 75
51–74   → guess 62
...
```

After each guess, use the hint to eliminate half of the remaining numbers.

Since each guess approximately halves the search space, a number between 1 and 100 can always be found in at most:

```text
ceil(log₂(100)) = 7 guesses
```

The game provides 10 attempts, giving the player a small margin for error.

## Technologies

* Go
* Standard Library (`fmt`, `math/rand`, `bufio`, `os`, `strconv`, `strings`)
