## Desing doc of Graph implementation in boltdb

```
SubGraphIdByName: {
  <name>: <subgraph.id>
  ...
}

VertexIdByName: {
  <name>: <vertex.id>
}

EdgeIdsByName: {
  <name>: {
    <seq:_>: <edge.id>
    ...
  }
}

SubGraphs: {
  <seq:subgraph.id>: {
    Name: []byte
    Value: []byte]
    Vertices: {
      <vertex.name>: <vertex.id>
      ...
    }
  }
}

Verticies: {
  <seq:vertex.id>: {
    Name: []byte
    Value: []byte
    SubGraphs: {
      <subgraph.name>: <subgraph.id>
      ...
    },
    Edges: {
      <seq:_>: <edge.id>
      ...
    }
  }
}

Edges: {
  <seq:edge.id>: {
    Name: []byte
    Value: []byte
    Prev: <vertex.id>
    Next: <vertex.id>
  }
}
```
