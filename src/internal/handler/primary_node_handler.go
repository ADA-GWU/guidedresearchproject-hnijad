package handler

import (
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PrimaryNodeHandler struct {
	PrimaryServer *server.PrimaryServer
}

func (h *PrimaryNodeHandler) getClusterInfo(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.PrimaryServer.GetClusterInfo())
}

func AddPrimaryRoutes(e *echo.Echo, server *server.PrimaryServer) {
	handler := &PrimaryNodeHandler{
		PrimaryServer: server,
	}

	e.GET("/primary/cluster-info", handler.getClusterInfo)
}
