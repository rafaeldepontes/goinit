package builder

import (
	"fmt"
	"os"

	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

func createVolumes(rc *RootCmd) error {
	name := fmt.Sprintf("./%s/%s", rc.projectName, DockerCompose)
	f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, OwnerPropertyMode)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(templates.VolumeTemplate); err != nil {
		return err
	}

	for name, have := range rc.docker.volumes {
		if !have {
			continue
		}

		if _, err := f.Write(templates.Volumes[name]); err != nil {
			return err
		}
	}
	return nil
}
