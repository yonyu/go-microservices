package server

import (
	"github.com/labstack/echo/v4"
	"github.com/yonyu/go-microservices/internal/database"
	"github.com/yonyu/go-microservices/internal/models"
	"log"
	"net/http"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	GetAllCustomers(ctx echo.Context) error
	GetAllProducts(ctx echo.Context) error
	GetAllVendors(ctx echo.Context) error
	GetAllServices(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

// NewEchoServer creates a new EchoServer and registers routers
func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

// Start calls echo.Start() to start the web server
func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server shutdown occurred: #{err}")
		return err
	}
	return nil
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	cg := s.echo.Group("/customers")
	cg.GET("", s.GetAllCustomers)

	pg := s.echo.Group("/products")
	pg.GET("", s.GetAllProducts)

	vg := s.echo.Group("/vendors")
	vg.GET("", s.GetAllVendors)

	sg := s.echo.Group("/services")
	sg.GET("", s.GetAllServices)
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}
	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
