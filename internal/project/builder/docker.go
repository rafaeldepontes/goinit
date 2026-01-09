package builder

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

const (
	DockerCompose = "docker-compose.yml"
	DockerFile    = "Dockerfile"
)

// hasDocker handles the logic behind the docker-compose and the dockerfile, it appears only once at the start.
func hasDocker() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">> Are you going to use Docker? (y/n) ")

	ans := "n"
	if scanner.Scan() {
		ans = strings.ToLower(strings.TrimSpace(scanner.Text()))
	}
	return ans == "y"
}

// createDocker creates the DockerFile and the docker-compose for the user, initally they are empty
// but, as the user answers the questions, they should be populated.
func createDocker(name string) error {
	if err := os.WriteFile(
		DockerFile,
		templates.DockerFile,
		OwnerPropertyMode,
	); err != nil {
		return err
	}

	if err := os.WriteFile(
		DockerCompose,
		templates.DockerComposeTemplate,
		OwnerPropertyMode,
	); err != nil {
		return err
	}

	if err := os.Rename(DockerFile, filepath.Join(name, DockerFile)); err != nil {
		return err
	}

	return os.Rename(DockerCompose, filepath.Join(name, DockerCompose))
}
