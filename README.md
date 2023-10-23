# Album Store REST API

## Overview
This project is a RESTful API for an album store, implemented in Go. It's designed to demonstrate best practices in Go project structure and code design. The project uses the Echo web framework and MongoDB for data storage. 

**Note**: This project is for learning purposes and not intended for production use.

## Features
- CRUD operations for Albums (Create, Read, Update, Delete)
- Connection to MongoDB for data storage
- Organised into separate packages for better modularity
- Makefile for easy build and run commands

## Project Structure

.
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── internal
│   ├── handlers
│   │   └── album_handlers.go
│   ├── models
│   │   └── album.go
│   └── storage
│       ├── mongo.go
│       └── storage.go
├── main.go
├── podman-compose.yml
└── todo.md


## Prerequisites
- Go (version 1.17+)
- MongoDB
- Podman (for containerised MongoDB)

## Setup
1. Clone the repository to your local machine.
2. Navigate to the project folder and run `make build` to build the project.
3. Run `podman-compose up` to start MongoDB in a container.

## Running the Application
- Run `make run` to start the API server.
- The server will start on port 8080 (configurable).

## Endpoints
- GET `/albums`: Fetch all albums
- GET `/albums/:id`: Fetch a single album by ID
- POST `/albums`: Create a new album

## Future Enhancements
For a detailed list of future enhancements, see [TODO.md](./TODO.md).

