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

func createGoMod(name string, log *log.Logger) error {
	if err := os.Mkdir(name, OwnerPropertyMode); err != nil {
		return err
	}

	template := templates.GodModDefault
	scanner := bufio.NewScanner(os.Stdin)

	var gitUsername string
	log.Info("Git username: ")
	if scanner.Scan() {
		gitUsername = scanner.Text()
	}

	if gitUsername != "" {
		template = fmt.Sprintf(templates.GoMod, gitUsername, name)
	}

	if err := os.WriteFile(
		GoModFile,
		[]byte(template),
		OwnerPropertyMode,
	); err != nil {
		return err
	}

	return os.Rename(GoModFile, fmt.Sprintf("%s/%s", name, GoModFile))
}

func addGolangCompose(rc *RootCmd) error {
	name := fmt.Sprintf("./%s/%s", rc.projectName, DockerCompose)
	f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, OwnerPropertyMode)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(templates.GolangProjectComposeFirstHalf)
	if err != nil {
		return err
	}

	_, err = f.Write(templates.DependsOnTemplate)
	if err != nil {
		return err
	}

	// Depends-on logic
	for name, have := range rc.docker.dependsOn {
		if !have {
			continue
		}

		if _, err := f.Write(templates.DependsOn[name]); err != nil {
			return err
		}
	}

	_, err = f.Write(templates.GolangProjectComposeSecondHalf)
	if err != nil {
		return err
	}
	return nil
}
