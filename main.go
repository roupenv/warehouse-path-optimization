package main

import (
	"fmt"
	"github.com/yourbasic/graph"
)

func main() {
	var warehouse warehouse
	warehouse.initialize(warehouseConfig)
	warehouse.getAllEdges()
	warehouseGraph := initializeGraph(warehouse.getAllEdges())
	fmt.Println(warehouseGraph)
	// warehouse.workers[0].Move(6,warehouseGraph, &warehouse)
	// warehouse.workers[0].Move(15,warehouseGraph, &warehouse)
	warehouse.workers[0].getProduct("AD", 10, warehouseGraph, &warehouse)
	// warehouse.workers[0].GetProduct("A", warehouseGraph, &warehouse)
	// warehouse.workers[0].GetProduct("AD", warehouseGraph, &warehouse)

	fmt.Println("Current Location", warehouse.workers[0].currentLocation)

}

func initializeGraph(graphNodes warehouseNodes) *graph.Immutable {
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
