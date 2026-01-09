package builder

import (
	"context"
	"log"
	"os"
	"path/filepath"

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

// BuildProject initialize the workflow to build the project body.
//
// When called it will make some questions to the user and should build the whole project from it, so it's basically a bunch of
// of edge cases...
func (rc *RootCmd) BuildProject() *cobra.Command {
	return &cobra.Command{
		Use:   "gini build",
		Short: "Build the project based on some questions",
		Run: func(cmd *cobra.Command, args []string) {
			// Create the go mod
			projectName, err := createGoMod()
			if err != nil {
				log.Fatalln("[ERROR] didn't create the go.mod: ", err)
				return
			}

			if err := createDir(projectName); err != nil {
				log.Fatalln("[ERROR] didn't create the dir: ", err)
				return
			}

			if hasDocker() {
				// Manages part of the docker logic
				if err := createDocker(projectName); err != nil {
					log.Fatalln("[ERROR] didn't create the docker-compose/Dockerfile: ", err)
					return
				}

				// Manages brokers
				if err := messageBrokerFlow(projectName); err != nil {
					log.Fatalln("[ERROR] didn't create the message broker: ", err)
					return
				}

				// Manages databases
				if err := databaseFlow(projectName); err != nil {
					log.Fatalln("[ERROR] didn't create the database: ", err)
					return
				}
			}
		},
	}
}

func createDir(name string) error {
	if err := os.Mkdir(name, OwnerPropertyMode); err != nil {
		return err
	}
	return os.Rename("go.mod", filepath.Join(name, "go.mod"))
}
