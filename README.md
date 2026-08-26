# Gini

[![Language](https://img.shields.io/badge/language-Go-00ADD8?labelColor=2F2F2F)](https://go.dev/doc/)
[![Version](https://img.shields.io/badge/version-1.27-9C27B0?labelColor=2F2F2F)](https://go.dev/doc/install)
[![Go Reference](https://pkg.go.dev/badge/github.com/rafaeldepontes/gini.svg)](https://pkg.go.dev/github.com/rafaeldepontes/gini)

Gini is an open-source CLI tool that speeds up Go project setup by generating a clean, structured project layout with sensible defaults and essential tooling.

It removes repetitive boilerplate so you can focus on building your application logic.

---

## Features

- Fast project scaffolding
- Ready-to-use configuration and sensible defaults
- Minimal setup with no unnecessary friction
- Ideal for small Go projects and MVPs

---

<img src="public/Gini.png" alt="Gini preview" />

---

## Usage

```bash
gini build
# or
gini b
```

This creates a new directory with the predefined structure and configuration.

---

## Installation

### Go install

```bash
go install github.com/rafaeldepontes/gini@latest
```

This installs the `gini` binary into your `$GOPATH/bin` directory.

Make sure `$GOPATH/bin` is included in your `$PATH`.

---

## Requirements

- Go 1.27.0 or newer

---

## Why Gini?

Bootstrapping Go projects over and over can be repetitive. Setting up the Dockerfile, docker-compose, and basic project structure from scratch takes time that could be spent on actual development.

Gini provides a consistent starting point and helps you:

- Reduce setup time
- Keep project structure consistent
- Generate Docker-related files with the essentials already in place

---

## Roadmap

- Custom templates
- Docker support
- Configurable project layouts

---

## Contributing

Contributions are welcome.

Feel free to open an issue or submit a pull request.

---

## License

[MIT](LICENSE)

## Contact

If something goes wrong, contact: `rafael.cr.carneiro@gmail.com`
