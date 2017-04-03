package bolt

import (
	"encoding/binary"
	"log"

	"github.com/alinz/graph"
	boltdb "github.com/boltdb/bolt"
)

/*

	Graph {
		ValueToVertex: {
			<value>: <Vertex ID>
		}

		EdgeValueToVerticies: {
			<value>: {
				<seq: ID>: <Vertex ID>
			}
		}

		Verticies: {
			<seq: Vertex ID>: {
				Value: <value>
				Edges: {
					<seq: ID>: <seq: Edge ID>
				}
			}
		}

		Edges: {
			<seq: Edge ID>: {
				Value: <value>
				VertexID: <seq: Vertex ID>
			}
		}
	}

*/

var (
	valueToVertexBucketName        = []byte("valueToVertexBucketName")
	edgeValueToVerticiesBucketName = []byte("edgeValueToVerticiesBucketName")
	verticiesBucketName            = []byte("verticiesBucketName")
	edgesBucketName                = []byte("edgesBucketName")
)

func uint64Bytes(value uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(value))
	return b
}

// #
// # Graph implmentation in bolt

type boltGraph struct {
	db *boltdb.DB
}

func (bg *boltGraph) Vertex(value []byte) graph.Vertex {
	var vertex graph.Vertex

	bg.db.Update(func(tx *boltdb.Tx) error {
		return nil
	})

	return vertex
}

func (bg *boltGraph) Edge(value []byte) graph.Edge {
	var edge graph.Edge

	edge = &boltEdge{
		id:    nil,
		value: value,
	}

	return edge
}

func (bg *boltGraph) Connect(a graph.Vertex, e graph.Edge, b graph.Vertex) error {

	bg.db.Update(func(tx *boltdb.Tx) error {
		return nil
	})

	return nil
}

func (bg *boltGraph) RemoveEdge(v graph.Vertex, e graph.Edge) {

}

func (bg *boltGraph) RemoveVertex(v graph.Vertex) {

}

func (bg *boltGraph) Close() error {
	return bg.db.Close()
}

// NewBoltGraph creates a bolt backend for graph
func NewBoltGraph(path string) graph.Graph {
	db, err := boltdb.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// we need to create 2 buckets, one for vertices and one for edges
	err = db.Batch(func(tx *boltdb.Tx) error {
		_, err := tx.CreateBucketIfNotExists(valueToVertexBucketName)
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists(edgeValueToVerticiesBucketName)
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists(verticiesBucketName)
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists(edgesBucketName)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return &boltGraph{
		db: db,
	}
}
