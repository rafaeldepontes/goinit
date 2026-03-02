package builder

import (
	"context"
	"os"
	"path"
	"strings"

	"github.com/rafaeldepontes/gini/internal/log"
	"github.com/rafaeldepontes/gini/internal/project/builder/templates"
)

const (
	DockerCompose = "docker-compose.yml"
	DockerFile    = "Dockerfile"
)

// hasDocker handles the logic behind the docker-compose and the dockerfile, it appears only once at the start.
func hasDocker(ctx context.Context, log *log.Logger) (bool, error) {
	log.InfoPrefix(">>", " Are you going to use Docker? (y/n) ")

	ans, err := scanLine(ctx)
	if err != nil {
		return false, err
	}

	ans = strings.ToLower(strings.TrimSpace(ans))
	return ans == "y", nil
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
