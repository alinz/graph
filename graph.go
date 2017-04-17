package graph

import (
	"fmt"
)

var (
	// ErrNotImplemented feature is not implemented
	ErrNotImplemented = fmt.Errorf("not implemented")
	// ErrDuplicateName ...
	ErrDuplicateName = fmt.Errorf("duplicate name")
	// ErrNotFound when an item is not found based on it's name
	ErrNotFound = fmt.Errorf("not found")
	// ErrNotNative if interface is not the internal object
	ErrNotNative = fmt.Errorf("not native")
)

type Entity interface {
	Name() ([]byte, error)
	Value() ([]byte, error)
	SetValue(value []byte) error
}

type Edge interface {
	Entity
	Prev() (Vertex, error)
	Next() (Vertex, error)
}

type Vertex interface {
	Entity
	Connect(vertex Vertex, edge Edge) error
	Edges(name []byte) ([]byte, error)
	RemoveEdge(edge Edge) error
}

type SubGraph interface {
	Entity
	AddVertex(vertex Vertex) error
	Vertex(name []byte) (Vertex, error)
	RemoveVertex(vertex Vertex) error
}

type Graph interface {
	SubGraph(name []byte) (SubGraph, error)
	CreateVertex(name []byte) (Vertex, error)
	RemoveVertex(vertex Vertex) error
	CreateEdge(name []byte) (Edge, error)
}
