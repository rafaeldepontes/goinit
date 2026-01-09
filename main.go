package main

import (
	"log"

	"github.com/rafaeldepontes/goinit/internal/project/builder"
)

func main() {
	rootCmd := builder.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("[ERROR] ", err)
		return
	}
}
