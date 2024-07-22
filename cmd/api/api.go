package api

import (
	"database/sql"

	"github.com/berkaycubuk/mqtt-studio/service/health"
	"github.com/berkaycubuk/mqtt-studio/service/user"
	"github.com/labstack/echo/v4"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	e := echo.New()

	apiV1 := e.Group("/api/v1")

	healthHandler := health.NewHandler()
	healthHandler.RegisterRoutes(apiV1)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(apiV1)

	return e.Start(s.addr)
}
