package boltdb

type boltType int

const (
	_ boltType = iota
	edgeType
	vertexType
	subGraphType
)

var keys = struct {
	subGraphs      []byte
	verticies      []byte
	edges          []byte
	name           []byte
	value          []byte
	prev           []byte
	next           []byte
	indexSubGraphs []byte
	indexVerticies []byte
	indexEdges     []byte
}{
	subGraphs:      []byte("a"),
	verticies:      []byte("b"),
	edges:          []byte("c"),
	name:           []byte("d"),
	value:          []byte("e"),
	prev:           []byte("f"),
	next:           []byte("g"),
	indexSubGraphs: []byte("h"),
	indexVerticies: []byte("i"),
	indexEdges:     []byte("j"),
}
