package sos

import (
	"errors"
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/client"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/config"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/handler"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/server"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/storage"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime"
	"strings"
)

func init() {
	log.SetReportCaller(true)
	formatter := &log.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuration
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", formatFilePath(f.File), f.Line)
		},
	}
	log.SetFormatter(formatter)
}
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func RunDataNode(params *config.DataNodeParams) {
	e := echo.New()
	e.HideBanner = true

	dataStorage := storage.NewStorage(params.VolDir)

	// TODO ID  is already in params
	dataServer := server.NewDataServer(params.NodeId, dataStorage, client.NewMasterGrpcClient(params.PrimaryNodeUrl), params)

	go dataServer.StartHeartBeat()

	handler.AddDataRoutes(e, dataServer)

	if err := e.Start(":" + params.HttpPort); !errors.Is(err, http.ErrServerClosed) {
		log.Infoln("Error when starting data node http server", err.Error())
	}
}

func RunPrimaryNode(params *config.PrimaryNodeParams) {
	e := echo.New()
	e.HideBanner = true

	primaryServer := server.NewPrimaryServer(params)

	handler.AddPrimaryRoutes(e, primaryServer)

	go server.StartPrimaryNodeGrpcServer(primaryServer)

	if err := e.Start(":" + params.HttpPort); !errors.Is(err, http.ErrServerClosed) {
		log.Infoln("Error when starting primary node http server", err.Error())
	}
}
