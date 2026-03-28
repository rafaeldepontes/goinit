package builder

import (
	"context"
	_ "embed"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/rafaeldepontes/gini/internal/log"
)

//go:embed templates/flake.nix.tmpl
var nixFlakeTemplate string

//go:embed templates/compat.nix.tmpl
var NixCompatTemplate string

func nixFlow(ctx context.Context, rc RootCmd) error {
	want, err := hasNix(ctx, rc.Log)
	if err != nil {
		return err
	}

	if want {
		wantsNixCompatFiles, err := hasNixCompatFiles(ctx, rc.Log)
		if err != nil {
			return err
		}

		if err := createNixFiles(rc, wantsNixCompatFiles); err != nil {
			return err
		}
		if err := createDerivationGitignore(rc); err != nil {
			return err
		}
	}
	return nil
}

func askUser(ctx context.Context, log log.Logger, question string) (bool, error) {
	log.InfoPrefix(">>", question)

	ans, err := scanLine(ctx)
	if err != nil {
		return false, err
	}

	ans = strings.ToLower(strings.TrimSpace(ans))
	return ans == "y", nil
}

func hasNix(ctx context.Context, log log.Logger) (bool, error) {
	want, err := askUser(ctx, log, " Are you going to use Nix? (y/n) ")
	if err != nil {
		return false, err
	}

	return want, nil
}

func hasNixCompatFiles(ctx context.Context, log log.Logger) (bool, error) {
	want, err := askUser(ctx, log, " Do you want to create compatibility files for versions of nix that don't support flakes? (y/n) ")
	if err != nil {
		return false, err
	}

	return want, nil
}

func createNixFiles(rc RootCmd, flakeCompat bool) error {
	flakeT, err := template.New("flake.nix").Parse(nixFlakeTemplate)
	if err != nil {
		return err
	}

	compatT, err := template.New("compat.nix").Parse(NixCompatTemplate)
	if err != nil {
		return err
	}

	filesToBeCreated := []string{
		"flake.nix",
		"default.nix",
		"shell.nix",
	}

	if !flakeCompat {
		filesToBeCreated = filesToBeCreated[:1]
	}

	files := make([]*os.File, 0, len(filesToBeCreated))
	for _, v := range filesToBeCreated {
		f, err := os.Create(
			path.Join(rc.projectName, v),
		)
		if err != nil {
			return err
		} else {
			files = append(files, f)
		}
		defer f.Close()
	}

	templateData := map[string]string{
		"ProjectName": rc.projectName,
		"Function":    "defaultNix",
	}

	if flakeCompat {
		templateData["Compat"] = "    flake-compat = {\n" +
			"      url = \"github:NixOS/flake-compat\";\n" +
			"      flake = false;\n" +
			"    };"
	} else {
		templateData["Compat"] = ""
	}

	flakeF := files[0]

	if err := flakeT.Execute(flakeF, templateData); err != nil {
		return err
	}

	if flakeCompat {
		defaultNixF := files[1]
		shellNixF := files[2]

		if err := compatT.Execute(defaultNixF, templateData); err != nil {
			return err
		}

		templateData["Function"] = "shellNix"

		if err := compatT.Execute(shellNixF, templateData); err != nil {
			return err
		}
	}

	return nil
}

func createDerivationGitignore(rc RootCmd) error {
	filename := path.Join(rc.projectName, ".gitignore")
	textToAppend := "/out\n\n"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, DefaultFileMode)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(textToAppend); err != nil {
		return err
	}
	return nil
}
