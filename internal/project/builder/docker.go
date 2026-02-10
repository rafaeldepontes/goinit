package builder

import (
	"bufio"
	"os"
	"path"
	"strings"

	"github.com/rafaeldepontes/goinit/internal/log"
	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

const (
	DockerCompose = "docker-compose.yml"
	DockerFile    = "Dockerfile"
)

// hasDocker handles the logic behind the docker-compose and the dockerfile, it appears only once at the start.
func hasDocker(log *log.Logger) bool {
	scanner := bufio.NewScanner(os.Stdin)
	log.InfoPrefix(">>", " Are you going to use Docker? (y/n) ")

	ans := "n"
	if scanner.Scan() {
		ans = strings.ToLower(strings.TrimSpace(scanner.Text()))
	}
	return ans == "y"
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
