package methods

import "fmt"

// BFS: Breadth First Search
func BFS(mapGame []string, positions []Position) {
	playerPos := positions[0]
	goalPos := append(positions[:0], positions[1:]...)
	executeBFS(playerPos, mapGame, goalPos)
}

func executeBFS(playerPosition Position, mapGame []string, goalPos []Position) {
	if allGoals(mapGame, goalPos) {
		fmt.Println(playerPosition.Movements)
		return
	}
}
