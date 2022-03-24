package main

import "github.com/yourbasic/graph"

func initializeGraph(graphNodes warehouseGraph) *graph.Immutable {
	numberVertices := (len(graphNodes.edges) + 1) + (len(graphNodes.neighbors))
	g := graph.New(numberVertices)
	for v := 0; v < len(graphNodes.edges); v++ {
		g.AddBothCost(v, v+1, graphNodes.edges[v])
	}
	for v := 0; v < len(graphNodes.neighbors); v++ {
		g.AddBoth(int(graphNodes.neighbors[v].junctionNode), graphNodes.neighbors[v].id)
	}

	// Create Immutable Copy
	immutableG := graph.Sort(g)
	return immutableG
}