package server

import (
	"bytes"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/util"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

type PrimaryNodeState struct {
	mu           sync.Mutex
	stateFile    *os.File
	LastObjectId uint32
	MaxVolumeId  uint32
}

func NewPrimaryNodeState(stateFilePath string) *PrimaryNodeState {
	if util.FileExists(stateFilePath) {
		file, err := os.OpenFile(stateFilePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("Could not open primary state file")
			return nil
		}

		state := &PrimaryNodeState{
			stateFile: file,
		}
		_ = state.loadValuesFromDisk()
		return state
	}
	file, err := os.OpenFile(stateFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Could not open primary state file")
		return nil
	}

	state := &PrimaryNodeState{
		stateFile:    file,
		LastObjectId: 0,
		MaxVolumeId:  0,
	}
	_ = state.saveToDisk()

	return state
}

func (s *PrimaryNodeState) loadValuesFromDisk() error {
	// read max volume
	offset := int64(0)
	totalSizeByte := make([]byte, 4)

	at, err := s.stateFile.ReadAt(totalSizeByte, offset)
	if err != nil {
		log.Errorln("Error while reading from the offset", err.Error())
		return err
	}
	offset += int64(at)

	s.MaxVolumeId = util.BytesToUint(totalSizeByte)

	// Read last object id
	at, err = s.stateFile.ReadAt(totalSizeByte, offset)
	if err != nil {
		log.Errorln("Error while reading from the offset", err.Error())
		return err
	}
	s.LastObjectId = util.BytesToUint(totalSizeByte)
	return nil
}

func (s *PrimaryNodeState) saveToDisk() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	buffer := bytes.Buffer{}
	buffer.Write(util.UintToBytes(s.MaxVolumeId))
	buffer.Write(util.UintToBytes(s.LastObjectId))

	_, err := s.stateFile.Write(buffer.Bytes())

	if err != nil {
		log.Fatalln("Could not save the primary node state")
		return err
	}
	return nil
}
