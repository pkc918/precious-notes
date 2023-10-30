package main

import (
	"os"
	"precious-notes/internal/preciousnotes"
)

func main() {
	command := preciousnotes.NewPreciousNotesCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
