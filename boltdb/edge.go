package boltdb

var (
	edgeNameKey  = []byte("name")
	edgeValueKey = []byte("value")
	edgeVertexA  = []byte("vertex_a")
	edgeVertexB  = []byte("vertex_b")
)

type boltEdge struct {
}

// SetValue it sets value to Edge
func (e *boltEdge) SetValue(value []byte) {

}

// Value it returns the value of target edge
func (e *boltEdge) Value() []byte {
	return nil
}
