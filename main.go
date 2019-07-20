package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/inancgumus/screen"
)

var (
	regexMap      = regexp.MustCompile(`[W0X]{1,}`)
	regexPosition = regexp.MustCompile(`([0-9]),([0-9])`)
)

type position struct {
	Row    int
	Column int
}

func main() {
	level := os.Args[1]
	fmt.Println(level)
	//method := iaMethod()

	//executeMethod(method, level)
}

func iaMethod() int {
	var method int
	screen.Clear()

	fmt.Println("Seleccione el metodo que desea:")
	fmt.Println("1) Preferente por profundidad")
	fmt.Println("2) Preferente por anchura")
	fmt.Println("3) Profundidad iterativa")
	fmt.Print("Ingrese su opcion: ")
	_, err := fmt.Scan(&method)

	if err != nil || method < 1 || method > 3 {
		return iaMethod()
	}

	return method
}

/*func selectLevel() int {
	var level int
	screen.Clear()

	fmt.Println("Seleccione el nivel deseado:")
	fmt.Println("1) Nivel 1")
	fmt.Println("2) Nivel 2")
	fmt.Println("3) Nivel 3")
	fmt.Println("4) Nivel 4")
	fmt.Print("Ingrese su opcion: ")
	_, err := fmt.Scan(&level)

	if err != nil || level < 1 || level > 4 {
		return selectLevel()
	}

	return level
}*/

func loadLevel(level string) ([]string, []position) {
	levelName := fmt.Sprintf("levels/%s", level)

	dat, _ := os.Open(levelName)
	datBody, _ := ioutil.ReadAll(dat)
	worldStr := string(datBody)
	mapData := regexMap.FindAllString(worldStr, -1)
	positionFound := regexPosition.FindAllStringSubmatch(worldStr, -1)
	positions := make([]position, 0)

	for i := 0; i < len(positionFound); i++ {
		row, _ := strconv.Atoi(positionFound[i][1])
		column, _ := strconv.Atoi(positionFound[i][2])

		positions = append(positions, position{
			Row:    row,
			Column: column,
		})
	}

	return mapData, positions
}

func executeMethod(method int, level string) {
	mapGame, positions := loadLevel(level)

	switch method {
	case 1:
		profundidad(mapGame, positions)
	case 2:
		anchura(mapGame, positions)
	case 3:
		iterativa(mapGame, positions)
	}
}
