package main

import (
	"fmt"
	"github.com/tuxikus/env-forge/internal/envforge"
	"os"
)

func usage() {
	fmt.Println("usage: env-forge <config.e>")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Illegal argument count")
		usage()
		os.Exit(1)
	}
	eFile := os.Args[1]

	fmt.Println("Setting up environment...")

	envForge := envforge.NewEnvForge(eFile)
	envForge.Forge()

	fmt.Println("Done! 🥳")
}
