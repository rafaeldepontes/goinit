package templates

var DockerFile = []byte(
	"# Optimized build, if you need something smaller...\n" +
		"FROM golang:1.25-alpine AS build\n\n" +
		"WORKDIR /cmd\n\n" +
		"COPY go.mod go.sum ./\n" +
		"RUN go mod download\n\n" +
		"COPY . .\n\n" +
		"RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /app .\n\n" +
		"# SWEET LITTLE BINARY\n" +
		"FROM gcr.io/distroless/static-debian12\n\n" +
		"WORKDIR /\n\n" +
		"COPY --from=build /app /app\n\n" +
		"USER nonroot:nonroot\n\n" +
		"ENTRYPOINT [\"/app\"]\n",
)
