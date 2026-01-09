# Go Init

Go Init is an open-source CLI tool that accelerates Go project setup by generating a structured project layout with essential tooling and sensible defaults.

It removes repetitive boilerplate so you can focus on writing actual business logic.

---

## Features

- Opinionated Go project structure
- Ready-to-use configuration and defaults
- Fast project scaffolding
- Minimal setup, zero friction
- Designed for real-world Go applications

---

## Installation

### Using Go (recommended)

```bash
go install github.com/rafaeldepontes/goinit@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

---

## Usage

Initialize a new Go project:

```bash
ginit build
```

This will create a new directory with the predefined structure and configuration.

---

## Requirements

- Go 1.25.5 or newer

---

## Why Go Init?

Setting up Go projects repeatedly is tedious and setting up the dockerfile and docker-compose is kinda lame...
Go Init provides a consistent starting point, helping you:

- Reduce setup time
- Maintain project structure consistency
- Create your Dockerfile and docker-compose with everything you need.

---

## Roadmap

- Custom templates
- Docker
- Configurable project layouts

---

## Contributing

Contributions are welcome.
Feel free to open issues or submit pull requests.

---

## License

[MIT](LICENSE)

## Contact

If something went wrong, please contact: `rafael.cr.carneiro@gmail.com`
