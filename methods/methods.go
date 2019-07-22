package methods

// Position struct for box location
type Position struct {
	Row               int
	Column            int
	Movements         string
	MovementsPosition []Position
}

func availableMovements(playerPosition Position, mapGame []string) []Position {
	moves := make([]Position, 0)

	row := playerPosition.Row
	column := playerPosition.Column
	movements := playerPosition.Movements
	movementsPosition := append(playerPosition.MovementsPosition, playerPosition)

	if column+1 < len(mapGame[row]) && mapGame[row][column+1:] != "W" {
		newPos := Position{
			Row:               row,
			Column:            column + 1,
			Movements:         movements + "R",
			MovementsPosition: movementsPosition,
		}
		if !positionExist(newPos, movementsPosition) {
			moves = append(moves, newPos)
		}
	}
	if column-1 >= 0 && mapGame[row][column-1:] != "W" {
		newPos := Position{
			Row:               row,
			Column:            column - 1,
			Movements:         movements + "L",
			MovementsPosition: movementsPosition,
		}
		if !positionExist(newPos, movementsPosition) {
			moves = append(moves, newPos)
		}
	}

	if row+1 < len(mapGame) && mapGame[row+1][column:] != "W" {
		newPos := Position{
			Row:               row + 1,
			Column:            column,
			Movements:         movements + "D",
			MovementsPosition: movementsPosition,
		}
		if !positionExist(newPos, movementsPosition) {
			moves = append(moves, newPos)
		}
	}
	if row-1 >= 0 && mapGame[row-1][column:] != "W" {
		newPos := Position{
			Row:               row - 1,
			Column:            column,
			Movements:         movements + "U",
			MovementsPosition: movementsPosition,
		}
		if !positionExist(newPos, movementsPosition) {
			moves = append(moves, newPos)
		}
	}
	return moves
}

func positionExist(playerPosition Position, movementsPosition []Position) bool {
	for _, pos := range movementsPosition {
		if pos.Row == playerPosition.Row && pos.Column == playerPosition.Column {
			return true
		}
	}

	return false
}
