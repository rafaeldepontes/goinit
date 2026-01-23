package templates

var DockerFile = []byte(
	"FROM golang:1.25\n\n" +
		"WORKDIR /usr/src/app\n\n" +
		"COPY go.mod go.sum ./\n" +
		"RUN go mod download\n\n" +
		"COPY . .\n" +
		"RUN go build -v -o /usr/local/bin/app .\n\n" +
		"CMD [\"app\"]\n",
)
