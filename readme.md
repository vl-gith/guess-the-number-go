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

## Technologies

* Go
* Standard Library (`fmt`, `math/rand`, `bufio`, `os`, `strconv`, `strings`)