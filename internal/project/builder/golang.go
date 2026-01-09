package builder

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

const (
	GoModFile = "go.mod"
)

func createGoMod() (string, error) {
	template := templates.GodModDefault
	scanner := bufio.NewScanner(os.Stdin)

	var gitUsername string
	fmt.Print("Git username: ")
	if scanner.Scan() {
		gitUsername = scanner.Text()
	}

	var projectName string
	if gitUsername != "" {
		fmt.Print("Project name: ")
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
