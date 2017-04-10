package boltdb

import "github.com/alinz/graph"

type boltGraph struct {
}

func (g *boltGraph) SubGraph(label []byte) graph.SubGraph {
	return nil
}

// New it initialize the graph backed by boltdb. everytime this function is called,
// brand new bolt db will be created.
func New(path string) graph.Graph {
	return &boltGraph{}
}
