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

// TODO: Test this.
// var DockerFile = []byte(
// 	"FROM golang:1.25 AS build\n\n" +
// 		"WORKDIR /app\n\n" +
// 		"COPY go.mod go.sum ./\n" +
// 		"RUN go mod download\n\n" +
// 		"COPY . .\n" +
// 		"RUN go build -v -o /app .\n\n" +
// 		"FROM grc.io/distroless/base-debian10\n\n" +
// 		"WORKDIR /\n\n" +
// 		"COPY --from=build /app /app\n\n" +
// 		"USER nonroot:nonroot\n\n" +
// 		"CMD [\"/app\"]\n",
// )
