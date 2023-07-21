package server

import (
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/storage"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"unsafe"
)

type DataServer struct {
	ID      string
	Storage *storage.Storage
}

func (ds *DataServer) CreateNewVolume(id int) error {
	log.Infoln("starting to create a volume")
	if err := ds.Storage.CreateNewVolume(id); err != nil {
		return err
	}
	return nil
}

func (ds *DataServer) WriteObject(fid string, fileName string, file multipart.File, bytes []byte) error {
	log.Infoln(fid, fileName)

	tokens := strings.Split(fid, ",")
	vId, _ := strconv.Atoi(tokens[0])
	oId, _ := strconv.Atoi(tokens[1])

	fileNameByte := []byte(fileName)
	fileNameLength := len(fileNameByte)

	mimeType := []byte(http.DetectContentType(bytes))
	mimeTypeLength := len(mimeType)

	data := bytes
	dataSize := len(data)

	uintSize := int(unsafe.Sizeof(uint32(1)))

	needle := &storage.Needle{
		TotalSize: uint32(6*uintSize + fileNameLength + mimeTypeLength + dataSize),
		Id:        uint32(oId),
		NameSize:  uint32(fileNameLength),
		Name:      fileNameByte,
		MimeSize:  uint32(mimeTypeLength),
		Mime:      mimeType,
		DataSize:  uint32(dataSize),
		Data:      data,
		Checksum:  uint32(12),
	}

	if err := ds.Storage.WriteNeedle(vId, needle); err != nil {
		log.Errorln("Error writing needle to storage", err.Error())
	}
	return nil
}

func (ds *DataServer) ReadObject(fid string) (*storage.Needle, error) {
	tokens := strings.Split(fid, ",")
	vId, _ := strconv.Atoi(tokens[0])
	oId, _ := strconv.Atoi(tokens[1])
	return ds.Storage.ReadNeedle(vId, oId)
}

func NewDataServer(id string, store *storage.Storage) *DataServer {
	return &DataServer{
		ID:      id,
		Storage: store,
	}
}
