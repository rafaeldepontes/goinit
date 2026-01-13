package builder

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rafaeldepontes/goinit/internal/log"
	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

var databaseOptions = map[int]string{
	1: "PostgreSQL",
	2: "MySQL",
	3: "SQL Server",
	4: "MongoDB",
}

const (
	Postgres  = "1"
	MySql     = "2"
	SqlServer = "3"
	Mongo     = "4"
)

// DockerFlow handles the logic behind the docker-compose and the dockerfile, it appears only once at the start.
func databaseFlow(name string, log *log.Logger) error {
	if hasDatabase(log) {
		selecteds, err := askDatabase(log)
		if err != nil {
			return err
		}

		log.Infoln("Selecteds, " + selecteds)

		choices := strings.Split(selecteds, ",")
		for _, choice := range choices {
			switch strings.TrimSpace(choice) {
			case Postgres:
				if err := createCompose(name, templates.PostgresCompose); err != nil {
					return err
				}

			case MySql:
				if err := createCompose(name, templates.MySQLCompose); err != nil {
					return err
				}

			case SqlServer:
				if err := createCompose(name, templates.SQLServerCompose); err != nil {
					return err
				}

			case Mongo:
				if err := createCompose(name, templates.MongoCompose); err != nil {
					return err
				}

			default:
				log.Warningln("As none was selected, using PostgreSQL as the default...")
				if err := createCompose(name, templates.PostgresCompose); err != nil {
					return err
				}

			}
		}
	}
	return nil
}

func askDatabase(log *log.Logger) (string, error) {
	for i := 1; i <= len(databaseOptions); i++ {
		log.InfoPrefixf(">>>>", " [%d] %s\n", i, databaseOptions[i])
	}

	// Move cursor UP so we can put the prompt above the options.
	// Print len(options) lines, +1 to place the prompt above the first option.
	log.Infof("\033[%dA", len(databaseOptions)+1)

	// Clear the line and print prompt where the cursor currently is
	// \033[K clears from cursor to end-of-line.
	log.InfoPrefix(">>>>", " \033[KSelect the database: ")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	choices := strings.TrimSpace(text)

	// Move cursor DOWN to below the options so subsequent prints don't overwrite them.
	// Move down len(options) lines to end up after the list.
	log.Infof("\033[%dB", len(databaseOptions))
	log.Infoln("")

	return choices, nil
}

func createCompose(pn string, db []byte) error {
	name := fmt.Sprintf("./%s/%s", pn, DockerCompose)
	f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, OwnerPropertyMode)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(db)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
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
