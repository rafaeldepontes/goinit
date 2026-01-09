package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rafaeldepontes/goinit/internal/project/builder"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	rootCmd := builder.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		rootCmd.Log.Errorln("\n[ERROR] executing command prompt: " + err.Error())
		return
	}

	<-sigChan
	rootCmd.Log.Infoln("\nBuild finished early, excluding the changes...")

	if err := rootCmd.RevertChanges(); err != nil {
		rootCmd.Log.Errorln("\n[ERROR] couldn't revert all changes: " + err.Error())
		return
	}
}
