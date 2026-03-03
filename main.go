package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"atlas.screensaver/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

var Version = "dev"

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Printf("atlas.screensaver v%s\n", Version)
		return
	}
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help") {
		fmt.Println("Atlas Screensaver - A collection of aesthetic terminal screensavers.")
		fmt.Println("\nUsage:")
		fmt.Println("  atlas.screensaver        Start a random screensaver")
		fmt.Println("  atlas.screensaver [name] Start a specific screensaver (matrix, stars, pipes, bouncing)")
		fmt.Println("  atlas.screensaver -v     Show version")
		fmt.Println("  atlas.screensaver -h     Show this help")
		return
	}

	p := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
