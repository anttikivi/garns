package puzzle

import "fmt"

// findUnresolved searches the given Puzzle p and returns the x and y
// coordinates of the first unresolved place in the grid. If no unresolved
// places are found, it returns -1, -1.
func findUnresolved(p Puzzle) (int, int) {
	for y, row := range p {
		for x, b := range row {
			if b == Unresolved {
				return x, y
			}
		}
	}
	return -1, -1
}

func Solve(p Puzzle) (Puzzle, bool) {
	x, y := findUnresolved(p)
	if x == -1 || y == -1 {
		fmt.Println("Seems like the given input is already solved!")
		return p, true
	}

	for n := Place(1); n <= MaxValue; n++ {
		if IsValid(p, n, x, y) {
			p[y][x] = n
			if _, ok := Solve(p); ok {
				return p, true
			}
			p[y][x] = Unresolved
		}
	}

	return p, false
}
