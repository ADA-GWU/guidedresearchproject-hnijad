//go:build !linux

package storage

import (
	"errors"
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/util"
	log "github.com/sirupsen/logrus"
	"os"
)

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
		ID:        id,
		Dir:       dir,
		dataFile:  file,
		NeedleMap: make(map[int]*NeedleInfo),
		UsedSpace: 0,
	}
	return volume, nil
}
