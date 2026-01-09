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
		log.InfoPrefixln(">>>>", " Select the message broker: ")
		for i := 0; i < len(brokerOptions); i++ {
			log.InfoPrefixf(">>>>", " [%d] %s\n", i+1, brokerOptions[i+1])
		}

		if scanner.Scan() {
			switch strings.TrimSpace(scanner.Text()) {
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
	}
	return nil
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
