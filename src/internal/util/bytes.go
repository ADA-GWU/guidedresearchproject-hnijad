package util

import "encoding/binary"

func UintToBytes(value uint32) []byte {
	byteValue := make([]byte, 4)
	binary.LittleEndian.PutUint32(byteValue, value)
	return byteValue
}

func BytesToUint(bytes []byte) uint32 {
	return binary.LittleEndian.Uint32(bytes)
}
