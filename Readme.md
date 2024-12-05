# Create a Microservice with Golang

## Initialize the go module

go mod init github.com/yonyu/go-microservices

## Install the required packages

go get github.com/google/uuid

go get github.com/labstack/echo/v4 github.com/lib/pq gorm.io/gorm gorm.io/driver/postgres

## Set up database

internal/database/client.go

internal/dberrors/conflict.go

## Set up server

internal/server/server.go

## Models: database model and server model

