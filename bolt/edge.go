package bolt

import "github.com/alinz/graph"

// #
// # Edge implmentation in bolt

type boltEdge struct {
	id    []byte
	value []byte
}

func (be *boltEdge) Value() []byte {
	return nil
}

func (be *boltEdge) Vertex() graph.Vertex {
	return nil
}
