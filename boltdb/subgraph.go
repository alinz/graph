package boltdb

import (
	"github.com/alinz/graph"
	"github.com/alinz/graph/boltdb/names"
	"github.com/boltdb/bolt"
)

type boltSubGraph struct {
	id []byte
	db *bolt.DB
}

func (sg *boltSubGraph) Name() ([]byte, error) {
	var result []byte

	err := sg.db.View(func(tx *bolt.Tx) error {
		subGraphsBucket := tx.Bucket(names.SubGraphs)
		if subGraphsBucket == nil {
			return graph.ErrNotFound
		}

		subGraphBucket := subGraphsBucket.Bucket(sg.id)
		if subGraphBucket == nil {
			return graph.ErrNotFound
		}

		result = bytesCopy(subGraphBucket.Get(names.SubGraphName))

		return nil
	})

	return result, err
}

func (sg *boltSubGraph) Value() ([]byte, error) {
	var result []byte

	err := sg.db.View(func(tx *bolt.Tx) error {
		subGraphsBucket := tx.Bucket(names.SubGraphs)
		if subGraphsBucket == nil {
			return graph.ErrNotFound
		}

		subGraphBucket := subGraphsBucket.Bucket(sg.id)
		if subGraphBucket == nil {
			return graph.ErrNotFound
		}

		value := subGraphBucket.Get(names.SubGraphValue)
		result = make([]byte, len(value))
		copy(result, value)

		return nil
	})

	return result, err
}

func (sg *boltSubGraph) SetValue(value []byte) error {
	return sg.db.Update(func(tx *bolt.Tx) error {
		subGraphsBucket := tx.Bucket(names.SubGraphs)
		if subGraphsBucket == nil {
			return graph.ErrNotFound
		}

		subGraphBucket := subGraphsBucket.Bucket(sg.id)
		if subGraphBucket == nil {
			return graph.ErrNotFound
		}

		return subGraphBucket.Put(names.SubGraphValue, value)
	})
}

func (sg *boltSubGraph) AddVertex(vertex graph.Vertex) error {
	nativeVertex, ok := vertex.(*boltVertex)
	if !ok {
		return graph.ErrNotNative
	}

	return sg.db.Update(func(tx *bolt.Tx) error {
		subGraphsBucket := tx.Bucket(names.SubGraphs)
		if subGraphsBucket == nil {
			return graph.ErrNotFound
		}

		subGraphBucket, err := subGraphsBucket.CreateBucketIfNotExists(sg.id)
		if err != nil {
			return err
		}

		subGraphVertices, err := subGraphBucket.CreateBucketIfNotExists(names.SubGraphVertices)
		if err != nil {
			return err
		}

		vertexName, err := getVertexName(nativeVertex.id, tx)
		if err != nil {
			return err
		}

		err = subGraphVertices.Put(vertexName, nativeVertex.id)
		if err != nil {
			return err
		}

		verticiesBucket, err := tx.CreateBucketIfNotExists(names.Verticies)
		if err != nil {
			return err
		}

		vertexBucket := verticiesBucket.Bucket(nativeVertex.id)
		if vertexBucket == nil {
			return graph.ErrNotFound
		}

		vertexSubGraphsBucket, err := vertexBucket.CreateBucketIfNotExists(names.VertexSubGraphs)
		if err != nil {
			return err
		}

		subGraphName, err := getSubGraphName(sg.id, tx)
		if err != nil {
			return err
		}

		err = vertexSubGraphsBucket.Put(subGraphName, sg.id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (sg *boltSubGraph) Vertex(name []byte) (graph.Vertex, error) {
	var vertex graph.Vertex

	sg.db.Update(func(tx *bolt.Tx) error {
		subGraphsBucket, err := tx.CreateBucketIfNotExists(names.SubGraphs)
		if err != nil {
			return err
		}

		subGraphBucket := subGraphsBucket.Bucket(sg.id)
		if subGraphBucket == nil {
			return graph.ErrNotFound
		}

		subGraphVerticesBucket, err := subGraphBucket.CreateBucketIfNotExists(names.SubGraphVertices)
		if err != nil {
			return err
		}

		vertexID := subGraphVerticesBucket.Get(name)
		if vertexID == nil {
			return graph.ErrNotFound
		}

		vertex = &boltVertex{
			id: vertexID,
		}

		return nil
	})

	return vertex, nil
}

func (sg *boltSubGraph) RemoveVertex(vertex graph.Vertex) error {
	// nativeVertex, ok := vertex.(*boltVertex)
	// if !ok {
	// 	return graph.ErrNotNative
	// }

	return nil
}

func createSubGraphIfNotExists(name []byte, tx *bolt.Tx) (graph.SubGraph, error) {
	subGraphIndexBucket, err := tx.CreateBucketIfNotExists(names.SubGraphIndex)
	if err != nil {
		return nil, err
	}

	subGraphID := subGraphIndexBucket.Get(name)
	if subGraphID == nil {
		subGraphsBucket, err := tx.CreateBucketIfNotExists(names.SubGraphs)
		if err != nil {
			return nil, err
		}

		subGraphID, err = generateBucketID(subGraphsBucket)
		if err != nil {
			return nil, err
		}

		_, err = subGraphsBucket.CreateBucket(subGraphID)
		if err != nil {
			return nil, err
		}
	}

	return &boltSubGraph{
		id: subGraphID,
		db: tx.DB(),
	}, nil
}

// getSubGraphName it returns name for given id for subgraph
//
// Note: the return value []byte is only valid for the duration of Tx. make
// sure to copy the value using `bytesCopy`
func getSubGraphName(id []byte, tx *bolt.Tx) ([]byte, error) {
	subGraphsBucket := tx.Bucket(names.SubGraphs)
	if subGraphsBucket == nil {
		return nil, graph.ErrNotFound
	}

	subGraphBucket := subGraphsBucket.Bucket(id)
	if subGraphsBucket == nil {
		return nil, graph.ErrNotFound
	}

	return subGraphBucket.Get(names.SubGraphName), nil
}
