package main

import (
	"fmt"

	// Imports off the project
	proceduresgo "github.com/aleluchesi/ProceduresGo"

	// Imports within the project

	"github.com/aleluchesi/PayGo/model"
	"github.com/aleluchesi/PayGo/transferarea"
)

var (
	teste   bool
	comprou bool
)

func main() {

	proceduresgo.Writeconsole("I", "Main", "Start process")

	area := transferarea.Msgtype{}

	fmt.Printf("Chip 1: %v\r\n", area.Chip)

	area.Chip = true

	Campos := new(model.FieldsIso)

	Campos.Pan = "123"

	fmt.Printf("campos Entrada: %p - %-v\r\n", &Campos, Campos)

	//elo.Parse(&Campos)

	fmt.Printf("campos Entrada: %p - %-v\r\n", &Campos, Campos)

	//	fmt.Printf("PAN: %v\r\n", Camposentrada.GetPan())

	fmt.Printf("Chip 2: %v\r\n", area.Chip)

	proceduresgo.Writeconsole("I", "Main", "Finish process")

}
