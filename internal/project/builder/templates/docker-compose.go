package templates

var DockerComposeTemplate = []byte(
	"services:\n" +
		"  golang-project:\n" +
		"    container_name: \"my-project\"\n" +
		"    build:\n" +
		"      context: ./\n" +
		"      dockerfile: Dockerfile\n" +
		"    image: golang-project\n" +
		"    deploy:\n" +
		"      mode: replicated\n" +
		"      replicas: 1\n" +
		"    env_file:\n" +
		"      - .env\n" +
		"    ports:\n" +
		"      - \"8080:8080\"\n\n",
)

var PostgresCompose = []byte(
	"  postgres:\n" +
		"    image: postgres:latest\n" +
		"    container_name: \"postgres\"\n" +
		"    ports:\n" +
		"      - \"5432:5432\"\n" +
		"    environment:\n" +
		"      POSTGRES_DB: \"postgres\"\n" +
		"      POSTGRES_USER: \"user\"\n" +
		"      POSTGRES_PASSWORD: \"pass\"\n" +
		"    deploy:\n" +
		"      mode: replicated\n" +
		"      replicas: 1\n" +
		"    volumes:\n" +
		"      - postgres:/var/lib/postgresql/data\n\n",
)

var MySQLCompose = []byte(
	"  mysql:\n" +
		"    image: mysql:8.0\n" +
		"    container_name: \"mysql\"\n" +
		"    ports:\n" +
		"      - \"3306:3306\"\n" +
		"    environment:\n" +
		"      MYSQL_DATABASE: \"mysql\"\n" +
		"      MYSQL_USER: \"user\"\n" +
		"      MYSQL_PASSWORD: \"pass\"\n" +
		"      MYSQL_ROOT_PASSWORD: \"root\"\n" +
		"    deploy:\n" +
		"      mode: replicated\n" +
		"      replicas: 1\n" +
		"    volumes:\n" +
		"      - mysql:/var/lib/mysql\n\n",
)

var SQLServerCompose = []byte(
	"  sqlserver:\n" +
		"    image: mcr.microsoft.com/mssql/server:2022-latest\n" +
		"    container_name: \"sqlserver\"\n" +
		"    ports:\n" +
		"      - \"1433:1433\"\n" +
		"    environment:\n" +
		"      ACCEPT_EULA: \"Y\"\n" +
		"      SA_PASSWORD: \"StrongPassword!123\"\n" +
		"    deploy:\n" +
		"      mode: replicated\n" +
		"      replicas: 1\n" +
		"    volumes:\n" +
		"      - sqlserver:/var/opt/mssql\n\n",
)

var RabbitMQCompose = []byte(
	"  rabbitmq:\n" +
		"    image: \"rabbitmq:3-management\"\n" +
		"    container_name: rabbitmq\n" +
		"    environment:\n" +
		"      - RABBITMQ_DEFAULT_USER=user\n" +
		"      - RABBITMQ_DEFAULT_PASS=pass\n" +
		"    ports:\n" +
		"      - \"5672:5672\"\n" +
		"      - \"15672:15672\"\n" +
		"    deploy:\n" +
		"      mode: replicated\n" +
		"      replicas: 1\n\n",
)

var KafkaCompose = []byte(
	"  zookeeper:\n" +
		"    image: confluentinc/cp-zookeeper:7.6.1\n" +
		"    container_name: \"zookeeper\"\n" +
		"    ports:\n" +
		"      - \"2181:2181\"\n" +
		"    deploy:\n" +
		"      mode: replicated\n" +
		"      replicas: 1\n" +
		"    environment:\n" +
		"      ZOOKEEPER_CLIENT_PORT: 2181\n" +
		"      ZOOKEEPER_TICK_TIME: 2000\n\n" +
		"  kafka:\n" +
		"    image: confluentinc/cp-kafka:7.6.1\n" +
		"    container_name: \"kafka\"\n" +
		"    depends_on:\n" +
		"      - zookeeper\n" +
		"    deploy:\n" +
		"      mode: replicated\n" +
		"      replicas: 1\n" +
		"    ports:\n" +
		"      - \"9092:9092\"\n" +
		"      - \"29092:29092\"\n" +
		"    environment:\n" +
		"      KAFKA_BROKER_ID: 1\n" +
		"      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181\n" +
		"      KAFKA_ADVERTISED_LISTENERS: \"PLAINTEXT://localhost:9092\"\n" +
		"      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1\n\n",
)

var MongoCompose = []byte(
	"  mongo:\n" +
		"    image: mongo\n" +
		"    restart: always\n" +
		"    ports:\n" +
		"      - \"27017:27017\"\n" +
		"    deploy:\n" +
		"      mode: replicated\n" +
		"      replicas: 1\n" +
		"    environment:\n" +
		"      MONGO_INITDB_ROOT_USERNAME: root\n" +
		"      MONGO_INITDB_ROOT_PASSWORD: example\n\n" +
		"  mongo-express:\n" +
		"    image: mongo-express\n" +
		"    restart: always\n" +
		"    ports:\n" +
		"      - \"8081:8081\"\n" +
		"    environment:\n" +
		"      ME_CONFIG_MONGODB_ADMINUSERNAME: root\n" +
		"      ME_CONFIG_MONGODB_ADMINPASSWORD: example\n" +
		"      ME_CONFIG_MONGODB_URL: \"mongodb://root:example@mongo:27017/\"\n\n",
)
