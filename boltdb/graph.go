package boltdb

import (
	"github.com/alinz/graph"
	"github.com/boltdb/bolt"
)

var (
	globalVericiesBucketName = []byte("verticies_bucket_name")
	globalEdgesBucketName    = []byte("edges_bucket_name")
)

type boltGraph struct {
	db *bolt.DB
}

// SubGraph it creates or returns SubGraph with the given lable
// behind the scene, it creates a unique bucket for that given label
func (g *boltGraph) CreateSubGraph(label []byte) (graph.SubGraph, error) {
	var subGraph graph.SubGraph

	err := g.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(label)
		if err == nil {
			subGraph = &boltSubGraph{
				db: bucket,
			}
		}

		return err
	})

	return subGraph, err
}

func (g *boltGraph) CreateVertex(name []byte) (graph.Vertex, error) {
	var vertex graph.Vertex

	err := g.db.Update(func(tx *bolt.Tx) error {
		verticies := tx.Bucket(globalVericiesBucketName)

		// check if name is unique
		if verticies.Get(name) != nil {
			return graph.ErrDuplicateVertex
		}

		// we are using sequence id to reference this vertex
		// everywhere. It uses less memory in average, 8 bytes
		id, err := verticies.NextSequence()
		if err != nil {
			return err
		}

		idInBytes := uint64tobyte(id)

		// we are creating a bucket with vertex_id as name
		// all the information about that vertex will be save inside
		// this bucket
		_, err = tx.CreateBucket(idInBytes)
		if err != nil {
			return err
		}

		// verticies bucket (name -> vertex_id)
		err = verticies.Put(name, idInBytes)
		if err != nil {
			return err
		}

		vertex = &boltVertex{
			idInBytes,
			g.db,
		}

		return nil
	})

	return vertex, err
}

// New it initialize the graph backed by boltdb. everytime this function is called,
// brand new bolt db will be created.
func New(path string) (graph.Graph, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	// this update creates couple of global buckets
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(globalVericiesBucketName)
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists(globalEdgesBucketName)
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return &boltGraph{
		db,
	}, err
}
