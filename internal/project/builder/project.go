package builder

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/rafaeldepontes/gini/internal/log"
	"github.com/rafaeldepontes/gini/internal/project/builder/templates"
	"github.com/spf13/cobra"
)

var LongDescription string = `Build the project and put it into a new directory, if finished earlier it will delete every single change so far. Otherwise, it will create the docker-compose and Dockerfile if wanted and the "go.mod" file`

const (
	Version              = "1.0.18"
	Name                 = "Gini"
	DefaultFileMode      = 0644
	DefaultDirectoryMode = 0755
)

type docker struct {
	volumes   map[string]struct{}
	dependsOn map[string]struct{}
}

// This should be an interface maybe... But i'm not willing to make this change
// so this will be a struct holding some of the rc.log.c and then I will update it
// as needed.
type RootCmd struct {
	projectName string
	cmd         *cobra.Command
	Log         log.Logger
	docker      *docker
}

func newDocker() *docker {
	return &docker{
		volumes:   make(map[string]struct{}),
		dependsOn: make(map[string]struct{}),
	}
}

// NewRootCmd inits a new rootcmd.
func NewRootCmd() *RootCmd {
	rc := &RootCmd{
		Log:    log.NewLogger(),
		docker: newDocker(),
	}

	cmd := &cobra.Command{
		Use:           "gini",
		Version:       Version,
		Short:         "Initialize Go projects",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.AddCommand(rc.BuildProject())
	cmd.AddCommand(rc.Update())
	cmd.AddCommand(rc.Version())
	rc.Log.Info()

	rc.cmd = cmd
	return rc
}

// Execute uses the args (os.Args[1:] by default) and run through the command tree finding appropriate matches for commands and then corresponding flags. (I got this description from cobra function...)
func (rc RootCmd) Execute() error {
	return rc.cmd.Execute()
}

// ExecuteContext is the same as Execute(), but sets the ctx on the command. Retrieve ctx by calling cmd.Context() inside your *Run lifecycle or ValidArgs functions. (I got this description from cobra function...)
func (rc RootCmd) ExecuteContext(ctx context.Context) error {
	return rc.cmd.ExecuteContext(ctx)
}

func (rc RootCmd) SetContext(ctx context.Context) {
	rc.cmd.SetContext(ctx)
}

// RevertChanges delete the project's directory
func (rc RootCmd) RevertChanges() error {
	if rc.projectName != "" {
		return os.RemoveAll(fmt.Sprintf("./%s", rc.projectName))
	}
	return nil
}

// BuildProject initialize the workflow to build the project body.
//
// When called it will make some questions to the user and should build the whole
// project from it, so it's basically a bunch of edge cases...
func (rc *RootCmd) BuildProject() *cobra.Command {
	return &cobra.Command{
		Use:     "build",
		Aliases: []string{"b"},
		Short:   "Build the project based on some questions",
		Long:    LongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			rc.Log.PrintBanner(templates.ProjectBanner)
			ctx := cmd.Context()

			rc.Log.Info("Project name: ")
			projectName, err := scanLine(ctx)
			if err != nil {
				return err
			}
			rc.projectName = projectName

			if err := createGoMod(ctx, *rc); err != nil {
				return err
			}

			if err := dockerFlow(ctx, *rc); err != nil {
				return err
			}

			if err := nixFlow(ctx, *rc); err != nil {
				return err
			}

			if err := createGitEnv(ctx, *rc); err != nil {
				return err
			}
			return nil
		},
	}
}

// Update runs "go install github.com/rafaeldepontes/gini@latest"
//
// If any update is available it will be installed.
func (rc RootCmd) Update() *cobra.Command {
	return &cobra.Command{
		Use:     "update",
		Aliases: []string{"u"},
		Short:   "Update Gini for the latest version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Looking for updates...")
			ctx := cmd.Context()

			module := "github.com/rafaeldepontes/gini"

			installCmd := exec.CommandContext(ctx, "go", "install", module+"@latest")

			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr

			if err := installCmd.Run(); err != nil {
				return fmt.Errorf("go install failed: %w", err)
			}

			fmt.Println("Updated successfully!")
			return nil
		},
	}
}
