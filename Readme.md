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

### Create vendor

internal/database/database_vendor.go: AddVendor()
internal/database/vendor.go: add AddVendor() to DatabaseClient
internal/server/vendor.go: AddVendor()
internal/server/server.go: paste AddVendor() to Server and  add it to registerRoutes()

run main()
http POST :8080/vendors name=Vendor1 contact="John Smith" phoneNumber="1-777-123-1234" emailAddress="test@gmail.com" address="1234 Main Street, Vancouver, BC V6V 1V1"

## GetOne operations

### Get a customer by ID

internal/database/database_customer.go: GetCustomerById()
internal/database/client.go: add GetCustomerById() to DatabaseClient
internal/server/customer.go: GetCustomerById()
internal/server/server.go: paste GetCustomerById() to Server and  add it to registerRoutes()

run main()
http :8080/customers/50cb40eb-cdc0-4235-9d8f-f598dcc7b9a9


### Get a product, service and vendor by ID

http :8080/products/ff3bef2f-9bf7-4edd-a141-89b312d5fd21
http :8080/services/8983ca78-37c5-4e78-9b15-2b2899163758
http :8080/vendors/e27d6cda-ef90-4d92-8f89-648389edb8af

## Update operations

### Update customer

internal/database/database_customer.go: UpdateCustomer()
internal/database/client.go: add UpdateCustomer() to DatabaseClient
internal/server/customer.go: UpdateCustomer()
internal/server/server.go: paste UpdateCustomer() to Server and add it to registerRoutes()

run main()
http PUT :8080/customers/50cb40eb-cdc0-4235-9d8f-f598dcc7b9a9 customerId="50cb40eb-cdc0-4235-9d8f-f598dcc7b9a9" firstName=John lastName=Doe emailAddress="jdoe@example.com" phoneNumber="515-555-1234" address="1234 Main St; Anytown, KS 66854"


### Update product, service and vendor

internal/database/database_customer.go: UpdateCustomer()
internal/database/client.go: add UpdateCustomer() to DatabaseClient
internal/server/customer.go: UpdateCustomer()
internal/server/server.go: paste UpdateCustomer() to Server and add it to registerRoutes()

run main()
http PUT :8080/products/ff3bef2f-9bf7-4edd-a141-89b312d5fd21 productId="ff3bef2f-9bf7-4edd-a141-89b312d5fd21" name=test3 price:=3.15 vendorId="e27d6cda-ef90-4d92-8f89-648389edb8af"
http PUT :8080/services/45763e43-e7e3-4114-8eb1-ca7f0cc2d88a serviceId="45763e43-e7e3-4114-8eb1-ca7f0cc2d88a" name=Service2 price:=7.99
http PUT :8080/vendors/1eda44e0-d29e-4df8-86c3-dee8d578e34b vendorId="1eda44e0-d29e-4df8-86c3-dee8d578e34b" name="vendor2" phoneNumber="1-777-123-1234" emailAddress="test@gmail.com" address="1234 Main Street, Vancouver, BC V6V 1V1"