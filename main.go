package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"syscall"

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
	mapGame, positions := loadLevel()
	method := iaMethod()

	executeMethod(method, mapGame, positions)
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

func loadLevel() ([]string, []position) {
	dat := os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	datBody, _ := ioutil.ReadAll(dat)
	syscall.Close(syscall.Stdin)

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

func executeMethod(method int, mapGame []string, positions []position) {
	switch method {
	case 1:
		profundidad(mapGame, positions)
	case 2:
		anchura(mapGame, positions)
	case 3:
		iterativa(mapGame, positions)
	}
}
