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

## GetAll operations

### Implemented customer table and GetAllCustomers

internal/models/model_customer.go
internal/database/database_customer.go

Create schema and import data

run main()
http :8080/customers
http :8080/customers emailAddress=="turpis@loremvehicula.com"

### Implement Product table and GetAll

internal/models/model_product.go
internal/database/database_product.go
internal/database/client.go DatabaseClient: DatabaseClient
internal/server/product.go
internal/server/server.go

run main()
http :8080/products

### Implement Vendor table and GetAll
internal/models/model_vendor.go
internal/database/database_vendor.go
internal/database/client.go DatabaseClient: DatabaseClient
internal/server/vendor.go
internal/server/server.go

run main()
http :8080/vendors

### Implement Service table and GetAll
internal/models/model_service.go
internal/database/database_service.go
internal/database/client.go DatabaseClient: DatabaseClient
internal/server/service.go
internal/server/server.go 

run main()
http :8080/services