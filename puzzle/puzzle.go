package puzzle

import (
	"errors"
	"fmt"
	"unicode"
)

type Place byte

// TODO: It might be strange to use 'uint8' but let's try it since it uses
// less memory.
type Puzzle [][]Place

const Unresolved Place = 0
const Len = 9

func Parse(input []string) (Puzzle, error) {
	var p Puzzle
	for _, line := range input {
		row := make([]Place, 0, len(line))
		for _, c := range line {
			if unicode.IsDigit(c) {
				row = append(row, Place(c-'0'))
			} else if c == '.' {
				row = append(row, 0)
			} else {
				return nil, errors.New("Invalid input character '" + string(c) + "'")
			}
		}
		p = append(p, row)
	}
	return p, nil
}
