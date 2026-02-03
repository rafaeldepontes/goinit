package builder

import (
	"bufio"
	_ "embed"
	"os"
	"path"
	"text/template"

	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

const (
	// Files
	GoModFile = "go.mod"
)

//go:embed templates/go.mod.tmpl
var goModTemplate string

func createGoMod(rc *RootCmd) error {
	name := path.Join(rc.projectName, GoModFile)

	if path.Base(rc.projectName) == rc.projectName {
		if err := os.Mkdir(rc.projectName, DefaultDirectoryMode); err != nil {
			return err
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	var gitUsername string
	rc.Log.Info("Git username: ")
	if scanner.Scan() {
		gitUsername = scanner.Text()
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	goModT, err := template.New(name).Parse(goModTemplate)
	if err != nil {
		return err
	}

	templateData := make(map[string]string)
	templateData["ProjectName"] = rc.projectName
	templateData["Username"] = gitUsername

	if err := goModT.Execute(f, templateData); err != nil {
		return err
	}

	return nil
}

func addGolangCompose(rc *RootCmd) error {
	f, err := os.OpenFile(
		path.Join(rc.projectName, DockerCompose),
		os.O_RDWR|os.O_APPEND,
		OwnerPropertyMode,
	)
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
	for name := range rc.docker.dependsOn {
		if _, err := f.Write(templates.DependsOn[name]); err != nil {
			return err
		}
	}

	// TODO: Create the same logic as above for the Network...

	_, err = f.Write(templates.GolangProjectComposeSecondHalf)
	if err != nil {
		return err
	}
	return nil
}
