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

func (h *PrimaryNodeHandler) findAvailableDataNode(ctx echo.Context) error {
	res, _ := h.PrimaryServer.FindDataNode()
	return ctx.JSON(http.StatusOK, res)
}

func (h *PrimaryNodeHandler) findDataNodeByObjectId(ctx echo.Context) error {
	objectId := ctx.QueryParam("id")
	res, _ := h.PrimaryServer.FindDataNodeByObjectId(objectId)
	return ctx.JSON(http.StatusOK, res)
}

func AddPrimaryRoutes(e *echo.Echo, server *server.PrimaryServer) {
	handler := &PrimaryNodeHandler{
		PrimaryServer: server,
	}

	e.GET("/primary/cluster-info", handler.getClusterInfo)
	e.GET("/primary/volume", handler.findAvailableDataNode)
	e.GET("/primary/search", handler.findDataNodeByObjectId)
}
