package server

import (
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/client"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/config"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/storage"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type DataServer struct {
	pb.UnimplementedDataNodeServer
	ID                string
	Storage           *storage.Storage
	PrimaryGrpcClient *client.PrimaryGrpcClientWrapper
	Params            *config.DataNodeParams
}

func (ds *DataServer) CreateNewVolume(id int) error {
	log.Infoln("starting to create a volume")
	if err := ds.Storage.CreateNewVolume(id); err != nil {
		return err
	}
	return nil
}

func (ds *DataServer) WriteObject(fid string, fileName string, file multipart.File, bytes []byte) error {
	//log.Infoln(fid, fileName)

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

func NewDataServer(id string, store *storage.Storage, primaryGrpcClient *client.PrimaryGrpcClientWrapper, params *config.DataNodeParams) *DataServer {
	return &DataServer{
		ID:                id,
		Storage:           store,
		PrimaryGrpcClient: primaryGrpcClient,
		Params:            params,
	}
}

func (ds *DataServer) StartHeartBeat() {
	log.Infoln("Starting the heartbeat")
	ticker := time.NewTicker(1000 * time.Millisecond)

	go func() {
		for {
			select {
			case <-ticker.C:
				//log.Infoln("Heartbeat at", ticker)
				info := &pb.DataNodeInfo{
					Id:       ds.ID,
					Address:  getOutboundUp(),
					Volumes:  asVolumeList(ds.Storage.Volumes),
					HttpPort: ds.Params.HttpPort,
					GrpcPort: ds.Params.GRPCPort,
				}
				ds.PrimaryGrpcClient.HeartBeat(info)
			}
		}
	}()
}

func getOutboundUp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Errorln("Error getting outbound IP:", err)
		os.Exit(1)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr).IP.String()
	return localAddr
}

func asVolumeList(volumeMap map[int]*storage.Volume) []*pb.Volume {
	volumes := make([]*pb.Volume, 0)
	for _, value := range volumeMap {
		volume := &pb.Volume{
			Id:        int32(value.ID),
			Dir:       value.Dir,
			UsedSpace: value.UsedSpace,
			FreeSpace: 12,
		}
		volumes = append(volumes, volume)
	}
	return volumes
}
