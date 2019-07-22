package methods

// Position struct for box location
type Position struct {
	Row       int
	Column    int
	Movements string
}

func availableMovements(playerPosition Position, mapGame []string) []Position {
	moves := make([]Position, 0)

	row := playerPosition.Row
	column := playerPosition.Column
	movements := playerPosition.Movements

	if column+1 < len(mapGame[row]) && mapGame[row][column+1:] != "W" {
		moves = append(moves, Position{
			Row:       row,
			Column:    column + 1,
			Movements: movements + "R",
		})
	}
	if column-1 >= 0 && mapGame[row][column-1:] != "W" {
		moves = append(moves, Position{
			Row:       row,
			Column:    column - 1,
			Movements: movements + "L",
		})
	}

	if row+1 < len(mapGame) && mapGame[row+1][column:] != "W" {
		moves = append(moves, Position{
			Row:       row + 1,
			Column:    column,
			Movements: movements + "D",
		})
	}
	if row-1 >= 0 && mapGame[row-1][column:] != "W" {
		moves = append(moves, Position{
			Row:       row - 1,
			Column:    column,
			Movements: movements + "U",
		})
	}
	return moves
}
