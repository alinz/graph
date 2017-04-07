package graph

// Vertex is a simple interface which contains Value and Edges
type Vertex interface {
	Name() ([]byte, error)
	SetName([]byte) error
	Value() ([]byte, error)
	SetValue([]byte) error
	Edges() ([]Edge, error)
}

// Edge is a simple interface which represents an edge
type Edge interface {
	Name() ([]byte, error)
	SetName([]byte) error
	Value() ([]byte, error)
	SetValue([]byte) error
	Vertex() Vertex
}

// Graph is the basic interface represents the graph,
// graph is backed by a reliable backend such as Boltdb or ...
type Graph interface {
	// Vertex does 2 things, find an existing Vertex with the same value or create one oterwise
	Vertex([]byte) (Vertex, error)
	// Edge creates a brand new Edge
	Edge([]byte) (Edge, error)

	// Connect tries to connect 2 vertices with given edge
	// for bidirectional connection this method must be called twice with rotating vertices
	Connect(Vertex, Edge, Vertex) error
	// RemoveEdge because edge needs to be removed from Vertex in backend, it requires vertex to be
	// pass as first argument, so backend can safely remove the edge
	RemoveEdge(Edge)
	// RemoveVertex will remove Vertex from graph and remove all edges to and from this vertex
	RemoveVertex(Vertex)
	// Close closes and cleans up the underneath data
	Close() error
}
