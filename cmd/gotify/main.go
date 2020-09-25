package main

import (
	"fmt"
	"os"

	"github.com/cthit/gotify/internal/app"
)

func main() {
	err := app.Start()
	if err != nil {
		fmt.Printf("Crash: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
