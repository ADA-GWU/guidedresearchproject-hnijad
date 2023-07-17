package storage

import (
	"errors"
	"fmt"
	log "github.com/google/logger"
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
	volume, err := CreateNewVolume(id, s.Dir)
	if err != nil {
		return err
	}
	s.Volumes[id] = volume
	return nil
}

func (s *Storage) WriteNeedle(volumeId int, needle *Needle) error {
	if _, ok := s.Volumes[volumeId]; !ok {
		log.Warningf("Volume with id=%d does not exist in volume map", volumeId)
		return errors.New(fmt.Sprintf("volume  does not exist with id %d", volumeId))
	}
	err := s.Volumes[volumeId].WriteNeedle(needle)
	if err != nil {
		log.Errorln("Could not write the needle with id %d", needle.Id)
		return err
	}
	return nil
}

func (s *Storage) ReadNeedle(volumeId int, needleId int) (*Needle, error) {
	if _, ok := s.Volumes[volumeId]; !ok {
		log.Warningf("Volume with id=%d does not exist in volume map", volumeId)
		return nil, errors.New(fmt.Sprintf("volume  does not exist with id %d", volumeId))
	}
	needle, err := s.Volumes[volumeId].ReadNeedle(needleId)
	if err != nil {
		log.Errorln("Could not read the needle with id %d", needleId)
		return nil, err
	}
	return needle, err
}
