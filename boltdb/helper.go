package boltdb

import (
	"encoding/binary"
	"fmt"

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

func cloneBytes(value []byte) []byte {
	if value == nil {
		return nil
	}

	valueCopy := make([]byte, len(value))
	copy(valueCopy, value)
	return valueCopy
}

func getNestedBuckets(tx *bolt.Tx, buckets ...[]byte) (*bolt.Bucket, error) {
	var targetBucket *bolt.Bucket

	for _, bucket := range buckets {
		if targetBucket == nil {
			targetBucket = tx.Bucket(bucket)
		} else {
			targetBucket = targetBucket.Bucket(bucket)
		}

		if targetBucket == nil {
			return nil, fmt.Errorf("bucket '%s' not found", bucket)
		}
	}

	return targetBucket, nil
}

// getValueFromNestedBuckets NOTE: the return bytes is only available inside Tx. make sure to clone it
func getValueFromNestedBuckets(tx *bolt.Tx, key []byte, buckets ...[]byte) ([]byte, error) {
	bucket, err := getNestedBuckets(tx, buckets...)
	if err != nil {
		return nil, err
	}

	value := bucket.Get(key)

	return value, nil
}

func setValueToNestedBuckets(tx *bolt.Tx, key, value []byte, buckets ...[]byte) error {
	bucket, err := getNestedBuckets(tx, buckets...)
	if err != nil {
		return err
	}

	return bucket.Put(key, value)
}
