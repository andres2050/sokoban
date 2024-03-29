package methods

import "fmt"

// DFS: Depth First Search
func DFS(mapGame []string, positions []Position) {
	playerPos := positions[0]
	goalPos := append(positions[:0], positions[1:]...)
	executeDFS(playerPos, mapGame, goalPos)
}

func executeDFS(playerPosition Position, mapGame []string, goalPos []Position) {
	if allGoals(mapGame, goalPos) {
		fmt.Println(playerPosition.Movements)
		return
	}

		
}
