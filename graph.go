package graph

import "fmt"

var (
	ErrAlreadyConnected = fmt.Errorf("vertices already connected with the same edge")
	ErrNotFound         = fmt.Errorf("not found")
)

// Vertex is a simple interface which contains Value and Edges
type Vertex interface {
	Value() []byte
	Edges() []Edge
}

// Edge is a simple interface which represents an edge
type Edge interface {
	Value() []byte
	Vertex() Vertex
}

// Graph is the basic interface represents the graph,
// graph is backed by a reliable backend such as Boltdb or ...
type Graph interface {
	// Vertex does 2 things, find an existing Vertex with the same value or create one oterwise
	Vertex(value []byte) Vertex
	// Edge creates a brand new Edge
	Edge(value []byte) Edge
	// Connect tries to connect 2 vertices with given edge
	// for bidirectional connection this method must be called twice with rotating vertices
	Connect(a Vertex, e Edge, b Vertex) error
	// RemoveEdge because edge needs to be removed from Vertex in backend, it requires vertex to be
	// pass as first argument, so backend can safely remove the edge
	RemoveEdge(v Vertex, e Edge)
	// RemoveVertex will remove Vertex from graph and remove all edges to and from this vertex
	RemoveVertex(v Vertex)
	// Close closes and cleans up the underneath data
	Close() error
}

// we need to add Vertex to graph
// we need to remove Vertex from graph
// we need to connect 2 Vertices

// we need to add data to Vertex and Edge
// we need to be able to read from it
