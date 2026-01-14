package templates

import (
	enums "github.com/rafaeldepontes/goinit/internal/project/builder/enums/database"
)

var VolumeTemplate = []byte(
	"volumes:\n",
)

var Volumes map[string][]byte = map[string][]byte{
	enums.Postgres:  []byte("  postgres:\n"),
	enums.MySql:     []byte("  mysql:\n"),
	enums.SqlServer: []byte("  sqlserver:\n"),
	enums.Mongo:     []byte("  mongo:\n"),
}
