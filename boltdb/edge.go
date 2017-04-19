package boltdb

import (
	"fmt"

	"github.com/alinz/graph"
	"github.com/boltdb/bolt"
)

type boltEdge struct {
	boltEntity
}

func (e *boltEdge) setName(name []byte) error {
	return e.db.View(func(tx *bolt.Tx) error {
		edgesBucket, err := getNestedBuckets(tx, keys.edges)
		if err != nil {
			return err
		}

		id, err := generateBucketID(edgesBucket)
		if err != nil {
			return err
		}

		edgeBucket, err := edgesBucket.CreateBucket(id)
		if err != nil {
			return err
		}

		return edgeBucket.Put(keys.name, name)
	})
}

func (e *boltEdge) Prev() (graph.Vertex, error) {
	var vertex graph.Vertex

	err := e.db.View(func(tx *bolt.Tx) error {
		vertexID, err := getValueFromNestedBuckets(tx, keys.prev, keys.edges, e.id)
		if err != nil {
			return err
		}

		if vertexID == nil {
			return fmt.Errorf("eof")
		}

		vertex = newVertex(e.db, vertexID)
		return nil
	})

	return vertex, err
}

func (e *boltEdge) Next() (graph.Vertex, error) {
	var vertex graph.Vertex

	err := e.db.View(func(tx *bolt.Tx) error {
		vertexID, err := getValueFromNestedBuckets(tx, keys.next, keys.edges, e.id)
		if err != nil {
			return err
		}

		if vertexID == nil {
			return fmt.Errorf("eof")
		}

		vertex = newVertex(e.db, vertexID)
		return nil
	})

	return vertex, err
}

func newEdge(db *bolt.DB, id, name []byte) (graph.Edge, error) {
	edge := &boltEdge{
		boltEntity{
			id:   nil,
			kind: edgeType,
			db:   db,
		},
	}

	err := edge.setName(name)
	if err != nil {
		return nil, err
	}

	return edge, nil
}
