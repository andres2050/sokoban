package methods

// Position struct for box location
type Position struct {
	Row    int
	Column int
}

func availableMovements(playerPosition Position, mapGame []string) []Position {
	moves := make([]Position, 0)

	row := playerPosition.Row
	column := playerPosition.Column

	if column+1 < len(mapGame[row]) && mapGame[row][column+1:] != "W" {
		moves = append(moves, Position{
			Row:    row,
			Column: column + 1,
		})
	}
	if column-1 >= 0 && mapGame[row][column-1:] != "W" {
		moves = append(moves, Position{
			Row:    row,
			Column: column - 1,
		})
	}

	if row+1 < len(mapGame) && mapGame[row+1][column:] != "W" {
		moves = append(moves, Position{
			Row:    row + 1,
			Column: column,
		})
	}
	if row-1 >= 0 && mapGame[row-1][column:] != "W" {
		moves = append(moves, Position{
			Row:    row - 1,
			Column: column,
		})
	}
	return moves
}
