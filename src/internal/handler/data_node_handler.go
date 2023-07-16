package handler

import (
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/server"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type DataNodeHandler struct {
	DataServer *server.DataServer
}

type VolumeCreateRequest struct {
	VolumeId int `json:"volumeId"`
}

func (h *DataNodeHandler) CreateVolume(ctx echo.Context) error {
	req := &VolumeCreateRequest{}

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": "could not unmarshall json in the request",
		})
	}
	if err := h.DataServer.CreateNewVolume(req.VolumeId); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": "could not create the volume: err =  " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "volume created successfully",
	})
}

func (h *DataNodeHandler) UploadObject(ctx echo.Context) error {
	fid := ctx.Param("fid")
	file, err := ctx.FormFile("file")

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": "could not get object",
		})
	}

	src, err := file.Open()

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": "could not open the object",
		})
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	fileName := file.Filename

	err = h.DataServer.WriteObject(fid, fileName, src, fileBytes)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": "could not write the object",
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "object uploaded successfully",
	})
}

func (h *DataNodeHandler) DownloadObject(ctx echo.Context) error {
	fid := ctx.Param("fid")

	needle, err := h.DataServer.ReadObject(fid)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "could not read the object" + err.Error(),
		})
	}

	ctx.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+string(needle.Name))
	ctx.Response().Header().Set(echo.HeaderContentType, string(needle.Mime))
	return ctx.Blob(http.StatusOK, echo.MIMEOctetStream, needle.Data)
}

func AddDataRoutes(e *echo.Echo, server *server.DataServer) {
	handler := &DataNodeHandler{
		DataServer: server,
	}
	e.POST("/data/volume", handler.CreateVolume)
	e.POST("/data/:fid", handler.UploadObject)
	e.GET("/data/:fid", handler.DownloadObject)
}
