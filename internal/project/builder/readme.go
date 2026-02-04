package builder

import (
	_ "embed"
	"os"
	"path"
	"text/template"
)

//go:embed templates/readme.md.tmpl
var readmeTemplate string
var ReadMeName = "README.md"

// TODO: could be a good thing to create the git ignore right here, after all...
// the best case scenario possible would be getting the git ignore from the golang
// repository!
func createReadMe(rc *RootCmd) error {
	name := path.Join(rc.projectName, ReadMeName)
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	readT, err := template.New(name).Parse(readmeTemplate)
	if err != nil {
		return err
	}

	templateData := map[string]string{
		"Name": rc.projectName,
	}
	return readT.Execute(f, templateData)
}
