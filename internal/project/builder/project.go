package builder

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

const (
	OwnerPropertyMode = 0644
)

// This should be an interface maybe... But i'm not willing to make this change
// so this will be a struct holding some of the logic and then I will update it
// as needed.
type RootCmd struct {
	cmd *cobra.Command
}

// NewRootCmd inits a new rootcmd.
func NewRootCmd() *RootCmd {
	rc := &RootCmd{}

	cmd := &cobra.Command{}
	cmd.AddCommand(rc.BuildProject())

	rc.cmd = cmd
	return rc
}

// Execute uses the args (os.Args[1:] by default) and run through the command tree finding appropriate matches for commands and then corresponding flags. (I got this description from cobra function...)
func (rc *RootCmd) Execute() error {
	return rc.cmd.Execute()
}

// ExecuteContext is the same as Execute(), but sets the ctx on the command. Retrieve ctx by calling cmd.Context() inside your *Run lifecycle or ValidArgs functions. (I got this description from cobra function...)
func (rc *RootCmd) ExecuteContext(ctx context.Context) error {
	return rc.cmd.ExecuteContext(ctx)
}

// <NAME> initialize the workflow to build the project body.
//
// When called it will make some questions to the user and should build the whole project from it, so it's basically a bunch of
// of edge cases...
func (rc *RootCmd) BuildProject() *cobra.Command {
	return &cobra.Command{
		Use:   "gini build",
		Short: "Build the project based on some questions",
		Run: func(cmd *cobra.Command, args []string) {
			if hasDocker() {
				// Manages part of the docker logic
				if err := createDocker(); err != nil {
					log.Fatalln("[ERROR] ", err)
					return
				}

				// Manager the database
				if err := databaseFlow(); err != nil {
					log.Fatalln("[ERROR] ", err)
					return
				}

				// >> Do you want a message broker on your docker-compose? (y/n)
				// >>>> Select the message broker:
				// >>>> [1]
				// >>>> [2]

			}
			// Create the go mod
		},
	}
}
