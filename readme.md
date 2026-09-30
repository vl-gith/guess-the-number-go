# Guess the Number

A simple command-line number guessing game written in Go.

## About

The program generates a random number between 1 and 100. The player chooses a difficulty level that determines the number of available attempts.

After each incorrect guess, the game provides a hint indicating whether the target number is higher or lower.

At the end of the game, the program reveals the target number.

## Features

* Random number generation
* Three difficulty levels
* Different number of attempts depending on difficulty
* Higher/lower hints after each incorrect guess
* Input validation
* Target number revealed at the end of the game
* Command-line interface

## Difficulty Levels

| Level  | Attempts |
| ------ | -------: |
| Easy   |       10 |
| Medium |        5 |
| Hard   |        3 |

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
2. Choose a difficulty level:

   * `1` — Easy
   * `2` — Medium
   * `3` — Hard
3. Enter a number between 1 and 100.
4. Follow the hints after each incorrect guess.
5. Try to guess the target number before you run out of attempts.
6. The target number is revealed when the game ends.

## Strategy

The most efficient way to play is to use a **binary search strategy**.

Instead of guessing random numbers, always choose the middle of the remaining range:

```text
1–100   → guess 50
51–100  → guess 75
51–74   → guess 62
...
```

After each guess, use the hint to eliminate approximately half of the remaining numbers.

Since each guess approximately halves the search space, a number between 1 and 100 can always be found in at most:

```text
ceil(log₂(100)) = 7 guesses
```

The Easy difficulty provides 10 attempts, giving the player a margin for error. The Medium and Hard difficulties require a more efficient strategy.

## Technologies

* Go
* Standard Library (`fmt`, `math/rand`, `bufio`, `os`, `strconv`, `strings`)
