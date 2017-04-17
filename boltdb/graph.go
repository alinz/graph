package boltdb

import (
	"github.com/alinz/graph"
	"github.com/boltdb/bolt"
)

type boltGraph struct {
	db *bolt.DB
}

func (g *boltGraph) SubGraph(name []byte) (graph.SubGraph, error) {
	var subgraph graph.SubGraph

	err := g.db.Update(func(tx *bolt.Tx) error {
		result, err := createSubGraphIfNotExists(name, tx)
		subgraph = result
		return err
	})

	return subgraph, err
}

func (g *boltGraph) CreateVertex(name []byte) (graph.Vertex, error) {
	var vertex graph.Vertex
	var err error

	g.db.Update(func(tx *bolt.Tx) error {
		vertex, err = createVertex(name, tx)
		return err
	})

	return vertex, err
}

func (g *boltGraph) RemoveVertex(vertex graph.Vertex) error {
	nativeVertex, ok := vertex.(*boltVertex)
	if !ok {
		return graph.ErrNotNative
	}

	return g.db.Update(func(tx *bolt.Tx) error {
		return removeVertex(nativeVertex.id, tx)
	})
}

func (g *boltGraph) CreateEdge(name []byte) (graph.Edge, error) {
	return nil, nil
}

// NewGraph creates a new graph which backed by a single bolt file
func NewGraph(path string) (graph.Graph, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &boltGraph{
		db,
	}, nil
}
