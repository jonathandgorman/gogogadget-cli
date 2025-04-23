// main.go
package main

import (
	"flag"
	"fmt"
	"go-go-gadget-cli/internal/dice"
	"go-go-gadget-cli/internal/help"
	"go-go-gadget-cli/internal/periscope"
)

func main() {
	var helpFlag bool
	var periscopeFlag bool
	var diceFlag bool
	var diceRollsFlag int

	flag.BoolVar(&helpFlag, "help", false, "Show help")
	flag.BoolVar(&helpFlag, "h", false, "Show help (shorthand)")

	flag.BoolVar(&periscopeFlag, "periscope", false, "Activate the periscope gadget")
	flag.BoolVar(&periscopeFlag, "p", false, "Activate the periscope gadget (shorthand)")

	flag.BoolVar(&diceFlag, "dice", false, "Activate the dice gadget")
	flag.BoolVar(&diceFlag, "d", false, "Activate the dice gadget (shorthand)")
	flag.IntVar(&diceRollsFlag, "rolls", 1, "Number of dice rolls (default: 1)")
	flag.IntVar(&diceRollsFlag, "n", 1, "Number of dice rolls (default: 1)(shorthand)")

	flag.Parse()

	if helpFlag {
		help.Run()
	} else if periscopeFlag {
		periscope.Run()
	} else if diceFlag {
		dice.Run(diceRollsFlag)
	} else {
		fmt.Println("No gadget specified. Use --help for options.")
		return
	}
}
