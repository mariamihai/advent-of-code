package day10

func isConnectionN(pipe string) bool {
	return pipe == "|" || pipe == "F" || pipe == "7"
}
func isConnectionS(pipe string) bool {
	return pipe == "|" || pipe == "L" || pipe == "J"
}
func isConnectionE(pipe string) bool {
	return pipe == "-" || pipe == "L" || pipe == "F"
}
func isConnectionW(pipe string) bool {
	return pipe == "-" || pipe == "J" || pipe == "7"
}

func checkN(puzzle [][]string, x, y int) (bool, *Point) {
	if y-1 >= 0 {
		currentPipe := puzzle[y-1][x]

		if isConnectionN(currentPipe) {
			return true, &Point{
				X: x,
				Y: y - 1,
			}
		}
	}

	return false, nil
}

func checkS(puzzle [][]string, x, y int) (bool, *Point) {
	if y+1 <= len(puzzle) {
		currentPipe := puzzle[y+1][x]

		if isConnectionS(currentPipe) {
			return true, &Point{
				X: x,
				Y: y + 1,
			}
		}
	}

	return false, nil
}

func checkE(puzzle [][]string, x, y int) (bool, *Point) {
	if x-1 >= 0 {
		currentPipe := puzzle[y][x-1]

		if isConnectionE(currentPipe) {
			return true, &Point{
				X: x - 1,
				Y: y,
			}
		}
	}

	return false, nil
}

func checkW(puzzle [][]string, x, y int) (bool, *Point) {
	if x+1 <= len(puzzle[y]) {
		currentPipe := puzzle[y][x+1]

		if isConnectionW(currentPipe) {
			return true, &Point{
				X: x + 1,
				Y: y,
			}
		}
	}

	return false, nil
}
