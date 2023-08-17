# Proof Of Concept

## Description
This is a proof of concept project that is centered around generating a wep api from an openapi specification. See CONTRIBUTING.md for more information.

### Quick Start
The openapi specification is located in the root of the src folder (`src`) and is called `openapi.yaml`. This file is used to generate the web api.

#### Install dependencies
Run `make` to check which dependencies are needed to install. If `task` is not installed, it will install it at `./bin/task`. `task` is used from here on out to run the project.

#### List all tasks
Run: `task`

This will show all available tasks.

#### Run the api
`task up`
This will start the api at `localhost:$APP_PORT` (default is `8888`).
