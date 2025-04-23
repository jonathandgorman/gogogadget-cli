package help

import (
	"flag"
	"fmt"
	"os"
)

func Run() {
	fmt.Println("\n   ___        ___        ___         _          _      ___ _    ___ \n  / __|___   / __|___   / __|__ _ __| |__ _ ___| |_   / __| |  |_ _|\n | (_ / _ \\ | (_ / _ \\ | (_ / _` / _` / _` / -_)  _| | (__| |__ | | \n  \\___\\___/  \\___\\___/  \\___\\__,_\\__,_\\__, \\___|\\__|  \\___|____|___|\n                                      |___/                         ")
	fmt.Println("Usage: gogogadget [options]")
	flag.PrintDefaults()
	os.Exit(0)
}
