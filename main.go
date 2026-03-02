package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/rafaeldepontes/gini/internal/project/builder"
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
		rootCmd.Log.Errorln("\n[ERROR] command failed:", err)
		revert(rootCmd)
		os.Exit(1)
	}
}

func revert(cmd *builder.RootCmd) {
	if err := cmd.RevertChanges(); err != nil {
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			cmd.Log.Errorln("[ERROR] cannot delete root directory...")
			return
		}
		cmd.Log.Errorln("[ERROR] couldn't revert changes:", err)
	}
}
