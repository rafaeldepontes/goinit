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

func createReadMe(rc *RootCmd) error {
	name := path.Join(rc.projectName, ReadMeName)
	f, err := os.Create(name)
	if err != nil {
		rc.Log.Error("[ERROR] Could not create the README.md:", err)
		return err
	}
	defer f.Close()

	readT, err := template.New(name).Parse(readmeTemplate)
	if err != nil {
		rc.Log.Errorln("[ERROR] Could not parse the README file:", err)
		return err
	}

	templateData := map[string]string{
		"Name": rc.projectName,
	}
	return readT.Execute(f, templateData)
}
