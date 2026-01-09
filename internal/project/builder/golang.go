package builder

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rafaeldepontes/goinit/internal/log"
	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

const (
	// Files
	GoModFile = "go.mod"
)

func createGoMod(log *log.Logger) (string, error) {
	template := templates.GodModDefault
	scanner := bufio.NewScanner(os.Stdin)

	var gitUsername string
	log.Info("Git username: ")
	if scanner.Scan() {
		gitUsername = scanner.Text()
	}

	var projectName string
	if gitUsername != "" {
		log.Info("Project name: ")
		if scanner.Scan() {
			projectName = scanner.Text()
			template = fmt.Sprintf(templates.GoMod, gitUsername, projectName)
		}
	}

	if err := os.WriteFile(
		GoModFile,
		[]byte(template),
		OwnerPropertyMode,
	); err != nil {
		return "", err
	}
	return projectName, nil
}
