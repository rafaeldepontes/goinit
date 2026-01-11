package builder

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/rafaeldepontes/goinit/internal/log"
	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
	"github.com/spf13/cobra"
)

var LongDescription string = `Build the project and put it into a new directory, if finished earlier it will delete every single change so far. Otherwise, it will create the docker-compose and Dockerfile if wanted and the "go.mod" file`

const (
	OwnerPropertyMode = 0644
)

// This should be an interface maybe... But i'm not willing to make this change
// so this will be a struct holding some of the rc.log.c and then I will update it
// as needed.
type RootCmd struct {
	projectName string
	cmd         *cobra.Command
	Log         *log.Logger
}

// NewRootCmd inits a new rootcmd.
func NewRootCmd() *RootCmd {
	rc := &RootCmd{
		Log: log.NewLogger(),
	}

	cmd := &cobra.Command{
		Use:   "gini",
		Short: "Initialize Go projects",
	}
	cmd.AddCommand(rc.BuildProject())

	rc.cmd = cmd
	return rc
}

// Execute uses the args (os.Args[1:] by default) and run through the command tree finding appropriate matches for commands and then corresponding flags. (I got this description from cobra function...)
func (rc *RootCmd) Execute() error {
	rc.Log.PrintBanner(templates.ProjectBanner)
	return rc.cmd.Execute()
}

// ExecuteContext is the same as Execute(), but sets the ctx on the command. Retrieve ctx by calling cmd.Context() inside your *Run lifecycle or ValidArgs functions. (I got this description from cobra function...)
func (rc *RootCmd) ExecuteContext(ctx context.Context) error {
	return rc.cmd.ExecuteContext(ctx)
}

func (rc *RootCmd) SetContext(ctx context.Context) {
	rc.cmd.SetContext(ctx)
}

// RevertChanges delete the project's directory
func (rc *RootCmd) RevertChanges() error {
	return os.RemoveAll(fmt.Sprintf("./%s", rc.projectName))
}

// BuildProject initialize the workflow to build the project body.
//
// When called it will make some questions to the user and should build the whole project from it, so it's basically a bunch of
// of edge cases...
func (rc *RootCmd) BuildProject() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Build the project based on some questions",
		Long:  LongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			scanner := bufio.NewScanner(os.Stdin)
			rc.Log.Info("Project name: ")

			var projectName string
			if scanner.Scan() {
				projectName = scanner.Text()
			}
			rc.projectName = projectName

			if err := os.Mkdir(rc.projectName, OwnerPropertyMode); err != nil {
				return err
			}

			if err := createGoMod(rc.projectName, rc.Log); err != nil {
				return err
			}

			if hasDocker(rc.Log) {
				// Manages part of the docker logic
				if err := createDocker(projectName); err != nil {
					return err
				}

				// Manages brokers
				if err := messageBrokerFlow(rc.projectName, rc.Log); err != nil {
					return err
				}

				// Manages databases
				if err := databaseFlow(rc.projectName, rc.Log); err != nil {
					return err
				}
			}
			return nil
		},
	}
}
