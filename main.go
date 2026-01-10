package main

import (
	"context"
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

	done := make(chan error, 1)

	go func() {
		done <- rootCmd.Execute()
	}()

	select {
	case err := <-done:
		if err != nil {
			rootCmd.Log.Errorln("\n[ERROR] command failed: " + err.Error())
			revert(rootCmd)
		}
		return

	case <-ctx.Done():
		rootCmd.Log.Warningln("\nCommand interrupted — reverting changes...")
		revert(rootCmd)
		return
	}
}

func revert(cmd *builder.RootCmd) {
	if err := cmd.RevertChanges(); err != nil {
		cmd.Log.Errorln("\n[ERROR] couldn't revert changes: " + err.Error())
	}
}
