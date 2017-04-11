package boltdb

import (
	"github.com/alinz/graph"
	"github.com/boltdb/bolt"
)

var (
	vertexValueKey = []byte("value")
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
	return nil, nil
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
