package puzzle

func isValidInRow(p Puzzle, n Place, y int) bool {
	for x := 0; x < Len; x++ {
		if p[y][x] == n {
			return false
		}
	}
	return true
}

func isValidInCol(p Puzzle, n Place, x int) bool {
	for y := 0; y < Len; y++ {
		if p[y][x] == n {
			return false
		}
	}
	return true
}

func isValidInBox(p Puzzle, n Place, xStart, yStart int) bool {
	// TODO: Don't use hard-coded threes here but set the upper limit
	// dynamically according to the grid size.
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if p[y+yStart][x+xStart] == n {
				return false
			}
		}
	}
	return true
}
func IsValid(p Puzzle, n Place, x, y int) bool {
	if isValidInRow(p, n, y) && isValidInCol(p, n, x) && isValidInBox(p, n, x-(x%3), y-(y%3)) {
		return true
	}
	return false
}
