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
	scanner := bufio.NewScanner(os.Stdin)
	if hasDatabase(scanner, log) {
		log.InfoPrefixln(">>>>", " Select the database: ")
		for i := 0; i < len(databaseOptions); i++ {
			log.InfoPrefixf(">>>>", " [%d] %s\n", i+1, databaseOptions[i+1])
		}

		if scanner.Scan() {
			switch strings.TrimSpace(scanner.Text()) {
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
func hasDatabase(scanner *bufio.Scanner, log *log.Logger) bool {
	log.InfoPrefix(">>", " Do you want a database on your docker-compose? (y/n) ")

	ans := "n"
	if scanner.Scan() {
		ans = strings.ToLower(strings.TrimSpace(scanner.Text()))
	}
	return ans == "y"
}
