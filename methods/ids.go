package methods

import "fmt"

// IDS: iterative deepening search
func IDS(mapGame []string, positions []Position) {
	playerPos := positions[0]
	goalPos := append(positions[:0], positions[1:]...)
	executeIDS(playerPos, mapGame, goalPos)
}

func executeIDS(playerPosition Position, mapGame []string, goalPos []Position) {
	if allGoals(mapGame, goalPos) {
		fmt.Println(playerPosition.Movements)
		return
	}
}
