package boltdb

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

func uint64tobyte(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func generateBucketID(bucket *bolt.Bucket) ([]byte, error) {
	id, err := bucket.NextSequence()
	if err != nil {
		return nil, err
	}

	return uint64tobyte(id), nil
}

func bytesCopy(value []byte) []byte {
	if value == nil {
		return nil
	}

	valueCopy := make([]byte, len(value))
	copy(valueCopy, value)
	return valueCopy
}
