# Guess the Number

A simple command-line number guessing game written in Go.

## About

The program generates a random number between 1 and 100. Before starting the game, the player enters their name and chooses a difficulty level that determines the number of available attempts.

The player's name is validated to ensure that it contains only letters.

After each incorrect guess, the game provides a hint indicating whether the target number is higher or lower.

The game also calculates the player's score based on the selected difficulty and the number of attempts used.

At the end of the game, the program reveals the target number and displays the final score.

## Features

* Player name input and validation
* Support for names containing letters from different alphabets
* Random number generation
* Three difficulty levels
* Different number of attempts depending on difficulty
* Higher/lower hints after each incorrect guess
* Input validation
* Score calculation based on difficulty and attempts used
* Higher score for fewer attempts
* Target number revealed at the end of the game
* Command-line interface

## Difficulty Levels

| Level  | Attempts | Maximum Score |
| ------ | -------: | ------------: |
| Easy   |       10 |          1800 |
| Medium |        5 |          3000 |
| Hard   |        3 |          7000 |

## Scoring

The player's score depends on the selected difficulty and the number of attempts used.

The maximum score is determined by the difficulty level:

```text
Easy   → 1800 points
Medium → 3000 points
Hard   → 7000 points
```

The final score is calculated using the following formula:

```text
Score = Maximum Score / Attempts Used
```

For example, on the Easy difficulty:

```text
1 attempt  → 1800 points
2 attempts →  900 points
4 attempts →  450 points
```

If the player fails to guess the target number, the final score is `0`.

The scoring system rewards players for solving the game using as few attempts as possible.

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
2. Enter your name.

   * The name must contain letters only.
3. Choose a difficulty level:

   * `1` — Easy
   * `2` — Medium
   * `3` — Hard
4. Enter a number between 1 and 100.
5. Follow the hints after each incorrect guess.
6. Try to guess the target number before you run out of attempts.
7. The target number is revealed when the game ends.
8. If you win, your final score is displayed.

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

Using fewer attempts also results in a higher score.

## Technologies

* Go
* Standard Library (`fmt`, `math/rand`, `bufio`, `os`, `strconv`, `strings`, `unicode`)
