package storage

import (
	"errors"
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/util"
	log "github.com/google/logger"
	"os"
)

type Storage struct {
	Dir     string
	Volumes map[int]*Volume
}

func NewStorage(dir string) *Storage {
	return &Storage{
		Dir:     dir,
		Volumes: make(map[int]*Volume),
	}
}

func (s *Storage) CreateNewVolume(id int) error {

	if _, ok := s.Volumes[id]; ok {
		log.Warningf("Volume with id=%d exists in volume map", id)
		return errors.New(fmt.Sprintf("volume already exists with id %d", id))
	}
	if err := CreateVolume(id, s.Dir); err != nil {
		return err
	}
	s.Volumes[id] = &Volume{id}
	return nil
}

func (s *Storage) WriteNeedle(volumeId int, needle *Needle) error {
	file, err := os.OpenFile("tmp/"+fmt.Sprintf("%d.data", volumeId), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Errorln("Could not open the volume ", volumeId)
		return err
	}
	defer file.Close()
	bytes := needle.MapNeedleToBytes()
	_, err = file.Write(bytes)
	if err != nil {
		log.Errorln("Error writing to file:", err)
		return err
	}
	return nil
}

func (s *Storage) ReadNeedle(volumeId int, needleId int) (*Needle, error) {
	volume, err := os.Open("tmp/" + fmt.Sprintf("%d.data", volumeId))
	if err != nil {
		log.Errorln("Could not open the volume ", volumeId)
		return nil, err
	}
	defer volume.Close()

	needles := make([]Needle, 0)
	offset := int64(0)
	totalSizeByte := make([]byte, 4)

	for {
		at, err := volume.ReadAt(totalSizeByte, offset)
		if err != nil {
			log.Errorln("Error while reading from the offset", err.Error())
			break
		}
		offset += int64(at)

		needle := Needle{}
		needle.TotalSize = util.BytesToUint(totalSizeByte)

		leftBytes := needle.TotalSize - 4
		data := make([]byte, leftBytes)

		at, err = volume.ReadAt(data, offset)
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
