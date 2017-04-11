package graph

import (
	"fmt"
)

var (
	// ErrDuplicateEdge it happens when an edge tries to be added to a vertex which already has the same edge
	// each edge must have a label. this lable will unique per 2 connected verties.
	ErrDuplicateEdge = fmt.Errorf("source vertex already has a connection with same label")
	//ErrEdgeNotFound if an edge with given label does not exist
	ErrEdgeNotFound = fmt.Errorf("vertex does not have reqsuetd edge")
	// ErrDuplicateVertex each graph can't have the same vertex with the same label
	ErrDuplicateVertex = fmt.Errorf("vertex with given name already exists in this graph")
	// ErrVertexNotFound if the given vertex is not part of that graph
	ErrVertexNotFound = fmt.Errorf("vertex does not found in this graph")
	// ErrNotSameType if the given vertex/edge created from different backend.
	// for example, you can not pass the non-boltdb backend edge and vertex to
	// boltdb graph.
	ErrNotSameType = fmt.Errorf("not the same type as backend")
)

// Edge this is base structure for Edge
type Edge interface {
	// SetValue it sets value to Edge
	SetValue(value []byte)

	// Value it returns the value of target edge
	Value() []byte
}

// Vertex this is a base structure for vertex
type Vertex interface {
	// SetValue this sets value to target vertex
	SetValue(value []byte) error

	// Value it returns the value of vertex
	Value() ([]byte, error)

	// Connects it connects a -[ e ]-> b with edge e
	//
	// errors that might happen
	// - source vertex already has a connection with same label
	// - not the same type as backend
	Connects(vertex Vertex) (Edge, error)

	// Edges it returns all edges
	Edges(label []byte) ([]Edge, error)

	// RemoveEdge it removes given edge from vertex
	//
	// errors that might happen
	// - vertex does not have reqsuetd edge
	// - not the same type as backend
	RemoveEdge(e Edge) error
}

// SubGraph this is a base structure for graph
type SubGraph interface {
	// AddVertex this adds a vertex to Graph.
	//
	// errors that might happen
	// - vertex with given name already exists in this graph
	// - not the same type as backend
	AddVertex(vertex Vertex) error

	// Vertex this finds vertex by given name in the graph
	//
	// errors that might happen
	// - give name doesn not found in this graph
	Vertex(name []byte) (Vertex, error)

	// RemoveVertex removes vertex from target graph
	//
	// errors that might happen
	// - vertex does not found in this graph
	// - not the same type as backend
	RemoveVertex(vertex Vertex) error
}

// Graph this is a base structure for Graph
type Graph interface {
	// SubGraph will create a logical sub graph. The main use case of sub graph is
	// to label verticies which can be query and find faster.
	//
	// errors that might happen
	// - mostly the backend implementation error
	CreateSubGraph(label []byte) (SubGraph, error)

	// CreateVertex all the new vertex must be created here. the given name must
	// be unique.
	//
	// errors that might happen
	// - duplicate vertex name found
	CreateVertex(name []byte) (Vertex, error)
}
