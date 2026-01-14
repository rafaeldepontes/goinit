package builder

import (
	"bufio"
	_ "embed"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/rafaeldepontes/goinit/internal/log"
)

//go:embed templates/flake.nix.tmpl
var nixFlakeTemplate string

//go:embed templates/compat.nix.tmpl
var NixCompatTemplate string

func askUser(log *log.Logger, question string) bool {
	scanner := bufio.NewScanner(os.Stdin)
	log.InfoPrefix(">>", question)

	ans := "n"
	if scanner.Scan() {
		ans = strings.ToLower(strings.TrimSpace(scanner.Text()))
	}
	return ans == "y"
}

func hasNix(log *log.Logger) bool {
	return askUser(log, " Are you going to use Nix? (y/n) ")
}

func hasNixCompatFiles(log *log.Logger) bool {
	return askUser(log, " Do you want to create compatibility files for versions of nix that don't support flakes? (y/n) ")
}

func createNixFiles(rc *RootCmd, flakeCompat bool) error {
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

	files := []*os.File{}
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

func createDerivationGitignore(rc *RootCmd) error {
	filename := path.Join(rc.projectName, ".gitignore")
	textToAppend := "/out\n"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.WriteString(textToAppend); err != nil {
		return err
	}
	return nil
}
