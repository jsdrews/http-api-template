This is a template for a web api project. It uses the following technologies:

- [Go](https://golang.org/)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [OpenAPI](https://swagger.io/specification/)
- [Docker](https://www.docker.com/)
- [Task](https://taskfile.dev/#/)
- [MockServer](https://www.mock-server.com/)
- [Schemathesis](https://schemathesis.readthedocs.io/en/stable/index.html)
- [Direnv](https://direnv.net/)

## Quick Start

### Use Cruft to Create a New Project

#### Install Cruft

https://cruft.github.io/cruft/#installation


#### Create a New Project

```bash
cruft create https://github.com/jsdrews/openapi-codegen
```

### Install Task

This project uses `task` instead of make. Install it via: https://taskfile.dev/installation
  
  **Note:** there is a `Makefile` in the project's root directory. This is only used to check to make sure you have installed everything that is needed for development. However, it will install `task` into a directory in the project root for convenience. For longterm, you will probably want this in a global system location as opposed to local to this project.

### Run the Web API

```bash
task up
```

### Run the Web API Mock

```bash
task mock
```
