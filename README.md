# Garns

A simple Sudoku-solver named after Howard Garns, the creator of Number Place
which became later known as Sudoku. You can provide it an input file containing
the starting position of a Sudoku puzzle. By default, `garns` outputs the
solution into `stdout`.

## Getting started

To get started, build the project.

    go build -o garns main.go

## Solving puzzles

The command-line interface takes a file containing the puzzle as input. The
puzzle must be nine lines of nine character. Each known number in the puzzle
must be given as is in the input. Each unknown place can be either a dot (`.`)
or a zero (`0`).

To solve a puzzle from an input file, run the `solve` command and the `--file`
options to specify the input file.

    garns solve -f path/to/input.txt

## Licence

This project is licensed under the Apache License 2.0. For more information,
please see the [LICENSE](LICENSE) file.
