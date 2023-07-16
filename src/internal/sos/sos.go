package sos

import (
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/handler"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/server"
	storage2 "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/storage"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Run() {
	e := echo.New()
	e.HideBanner = true

	storage := storage2.NewStorage("tmp")

	dataServer := server.NewDataServer("dnode_1", storage)

	handler.AddDataRoutes(e, dataServer)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Println("Error when starting http server", err.Error())
	}
}
