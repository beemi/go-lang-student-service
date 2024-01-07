# 🌟Student service🌟

This is a simple student service that allows you to create, read, update and delete students using go lang and mongodb.

## Installation

Use the package manager [go](https://golang.org/) to install the student service.

```bash
go get github.com/gorilla/mux

go get go.mongodb.org/mongo-driver/mongo

go get github.com/joho/godotenv

```

## Usage

Run mongodb server on your local machine using the Docker compose file in the root directory.
```bash
docker-compose up -d
```

```bash
go run main.go
```