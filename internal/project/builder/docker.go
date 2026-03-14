package builder

import (
	"context"
	"os"
	"path"

	"github.com/rafaeldepontes/gini/internal/log"
	"github.com/rafaeldepontes/gini/internal/project/builder/templates"
)

const (
	DockerCompose = "docker-compose.yml"
	DockerFile    = "Dockerfile"
)

func dockerFlow(ctx context.Context, rc *RootCmd) error {
	want, err := hasDocker(ctx, rc.Log)
	if err != nil {
		return err
	}

	if want {
		// Manages part of the docker logic
		if err := createDocker(rc.projectName); err != nil {
			return err
		}

		// Manages brokers
		if err := messageBrokerFlow(ctx, rc); err != nil {
			return err
		}

		// Manages databases
		if err := databaseFlow(ctx, rc); err != nil {
			return err
		}

		// TODO: Add the ToolStack here for anyone that wants to use AWS.

		if err := addGolangCompose(rc); err != nil {
			return err
		}

		if err := createVolumes(rc); err != nil {
			return err
		}
	}

	return nil
}

// hasDocker handles the logic behind the docker-compose and the dockerfile, it appears only once at the start.
func hasDocker(ctx context.Context, log *log.Logger) (bool, error) {
	want, err := askUser(ctx, log, " Are you going to use Docker? (y/n) ")
	if err != nil {
		return false, err
	}

	return want, nil
}

// createDocker creates the DockerFile and the docker-compose for the user, initally they are empty
// but, as the user answers the questions, they should be populated.
func createDocker(name string) error {
	pathNameDF := path.Join(name, DockerFile)

	if err := os.WriteFile(
		pathNameDF,
		templates.DockerFile,
		OwnerPropertyMode,
	); err != nil {
		return err
	}

	pathNameDC := path.Join(name, DockerCompose)

	return os.WriteFile(
		pathNameDC,
		[]byte("services:\n"),
		OwnerPropertyMode,
	)
}
