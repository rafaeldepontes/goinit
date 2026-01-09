package builder

import (
	"context"
	"os"
	"path/filepath"

	"github.com/rafaeldepontes/goinit/internal/log"
	"github.com/spf13/cobra"
)

const (
	OwnerPropertyMode = 0644
)

// This should be an interface maybe... But i'm not willing to make this change
// so this will be a struct holding some of the rc.log.c and then I will update it
// as needed.
type RootCmd struct {
	cmd *cobra.Command
	Log *log.Logger
}

// NewRootCmd inits a new rootcmd.
func NewRootCmd() *RootCmd {
	rc := &RootCmd{
		Log: log.NewLogger(),
	}

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
			projectName, err := createGoMod(rc.Log)
			if err != nil {
				rc.Log.Errorln("[ERROR] didn't create the go.mod: " + err.Error())
				return
			}

			if err := createDir(projectName); err != nil {
				rc.Log.Errorln("[ERROR] didn't create the dir: " + err.Error())
				return
			}

			if hasDocker(rc.Log) {
				// Manages part of the docker rc.log.c
				if err := createDocker(projectName); err != nil {
					rc.Log.Errorln("[ERROR] didn't create the docker-compose/Dockerfile: " + err.Error())
					return
				}

				// Manages brokers
				if err := messageBrokerFlow(projectName, rc.Log); err != nil {
					rc.Log.Errorln("[ERROR] didn't create the message broker: " + err.Error())
					return
				}

				// Manages databases
				if err := databaseFlow(projectName, rc.Log); err != nil {
					rc.Log.Errorln("[ERROR] didn't create the database: " + err.Error())
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
