package builder

import (
	"bufio"
	"os"
	"strings"

	"github.com/rafaeldepontes/goinit/internal/log"
	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

var brokerOptions = map[int]string{
	1: "RabbitMQ",
	2: "Kafka",
}

const (
	RabbitMQ = "1"
	Kafka    = "2"
)

func messageBrokerFlow(name string, log *log.Logger) error {
	scanner := bufio.NewScanner(os.Stdin)
	if hasMessageBroker(scanner, log) {
		choice, err := askMessageBroker(log)
		if err != nil {
			return err
		}

		switch choice {
		case RabbitMQ:
			if err := createCompose(name, templates.RabbitMQCompose); err != nil {
				return err
			}

		case Kafka:
			if err := createCompose(name, templates.KafkaCompose); err != nil {
				return err
			}

		default:
			log.Warningln("As none was selected, using RabbitMQ as the default...")
			if err := createCompose(name, templates.RabbitMQCompose); err != nil {
				return err
			}
		}
	}
	return nil
}

func askMessageBroker(log *log.Logger) (string, error) {
	for i := 1; i <= len(brokerOptions); i++ {
		log.InfoPrefixf(">>>>", " [%d] %s\n", i, brokerOptions[i])
	}

	// Move cursor UP so we can put the prompt above the options.
	// Print len(options) lines, +1 to place the prompt above the first option.
	log.Infof("\033[%dA", len(brokerOptions)+1)

	// Clear the line and print prompt where the cursor currently is
	// \033[K clears from cursor to end-of-line.
	log.InfoPrefix(">>>>", " \033[KSelect the message broker: ")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	choice := strings.TrimSpace(text)

	// Move cursor DOWN to below the options so subsequent prints don't overwrite them.
	// Move down len(options) lines to end up after the list.
	log.Infof("\033[%dB", len(brokerOptions))
	log.Infoln("")

	return choice, nil
}

// hasDatabase checks to see if the user want or not a database in their docker-compose.
func hasMessageBroker(scanner *bufio.Scanner, log *log.Logger) bool {
	log.InfoPrefix(">>", " Do you want a message broker on your docker-compose? (y/n) ")

	ans := "n"
	if scanner.Scan() {
		ans = strings.ToLower(strings.TrimSpace(scanner.Text()))
	}
	return ans == "y"
}
