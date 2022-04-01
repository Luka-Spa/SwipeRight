# SwipeRight Profile Service

## Table of Contents

1. How to run?
2. Techstack
3. Folder Structure

## How to run?

1. Make sure [Apache Cassandra](https://cassandra.apache.org/_/index.html) is installed & running
2. Run `go get` in the current directory to install the dependencies
3. Run `go run .` to run the project

## Techstack

| Technology       | Description                                |
| ---------------- | ------------------------------------------ |
| Go               | Go is used as the programming language     |
| Gin              | Gin is the web framework                   |
| Apache Cassandra | Cassandra is used as the database          |
| Docker           | Docker is used to containerize the service |

## Folder Structure

```c++
profile
│   README.md
│   ...
│
└───router
│   │
│   └───http
│       │   router.go // Registers HTTP Routes
│       │   ...
|       └───handler
|           |   user_handler.go // Contains the user route handler functions
|           |   ...
│
└───repository
|   │   init.go // Initialized db connections
|   │   ...
|   └───cassandra
|       | user_repository.go // Communicates with the Cassandra database
└───logic
|   │   user_logic.go // Contains the user logic
|   │   ...
└───model
|   |   user_profile.go // Model for the user
|   |   ...
└───config
    |   config.go // Initializes the configuration & provides GetConfig function
    |   development.yaml // Contains the development configuration
    |   ...
```
