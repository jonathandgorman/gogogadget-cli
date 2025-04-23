package dice

import (
	"fmt"
	"math/rand"
	"time"
)

type SystemInfo struct {
	Hostname     string
	Platform     string
	OS           string
	Kernel       string
	CPU          string
	CPUCores     int32
	RAMTotal     uint64
	RAMAvailable uint64
	DiskTotal    uint64
	Uptime       string
}

func generateDiceFace(n int) {
	switch n {
	case 1:
		fmt.Println("   -------")
		fmt.Println("  |       |")
		fmt.Println("  |   *   |")
		fmt.Println("  |       |")
		fmt.Println("   -------")
	case 2:
		fmt.Println("   -------")
		fmt.Println("  | *     |")
		fmt.Println("  |       |")
		fmt.Println("  |     * |")
		fmt.Println("   -------")
	case 3:
		fmt.Println("   -------")
		fmt.Println("  | *     |")
		fmt.Println("  |   *   |")
		fmt.Println("  |     * |")
		fmt.Println("   -------")
	case 4:
		fmt.Println("   -------")
		fmt.Println("  | *   * |")
		fmt.Println("  |       |")
		fmt.Println("  | *   * |")
		fmt.Println("   -------")
	case 5:
		fmt.Println("   -------")
		fmt.Println("  | *   * |")
		fmt.Println("  |   *   |")
		fmt.Println("  | *   * |")
		fmt.Println("   -------")
	case 6:
		fmt.Println("   -------")
		fmt.Println("  | *   * |")
		fmt.Println("  | *   * |")
		fmt.Println("  | *   * |")
		fmt.Println("   -------")
	default:
	}
}

func Run(diceNumber int) {
	fmt.Println("\n   ___        ___        ___         _          _           ___  _        _ \n  / __|___   / __|___   / __|__ _ __| |__ _ ___| |_        |   \\(_)__ ___| |\n | (_ / _ \\ | (_ / _ \\ | (_ / _` / _` / _` / -_)  _|_ _ _  | |) | / _/ -_)_|\n  \\___\\___/  \\___\\___/  \\___\\__,_\\__,_\\__, \\___|\\__(_|_|_) |___/|_\\__\\___(_)\n                                      |___/                                 ")
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < diceNumber; i++ {
		randomNumber := rand.Intn(6) + 1
		generateDiceFace(randomNumber)
	}
}
