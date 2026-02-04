package builder

import (
	_ "embed"
	"os"
	"path"
	"text/template"
)

//go:embed templates/readme.md.tmpl
var readmeTemplate string

//go:embed templates/.gitignore.tmpl
var gitignoreTemplate string

const (
	ReadMeName    = "README.md"
	GitIgnoreName = ".gitignore"
)

func createGitEnv(rc *RootCmd) error {
	readmePath := path.Join(rc.projectName, ReadMeName)
	f, err := os.Create(readmePath)
	if err != nil {
		return err
	}
	defer f.Close()

	readT, err := template.New(readmePath).Parse(readmeTemplate)
	if err != nil {
		return err
	}

	templateData := map[string]string{
		"Name": rc.projectName,
	}

	if err = readT.Execute(f, templateData); err != nil {
		return err
	}

	gitPath := path.Join(rc.projectName, GitIgnoreName)
	gitF, err := os.OpenFile(
		gitPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		OwnerPropertyMode,
	)
	if err != nil {
		return err
	}
	defer gitF.Close()

	gitT, err := template.New(gitPath).Parse(gitignoreTemplate)
	if err != nil {
		return err
	}

	if err = gitT.Execute(gitF, nil); err != nil {
		return err
	}

	if _, err = gitF.WriteString("\n"); err != nil {
		return err
	}
	return nil
}
