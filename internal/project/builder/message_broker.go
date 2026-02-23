package builder

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/rafaeldepontes/gini/internal/log"
	enums "github.com/rafaeldepontes/gini/internal/project/builder/enums/broker"
	dbEnum "github.com/rafaeldepontes/gini/internal/project/builder/enums/database"
	"github.com/rafaeldepontes/gini/internal/project/builder/templates"
)

var brokerOptions = map[int]string{
	1: "RabbitMQ",
	2: "Kafka",
}

func messageBrokerFlow(rc *RootCmd) error {
	if hasMessageBroker(rc.Log) {
		choices, err := askMessageBroker(rc.Log)
		if err != nil {
			return err
		}

		for _, choice := range choices {
			choice = strings.TrimSpace(choice)

			if choice != "" {
				val, _ := strconv.Atoi(choice)
				lastVal, _ := strconv.Atoi(dbEnum.Redis)

				choice = strconv.Itoa(val + lastVal)
			}

			switch choice {
			case enums.RabbitMQ:
				if err := createGenericCompose(rc, templates.RabbitMQCompose, enums.RabbitMQ); err != nil {
					return err
				}

			case enums.Kafka:
				if err := createGenericCompose(rc, templates.KafkaCompose, enums.Kafka); err != nil {
					return err
				}

			default:
				rc.Log.Warningln("Invalid input, using RabbitMQ as the default...")
				if err := createGenericCompose(rc, templates.RabbitMQCompose, enums.RabbitMQ); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func askMessageBroker(log *log.Logger) ([]string, error) {
	for i := 1; i <= len(brokerOptions); i++ {
		log.InfoPrefixf(">>>>", " [%d] %s\n", i, brokerOptions[i])
	}

	// Move cursor UP so we can put the prompt above the options.
	// Print len(options) lines, +1 to place the prompt above the first option.
	log.Infof("\033[%dA", len(brokerOptions)+1)

	// Clear the line and print prompt where the cursor currently is
	// \033[K clears from cursor to end-of-line.
	log.InfoPrefix(">>>>", " \033[KSelect one or more message broker (use commas, e.g. 1,3): ")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	choices := strings.Split(strings.TrimSpace(text), ",")

	// Move cursor DOWN to below the options so subsequent prints don't overwrite them.
	// Move down len(options) lines to end up after the list.
	log.Infof("\033[%dB", len(brokerOptions))
	log.Infoln("")

	return choices, nil
}

// hasDatabase checks to see if the user want or not a database in their docker-compose.
func hasMessageBroker(log *log.Logger) bool {
	scanner := bufio.NewScanner(os.Stdin)
	log.InfoPrefix(">>", " Do you want a message broker on your docker-compose? (y/n) ")

	ans := "n"
	if scanner.Scan() {
		ans = strings.ToLower(strings.TrimSpace(scanner.Text()))
	}
	return ans == "y"
}
