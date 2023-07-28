package handler

import (
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/server"
	"github.com/labstack/echo/v4"
)

type PrimaryNodeHandler struct {
	PrimaryServer *server.PrimaryServer
}

func AddPrimaryRoutes(e *echo.Echo, server *server.PrimaryServer) {
	_ = &PrimaryNodeHandler{
		PrimaryServer: server,
	}
}
