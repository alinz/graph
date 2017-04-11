package boltdb

import (
	"crypto/rand"
	"encoding/binary"
)

func uuid() ([]byte, error) {
	b := make([]byte, 16)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func uint64tobyte(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
