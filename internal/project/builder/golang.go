package builder

import (
	"context"
	_ "embed"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/rafaeldepontes/gini/internal/project/builder/templates"
)

const (
	// Files
	GoModFile = "go.mod"
)

//go:embed templates/go.mod.tmpl
var goModTemplate string

// createGoMod literally just creates the 'go.mod'
func createGoMod(ctx context.Context, rc RootCmd) error {
	name := path.Join(rc.projectName, GoModFile)

	if !validatePath(rc.projectName) {
		// If the directory already exists, the system will delete the old dir...
		// to prevent this, we are ignoring the error... But I still want to find
		// a way to handle this... For now, this approach works.
		_ = os.Mkdir(rc.projectName, DefaultDirectoryMode)
	}

	rc.Log.Info("Git username: ")
	gitUsername, err := scanLine(ctx)
	if err != nil {
		return err
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
	if validatePath(rc.projectName) {
		dir, _ := os.Getwd()
		templateData["ProjectName"] = filepath.Base(dir)
	}

	templateData["Username"] = gitUsername

	if err := goModT.Execute(f, templateData); err != nil {
		return err
	}

	return nil
}

// addGolangCompose writes into the docker compose the Golang service if the user
// has one
func addGolangCompose(rc RootCmd) error {
	f, err := os.OpenFile(
		path.Join(rc.projectName, DockerCompose),
		os.O_RDWR|os.O_APPEND,
		DefaultFileMode,
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

	_, err = f.Write(templates.GolangProjectComposeSecondHalf)
	if err != nil {
		return err
	}
	return nil
}

func validatePath(src string) bool {
	return src == "" || src == "." || src == "/" || src == "./"
}
