package main

import (
	"fmt"
	"os"

	"github.com/claykom/lode/internal/tui"
)

func main() {
	if err := tui.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
