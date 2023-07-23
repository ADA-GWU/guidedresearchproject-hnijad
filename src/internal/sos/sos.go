package sos

import (
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/handler"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/server"
	storage2 "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/storage"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime"
	"strings"
)

func init() {
	log.SetReportCaller(true)
	formatter := &log.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
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

func RunDataNode(volDir, port, primaryNode, noteId string) {
	e := echo.New()
	e.HideBanner = true

	storage := storage2.NewStorage(volDir)

	dataServer := server.NewDataServer(noteId, storage)

	handler.AddDataRoutes(e, dataServer)

	if err := e.Start(":" + port); err != http.ErrServerClosed {
		log.Infoln("Error when starting http server", err.Error())
	}
}
