package boltdb

import (
	"fmt"

	"github.com/boltdb/bolt"
)

type boltEntity struct {
	id   []byte
	kind boltType
	db   *bolt.DB
}

func (e *boltEntity) baseBucketName() ([]byte, error) {
	switch e.kind {
	case edgeType:
		return keys.edges, nil
	case vertexType:
		return keys.verticies, nil
	case subGraphType:
		return keys.subGraphs, nil
	default:
		return nil, fmt.Errorf("internal type error")
	}
}

func (e *boltEntity) Name() ([]byte, error) {
	base, err := e.baseBucketName()
	if err != nil {
		return nil, err
	}

	var value []byte

	err = e.db.View(func(tx *bolt.Tx) error {
		data, err := getValueFromNestedBuckets(tx, keys.name, base, e.id, keys.name)
		value = cloneBytes(data)

		return err
	})

	return value, err
}

func (e *boltEntity) Value() ([]byte, error) {
	base, err := e.baseBucketName()
	if err != nil {
		return nil, err
	}

	var value []byte

	err = e.db.View(func(tx *bolt.Tx) error {
		data, err := getValueFromNestedBuckets(tx, keys.name, base, e.id, keys.value)
		value = cloneBytes(data)

		return err
	})

	return value, err
}

func (e *boltEntity) SetValue(value []byte) error {
	base, err := e.baseBucketName()
	if err != nil {
		return err
	}

	err = e.db.Update(func(tx *bolt.Tx) error {
		return setValueToNestedBuckets(tx, keys.value, value, base, e.id)
	})

	return err
}
