package boltdb

import "github.com/alinz/graph"

type boltVertex struct {
	id []byte
}

// SetValue this sets value to target vertex
func (v *boltVertex) SetValue(value []byte) {

}

// Value it returns the value of vertex
func (v *boltVertex) Value() []byte {
	return nil
}

// Connects it connects a -[ e ]-> b with edge e
//
// errors that might happen
// - source vertex already has a connection with same label
func (v *boltVertex) Connects(vertex graph.Vertex) (graph.Edge, error) {
	return nil, nil
}

// Edges it returns all edges
func (v *boltVertex) Edges(label []byte) []graph.Edge {
	return nil
}

// RemoveEdge it removes given edge from vertex
//
// errors that might happen
// - vertex does not have reqsuetd edge
func (v *boltVertex) RemoveEdge(e graph.Edge) error {
	return nil
}
