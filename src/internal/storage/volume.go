package storage

import (
	"errors"
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/util"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

type Volume struct {
	ID        int
	Dir       string
	dataFile  *os.File
	NeedleMap map[int]*NeedleInfo
	UsedSpace int64
	mu        sync.Mutex
}

//func NewVolume(id int) *Volume {
//	return &Volume{
//		ID: id,
//	}
//}

func (v *Volume) WriteNeedle(needle *Needle) error {
	bytes := needle.MapNeedleToBytes()
	v.mu.Lock()
	defer v.mu.Unlock()
	sz, err := v.dataFile.Write(bytes)
	if err == nil {
		v.NeedleMap[int(needle.Id)] = &NeedleInfo{
			Offset: v.UsedSpace,
			Size:   uint32(sz),
			Name:   string(needle.Name),
		}
		v.UsedSpace += int64(sz)
		//log.Infof("id = %v, offset = %v ", needle.Id, v.NeedleMap[int(needle.Id)].Name)
	}
	return err
}

func (v *Volume) syncNeedleMap() {
	offset := int64(0)
	totalSizeByte := make([]byte, 4)
	for {
		needleOffset := offset
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
			log.Errorln("Error while reading from the offset ", err.Error())
			break
		}

		needle = MapBytesToNeedle(data)

		offset += int64(at)
		v.NeedleMap[int(needle.Id)] = &NeedleInfo{
			Offset: needleOffset,
			Size:   needle.TotalSize,
			Name:   string(needle.Name),
		}
	}
	v.UsedSpace = offset
	log.Infoln("Synced the needle map")
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

	for _, needle := range needles {
		if needle.Id == uint32(needleId) {
			return &needle, nil
		}
	}

	return nil, errors.New("not Found")
}

func (v *Volume) FindNeedle(needleId int) (*Needle, error) {
	//log.Infof("Find needle with id = %d\n", needleId)
	if _, ok := v.NeedleMap[needleId]; !ok {
		log.Warningf("Needle with id=%d was not found", needleId)
		return nil, errors.New(fmt.Sprintf("Needle with id=%d was not found", needleId))
	}

	offset := v.NeedleMap[needleId].Offset

	totalSizeByte := make([]byte, 4)

	at, err := v.dataFile.ReadAt(totalSizeByte, offset)
	if err != nil {
		return nil, err
	}
	offset += int64(at)

	needle := Needle{}
	needle.TotalSize = util.BytesToUint(totalSizeByte)

	leftBytes := needle.TotalSize - 4
	data := make([]byte, leftBytes)

	at, err = v.dataFile.ReadAt(data, offset)
	if err != nil {
		return nil, err
	}

	needle = MapBytesToNeedle(data)
	return &needle, nil
}
