This is a template for a web api project. It uses the following technologies:

- [Go](https://golang.org/)
- [OpenAPI](https://swagger.io/specification/)
- [Swagger](https://swagger.io/)
- [Docker](https://www.docker.com/)
- [Task](https://taskfile.dev/#/)
- [MockServer](https://www.mock-server.com/)
- [Schemathesis](https://schemathesis.readthedocs.io/en/stable/index.html)

## Quick Start

### Use Cruft to Create a New Project

#### Install Cruft

https://cruft.github.io/cruft/#installation


#### Create a New Project

```bash
cruft create https://github.com/jsdrews/openapi-codegen
```

### Run the Web API

```bash
task up
```

### Run the Web API Mock

```bash
task mock
```
