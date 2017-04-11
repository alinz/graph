package boltdb

import (
	"github.com/alinz/graph"
	"github.com/boltdb/bolt"
)

var (
	vertexValueKey = []byte("value")
	vertexEdgesKey = []byte("edges")
)

type boltVertex struct {
	id []byte
	db *bolt.DB
}

// SetValue this sets value to target vertex
func (v *boltVertex) SetValue(value []byte) error {
	return v.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(v.id)
		if bucket == nil {
			return graph.ErrVertexNotFound
		}

		if err := bucket.Put(vertexValueKey, value); err != nil {
			return err
		}

		return nil
	})
}

// Value it returns the value of vertex
func (v *boltVertex) Value() ([]byte, error) {
	var value []byte

	err := v.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(v.id)
		if bucket == nil {
			return graph.ErrVertexNotFound
		}

		src := bucket.Get(vertexValueKey)
		if src == nil {
			// This should not happen
			return nil
		}

		// because src is only valid inside Update transaction, we need to copy it
		// to a new value
		value = make([]byte, len(src))
		copy(value, src)

		return nil
	})

	return value, err
}

// Connects it connects a -[ e ]-> b with edge e
//
// errors that might happen
// - source vertex already has a connection with same label
func (v *boltVertex) Connects(vertex graph.Vertex) (graph.Edge, error) {
	var edge graph.Edge

	targetVertex, ok := vertex.(*boltVertex)
	if !ok {
		return nil, graph.ErrNotSameType
	}

	err := v.db.Update(func(tx *bolt.Tx) error {
		// we need to make sure that the current vertex exists
		srcBucket := tx.Bucket(v.id)
		if srcBucket == nil {
			return graph.ErrVertexNotFound
		}

		// we need to make sure that the target vertex exists
		desBucket := tx.Bucket(targetVertex.id)
		if desBucket == nil {
			return graph.ErrVertexNotFound
		}

		// we get access to edges bucket of src vertex.
		// src edges bucket contains a seq id as key and value points to global seq bucket
		// which contains everything related to this particulat edge
		srcEdgesBucket, err := srcBucket.CreateBucketIfNotExists(vertexEdgesKey)
		if err != nil {
			return err
		}

		edges := tx.Bucket(globalEdgesBucketName)

		// we are using sequence id to reference this vertex
		// everywhere. It uses less memory in average, 8 bytes
		id, err := edges.NextSequence()
		if err != nil {
			return err
		}

		idInBytes := uint64tobyte(id)

		edgeBucket, err := edges.CreateBucket(idInBytes)
		if err != nil {
			return err
		}

		err = edgeBucket.Put(edgeVertexA, v.id)
		if err != nil {
			return err
		}

		err = edgeBucket.Put(edgeVertexB, targetVertex.id)
		if err != nil {
			return err
		}

		id, err = srcEdgesBucket.NextSequence()
		if err != nil {
			return err
		}

		err = srcEdgesBucket.Put(uint64tobyte(id), idInBytes)
		if err != nil {
			return err
		}

		return nil
	})

	return edge, err
}

// Edges it returns all edges
func (v *boltVertex) Edges(label []byte) ([]graph.Edge, error) {
	return nil, nil
}

// RemoveEdge it removes given edge from vertex
//
// errors that might happen
// - vertex does not have reqsuetd edge
func (v *boltVertex) RemoveEdge(e graph.Edge) error {
	return nil
}
