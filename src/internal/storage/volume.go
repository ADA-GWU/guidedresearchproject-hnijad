package storage

import (
	"errors"
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/util"
	log "github.com/google/logger"
	"os"
	"sync"
)

type Volume struct {
	ID       int
	Dir      string
	dataFile *os.File
	mu       sync.Mutex
}

//func NewVolume(id int) *Volume {
//	return &Volume{
//		ID: id,
//	}
//}

func CreateNewVolume(id int, dir string) (*Volume, error) {
	volumeName := dir + "/" + fmt.Sprintf("%d.data", id)

	if util.FileExists(volumeName) {
		log.Warningf("Volume %v exists", volumeName)
		return nil, errors.New(fmt.Sprintf("volume with %d id exists", id))
	}

	file, err := os.OpenFile(volumeName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Errorln("Error creating new volume:", err)
		return nil, err
	}
	volume := &Volume{
		ID:       id,
		Dir:      dir,
		dataFile: file,
	}
	return volume, nil
}

func (v *Volume) WriteNeedle(needle *Needle) error {
	bytes := needle.MapNeedleToBytes()
	_, err := v.dataFile.Write(bytes)
	return err
}

func (v *Volume) ReadNeedle(needleId int) (*Needle, error) {
	needles := make([]Needle, 0)
	offset := int64(0)
	totalSizeByte := make([]byte, 4)

	for {
		at, err := v.dataFile.ReadAt(totalSizeByte, offset)
		if err != nil {
			log.Errorln("Error while reading from the offset", err.Error())
			break
		}
		offset += int64(at)

		needle := Needle{}
		needle.TotalSize = util.BytesToUint(totalSizeByte)

		leftBytes := needle.TotalSize - 4
		data := make([]byte, leftBytes)

		at, err = v.dataFile.ReadAt(data, offset)
		if err != nil {
			log.Errorln("Error while reading from the offset", err.Error())
			break
		}

		needle = MapBytesToNeedle(data)

		needles = append(needles, needle)
		offset += int64(at)
	}

	return &needles[0], nil
}
