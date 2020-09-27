package main

import (
	"fmt"

	// Imports off the project
	"github.com/aleluchesi/proceduresgo"

	// Imports within the project

	"github.com/aleluchesi/paygo/transferarea"
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

	//campos = elo.Parse{}

	fmt.Print("PAN: %v\r\n")

	fmt.Printf("Chip 2: %v\r\n", area.Chip)

	proceduresgo.Writeconsole("I", "Main", "Finish process")

}
