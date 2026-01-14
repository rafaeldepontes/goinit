package templates

import (
	brokerEnum "github.com/rafaeldepontes/goinit/internal/project/builder/enums/broker"
	databaseEnum "github.com/rafaeldepontes/goinit/internal/project/builder/enums/database"
)

var DependsOnTemplate = []byte(
	"    depends-on:\n",
)

var DependsOn map[string][]byte = map[string][]byte{
	databaseEnum.Postgres:  []byte("      - postgres:\n"),
	databaseEnum.MySql:     []byte("      - mysql:\n"),
	databaseEnum.SqlServer: []byte("      - sqlserver:\n"),
	databaseEnum.Mongo:     []byte("      - mongo:\n"),
	brokerEnum.RabbitMQ:    []byte("      - rabbitmq:\n"),
	brokerEnum.Kafka:       []byte("      - kafka:\n"),
}
