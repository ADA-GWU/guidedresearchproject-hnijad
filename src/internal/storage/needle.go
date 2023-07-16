package storage

import (
	"bytes"
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/util"
	"io"
	"os"
	"path/filepath"
	"unsafe"
)

type Needle struct {
	TotalSize uint32
	Id        uint32
	NameSize  uint32
	Name      []byte
	MimeSize  uint32
	Mime      []byte
	DataSize  uint32
	Data      []byte
	Checksum  uint32
}

func (n *Needle) MapNeedleToBytes() []byte {
	buffer := bytes.Buffer{}
	buffer.Write(util.UintToBytes(n.TotalSize))
	buffer.Write(util.UintToBytes(n.Id))
	buffer.Write(util.UintToBytes(n.NameSize))
	buffer.Write(n.Name)
	buffer.Write(util.UintToBytes(n.MimeSize))
	buffer.Write(n.Mime)
	buffer.Write(util.UintToBytes(n.DataSize))
	buffer.Write(n.Data)
	buffer.Write(util.UintToBytes(n.Checksum))
	return buffer.Bytes()
}

func MapBytesToNeedle(data []byte) Needle {
	needle := Needle{}
	// read ID
	a, b := 0, 4
	needle.Id = util.BytesToUint(data[a:b])

	// read name size
	a = b
	b += 4
	needle.NameSize = util.BytesToUint(data[a:b])

	// read name
	a = b
	b += int(needle.NameSize)
	needle.Name = data[a:b]

	// read mime size
	a = b
	b += 4
	needle.MimeSize = util.BytesToUint(data[a:b])

	// read mime type
	a = b
	b += int(needle.MimeSize)
	needle.Mime = data[a:b]

	// read data size
	a = b
	b += 4
	needle.DataSize = util.BytesToUint(data[a:b])

	// read data content
	a = b
	b += int(needle.DataSize)
	needle.Data = data[a:b]

	// read the checksum
	a = b
	b += 4
	needle.Checksum = util.BytesToUint(data[a:b])

	return needle
}

func writeToFile(volume string, byteValues []byte) {
	file, err := os.OpenFile(volume, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err = file.Write(byteValues)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func readFile(path string) (Needle, error) {
	file, err := os.Open(path)
	if err != nil {
		return Needle{}, err
	}

	fmt.Println(filepath.Base(file.Name()), path)

	data, err := io.ReadAll(file)
	if err != nil {
		return Needle{}, err
	}

	dataSize := len(data)

	fileName := filepath.Base(file.Name())
	nameBytes := []byte(fileName)
	nameSize := len(nameBytes)

	uintSize := int(unsafe.Sizeof(uint32(1)))

	needle := Needle{
		TotalSize: uint32(nameSize + dataSize + uintSize*3),
		NameSize:  uint32(nameSize),
		Name:      nameBytes,
		DataSize:  uint32(dataSize),
		Data:      data,
	}
	return needle, nil
}
