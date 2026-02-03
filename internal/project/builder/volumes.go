package builder

import (
	"os"
	"path"

	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

func createVolumes(rc *RootCmd) error {
	f, err := os.OpenFile(
		path.Join(rc.projectName, DockerCompose),
		os.O_RDWR|os.O_APPEND,
		OwnerPropertyMode,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(templates.VolumeTemplate); err != nil {
		return err
	}

	for name := range rc.docker.volumes {
		if _, err := f.Write(templates.Volumes[name]); err != nil {
			return err
		}
	}
	return nil
}
