package main

import (
	"github.com/rafaeldepontes/goinit/internal/project/builder"
)

func main() {
	rootCmd := builder.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		rootCmd.Log.Errorln("[ERROR] " + err.Error())
		return
	}
}
