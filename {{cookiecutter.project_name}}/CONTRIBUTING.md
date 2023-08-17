# Proof Of Concept

## Description
This is a proof of concept project that is centered around generating a wep api from an openapi specification.

## Install dependencies
Run `make` to check which dependencies are needed to install. If `task` is not installed, it will install it at `./bin/task`. `task` is used from here on out to run the project.

### List all tasks
Run: `task`

This will show all available tasks.

## Front End Development

Everything is centered around the api spec. Make changes to the api spec to reflect what is needed from the api for the front end. While the backend features are developed, one can use `task mock` to mock the api and develop the front end.

If you need to run the full web api, you can run `task up` which will run the web api at address: `localhost:$API_PORT`.

### OpenAPI Specification
The openapi specification is located in the root of the src folder (`src`) and is called `openapi.yaml`. This file is used to generate the web api.

#### Editing the OpenAPI Specification
Run `task edit` to run the editor. Paste `http://localhost:$SWAGGER_PORT` into the browser to edit the api spec. Save & overwrite the `openapi.yaml` file to update the api (this can be done in the editor).

#### Quick Mocking the Web Server API
Run `task mock` to run the web api mock. This will start the web api mock on port `$MOCKSERVER_PORT` (`http://localhost:$MOCKSERVER_PORT/api/v1`). You can then use the mock to test the web api. The mock data it serves will be the examples in the api spec.

#### Running the Web API
Run `task up` to run the web api. This will start the web api on port `$API_PORT`.

## Backend Development

### OpenAPI Specification
The openapi specification is located in the root of the src folder (`src`) and is called `openapi.yaml`. This file is used to generate the web api.

#### Editing the OpenAPI Specification
Run `task edit` to run the editor. Paste `http://localhost:$SWAGGER_PORT` into the browser to edit the api spec. Save & overwrite the `openapi.yaml` file to update the api (this can be done in the editor).

#### Generating the Web Server API
Run `task generate-server` to generate the web api. This will generate the server interface and types code in the `src/server/api` folder (everytime you edit the spec you will need to regenerate the api). This code will be available internally via the `api` package. 

#### Developing the Web API
Every time you make a change to the api spec, you will need to run `task generate-server` to update the `api` package. On top of that you will need to add controllers to `src/server/controllers` to satisfy the auto generated interface. The controllers will be available internally via the `controllers` package.

An example of a controller is the following:
```go
func (s Server) GetUsers(ctx context.Context, request api.GetUsersRequestObject) (api.GetUsersResponseObject, error) {
    ...
}
```

You can predict what the controller method names will be by looking at the operationKey in the api spec. For example, the operationKey for the `GetUsers` controller is `getUsers`.

If in doubt, you can always look at the auto generated interface in `src/server/api/api.go` to see what the controller method names will be.
Look for the `StrictServerInterface` definition (example below):

```go
type StrictServerInterface interface {
	// Get Users
	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Add Users
	// (POST /users)
	AddUsers(ctx context.Context, request AddUsersRequestObject) (AddUsersResponseObject, error)
}
```

#### Running the Web API
Run `task up` to run the web api. This will start the web api on port `$API_PORT`.

#### Testing the Web API
Run `task test` to run the web api tests. This will run the tests in the `src/server/api` folder.
Currently these are auto generated per the openapi spec. See https://schemathesis.readthedocs.io/en/stable/index.html for more information.
