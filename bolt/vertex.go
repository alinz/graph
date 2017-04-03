package bolt

import "github.com/alinz/graph"

// #
// # Vertex implmentation in bolt

type boltVertex struct {
	id    []byte
	value []byte
}

func (bv *boltVertex) Value() []byte {
	return bv.value
}

func (bv *boltVertex) Edges() []graph.Edge {
	return nil
}
