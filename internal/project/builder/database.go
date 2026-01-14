package builder

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rafaeldepontes/goinit/internal/log"
	enums "github.com/rafaeldepontes/goinit/internal/project/builder/enums/database"
	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

// DockerFlow handles the logic behind the docker-compose and the dockerfile, it appears only once at the start.
func databaseFlow(rc *RootCmd) error {
	if hasDatabase(rc.Log) {
		choices, err := askDatabase(rc.Log)
		if err != nil {
			return err
		}

		for _, choice := range choices {
			switch strings.TrimSpace(choice) {
			case enums.Postgres:
				if err := createCompose(rc, templates.PostgresCompose, enums.Postgres); err != nil {
					return err
				}

			case enums.MySql:
				if err := createCompose(rc, templates.MySQLCompose, enums.MySql); err != nil {
					return err
				}

			case enums.SqlServer:
				if err := createCompose(rc, templates.SQLServerCompose, enums.SqlServer); err != nil {
					return err
				}

			case enums.Mongo:
				if err := createCompose(rc, templates.MongoCompose, enums.Mongo); err != nil {
					return err
				}

			case enums.Redis:
				if err := createGenericCompose(rc, templates.RedisCompose, enums.Redis); err != nil {
					return err
				}

			default:
				rc.Log.Warningln("Invalid input, using PostgreSQL as the default...")
				if err := createCompose(rc, templates.PostgresCompose, enums.Postgres); err != nil {
					return err
				}

			}
		}
	}
	return nil
}

func askDatabase(log *log.Logger) ([]string, error) {
	var databaseOptions = map[int]string{
		1: "PostgreSQL",
		2: "MySQL",
		3: "SQL Server",
		4: "MongoDB",
		5: "Redis",
	}

	for i := 1; i <= len(databaseOptions); i++ {
		log.InfoPrefixf(">>>>", " [%d] %s\n", i, databaseOptions[i])
	}

	// Move cursor UP so we can put the prompt above the options.
	// Print len(options) lines, +1 to place the prompt above the first option.
	log.Infof("\033[%dA", len(databaseOptions)+1)

	// Clear the line and print prompt where the cursor currently is
	// \033[K clears from cursor to end-of-line.
	log.InfoPrefix(">>>>", " \033[KSelect one or more databases (use commas, e.g. 1,3): ")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	choices := strings.Split(strings.TrimSpace(text), ",")

	// Move cursor DOWN to below the options so subsequent prints don't overwrite them.
	// Move down len(options) lines to end up after the list.
	log.Infof("\033[%dB", len(databaseOptions))
	log.Infoln("")

	return choices, nil
}

func createCompose(rc *RootCmd, compose []byte, dbName string) error {
	name := fmt.Sprintf("./%s/%s", rc.projectName, DockerCompose)
	f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, OwnerPropertyMode)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(compose)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	rc.docker.volumes[dbName] = true
	rc.docker.dependsOn[dbName] = true

	return nil
}

// hasDatabase checks to see if the user want or not a database in their docker-compose.
func hasDatabase(log *log.Logger) bool {
	scanner := bufio.NewScanner(os.Stdin)
	log.InfoPrefix(">>", " Do you want a database on your docker-compose? (y/n) ")

	ans := "n"
	if scanner.Scan() {
		ans = strings.ToLower(strings.TrimSpace(scanner.Text()))
	}
	return ans == "y"
}

func createGenericCompose(rc *RootCmd, compose []byte, serviceName string) error {
	name := fmt.Sprintf("./%s/%s", rc.projectName, DockerCompose)
	f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, OwnerPropertyMode)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(compose)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	rc.docker.dependsOn[serviceName] = true

	return nil
}
