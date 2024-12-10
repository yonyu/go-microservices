# Create a Microservice with Golang
This microservice implements CRUD operations on a PostgresDB hosted in
a docker container.

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

## Create operations

### Create customer

internal/database/database_customer.go: AddCustomer()
internal/database/client.go: add AddCustomer() to DatabaseClient
internal/server/customer.go: AddCustomer()
internal/server/server.go: paste AddCustomer() to Server and  add it to registerRoutes()

run main()
http POST :8080/customers firstName=John lastName=Doe emailAddress="jdoe@example.com" phoneNumber="515-555-1234" address="123 Main St; Anytown, KS 66854"

### Create product

internal/database/database_product.go: AddProduct()
internal/database/client.go: add AddProduct() to DatabaseClient
internal/server/product.go: AddProduct()
internal/server/server.go: paste AddProduct() to Server and  add it to registerRoutes()

run main()
http POST :8080/products name=test2 price:=3.14 vendorId="e27d6cda-ef90-4d92-8f89-648389edb8af"
Or Use Postman
{
    "name":"Test",
    "price": 3.14,
    "vendorId": "e27d6cda-ef90-4d92-8f89-648389edb8af"
}

### Create service

internal/database/database_service.go: AddService()
internal/database/service.go: add AddService() to DatabaseClient
internal/server/service.go: AddService()
internal/server/server.go: paste AddService() to Server and  add it to registerRoutes()

run main()
http POST :8080/services name=Service1 price:=6.99

