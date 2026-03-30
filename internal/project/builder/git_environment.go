package builder

import (
	"context"
	_ "embed"
	"os"
	"os/exec"
	"path"
	"path/filepath"
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

// createGitEnv holds the logic for any git related feature, like for instance
// the .git dir, .gitignore file and the README as well.
func createGitEnv(ctx context.Context, rc RootCmd) error {
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

	if isRoot(rc.projectName) {
		dir, _ := os.Getwd()
		templateData["Name"] = filepath.Base(dir)
	}

	if err = readT.Execute(f, templateData); err != nil {
		return err
	}

	gitPath := path.Join(rc.projectName, GitIgnoreName)
	gitF, err := os.OpenFile(
		gitPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		DefaultFileMode,
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

	if err = gitInit(ctx, rc); err != nil {
		return err
	}

	return nil
}

// gitInit initalizes the .git project by running the 'git init -b main' on the
// project dir, if the user doesn't have git installed... A warning message will
// be displayed and the application will ignore this part of the logic.
func gitInit(ctx context.Context, rc RootCmd) error {
	_, err := exec.LookPath("git")
	if err != nil {
		rc.Log.Warningln("[WARN] user doesn't have git installed")
		return nil
	}

	cmd := exec.CommandContext(ctx, "git", "init", "-b", "main")

	if !isRoot(rc.projectName) {
		cmd.Dir = rc.projectName
	}

	return cmd.Run()
}
