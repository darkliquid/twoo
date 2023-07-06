package main

import (
	"os"

	"github.com/darkliquid/twoo/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
