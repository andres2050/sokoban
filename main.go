package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/andres2050/sokoban/methods"
)

var (
	regexMap      = regexp.MustCompile(`[W0X]{1,}`)
	regexPosition = regexp.MustCompile(`([0-9]),([0-9])`)
)

func main() {
	mapGame, positions := loadLevel()
	executeMethod(mapGame, positions)
}

func loadLevel() ([]string, []methods.Position) {
	datBody, _ := ioutil.ReadAll(os.Stdin)
	worldStr := string(datBody)
	mapData := regexMap.FindAllString(worldStr, -1)
	positionFound := regexPosition.FindAllStringSubmatch(worldStr, -1)
	positions := make([]methods.Position, 0)

	for i := 0; i < len(positionFound); i++ {
		row, _ := strconv.Atoi(positionFound[i][1])
		column, _ := strconv.Atoi(positionFound[i][2])

		positions = append(positions, methods.Position{
			Row:    row,
			Column: column,
		})
	}

	return mapData, positions
}

func executeMethod(mapGame []string, positions []methods.Position) {
	switch method := os.Args[1]; method {
	case "profundidad":
		methods.DFS(mapGame, positions)
	case "anchura":
		methods.BFS(mapGame, positions)
	case "iterativa":
		methods.IDS(mapGame, positions)
	default:
		fmt.Println("Metodo incorrecto.")
	}
}
