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

internal/models/model_status.go

## Set up docker

make sure no docker image running

    docker ps

run docker script

    cd dat
    ./postgres_start.sh

## Wire the service

modify main()

Run the server by click run icon at func main()

Install http and then run http

    sudo snap install http

    http :8080/readiness

GetAll operations

Create schema and import data

http :8080/customers

http :8080/customers emailAddress=="turpis@loremvehicula.com"

