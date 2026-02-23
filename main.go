package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rafaeldepontes/goinit/internal/project/builder"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	rootCmd := builder.NewRootCmd()
	rootCmd.SetContext(ctx)

	if err := rootCmd.Execute(); err != nil {
		rootCmd.Log.Errorln("\n[ERROR] command failed: ", err)
		revert(rootCmd)
		os.Exit(1)
	}
}

func revert(cmd *builder.RootCmd) {
	if err := cmd.RevertChanges(); err != nil {
		cmd.Log.Errorln("\n[ERROR] couldn't revert changes: ", err)
	}
}
