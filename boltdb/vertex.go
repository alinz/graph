package boltdb

import (
	"github.com/alinz/graph"
	"github.com/alinz/graph/boltdb/names"
	"github.com/boltdb/bolt"
)

type boltVertex struct {
	id []byte
}

func (v *boltVertex) Name() ([]byte, error) {
	return nil, nil
}

func (v *boltVertex) Value() ([]byte, error) {
	return nil, nil
}

func (v *boltVertex) SetValue(value []byte) error {
	return nil
}

func (v *boltVertex) Connect(vertex graph.Vertex, edge graph.Edge) error {
	return nil
}

func (v *boltVertex) Edges(name []byte) ([]byte, error) {
	return nil, nil
}

func (v *boltVertex) RemoveEdge(edge graph.Edge) error {
	return nil
}

func createVertex(name []byte, tx *bolt.Tx) (graph.Vertex, error) {
	vertexIndexBucket, err := tx.CreateBucketIfNotExists(names.VertexIndex)
	if err != nil {
		return nil, err
	}

	vertexID := vertexIndexBucket.Get(name)
	if vertexID != nil {
		return nil, graph.ErrDuplicateName
	}

	verticiesBucket, err := tx.CreateBucketIfNotExists(names.Verticies)
	if err != nil {
		return nil, err
	}

	vertexID, err = generateBucketID(verticiesBucket)
	if err != nil {
		return nil, err
	}

	vertexBucket, err := verticiesBucket.CreateBucket(vertexID)
	if err != nil {
		return nil, err
	}

	err = vertexBucket.Put(names.VertexName, name)
	if err != nil {
		return nil, err
	}

	return &boltVertex{
		id: vertexID,
	}, nil
}

func getVertex(name []byte, tx *bolt.Tx) (graph.Vertex, error) {
	vertexIndexBucket := tx.Bucket(names.VertexIndex)
	if vertexIndexBucket == nil {
		return nil, graph.ErrNotFound
	}

	vertexID := vertexIndexBucket.Get(name)
	if vertexID == nil {
		return nil, graph.ErrNotFound
	}

	return &boltVertex{
		id: vertexID,
	}, nil
}

// getVertexName returns the name of vertex
// it name is a required field. if the return value is nil, it means that
// somehow, data got crrupted.
//
// Note: the return value []byte is only valid for the duration of Tx. make
// sure to copy the value using `bytesCopy`
func getVertexName(id []byte, tx *bolt.Tx) ([]byte, error) {
	verticiesBucket := tx.Bucket(names.Verticies)
	if verticiesBucket == nil {
		return nil, graph.ErrNotFound
	}

	vertexBucket := verticiesBucket.Bucket(id)
	if vertexBucket == nil {
		return nil, graph.ErrNotFound
	}

	return vertexBucket.Get(names.VertexName), nil
}

// removeVertex it removes the vertex fully from graph
func removeVertex(id []byte, tx *bolt.Tx) error {
	return nil
}
