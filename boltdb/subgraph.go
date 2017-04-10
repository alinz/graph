package boltdb

import "github.com/alinz/graph"

type boltSubGraph struct {
}

// AddVertex this adds a vertex to Graph.
//
// errors that might happen
// - vertex with given name already exists in this graph
func (s *boltSubGraph) AddVertex(vertex graph.Vertex) error {
	return nil
}

// Vertex this finds vertex by given name in the graph
//
// errors that might happen
// - give name doesn not found in this graph
func (s *boltSubGraph) Vertex(name []byte) (graph.Vertex, error) {
	return nil, nil
}

// RemoveVertex removes vertex from target graph
//
// errors that might happen
// - vertex does not found in this graph
func (s *boltSubGraph) RemoveVertex(vertex graph.Vertex) error {
	return nil
}
