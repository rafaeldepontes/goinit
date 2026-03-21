package builder

import (
	"github.com/spf13/cobra"
)

// Version represents the version command
func (rc *RootCmd) Version() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show the version of your Gini",
		Run: func(cmd *cobra.Command, args []string) {
			rc.Log.Infof("Gini v%s\n", Version)
		},
	}
}
