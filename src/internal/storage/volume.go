package storage

import (
	"errors"
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/util"
	log "github.com/google/logger"
	"os"
)

type Volume struct {
	ID int
}

//func NewVolume(id int) *Volume {
//	return &Volume{
//		ID: id,
//	}
//}

func CreateVolume(id int, dir string) error {
	volumeName := dir + "/" + fmt.Sprintf("%d.data", id)

	if util.FileExists(volumeName) {
		log.Warningf("Volume %v exists", volumeName)
		return errors.New(fmt.Sprintf("volume with %d id exists", id))
	}

	file, err := os.OpenFile(volumeName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Errorln("Error creating new volume:", err)
		return err
	}
	if err = file.Close(); err != nil {
		log.Errorln("Could not close the new volume")
		return err
	}
	return nil
}
