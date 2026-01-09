package templates

var DockerComposeTemplate = []byte("services:\n")

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
		"      - postgres:/var/lib/postgresql/data\n\n" +
		"volumes:\n" +
		"  postgres:\n",
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
		"      - mysql:/var/lib/mysql\n\n" +
		"volumes:\n" +
		"  mysql:\n",
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
		"      - sqlserver:/var/opt/mssql\n\n" +
		"volumes:\n" +
		"  sqlserver:\n",
)
