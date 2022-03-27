package main

import (
	"fmt"

	"github.com/yourbasic/graph"
)

type area struct {
	aisle
	workArea
}
type junctionNode int

type junctionData struct {
	areas                  []area
	nextJunction           int
	distanceToNextJunction int64
}

type neighborNode struct {
	junctionNode
	id int
}
type warehouseGraph struct {
	edges     []int64
	neighbors []neighborNode
}

type warehouse struct {
	junctions     map[junctionNode]junctionData
	workArea      workArea
	productLocs   map[sku]int
	aisleLocs     map[int]area
	internalGraph *graph.Immutable
}


func newWarehouse(init warehouse) *warehouse {
	w := &warehouse{
		junctions: init.junctions,
		workArea:  init.workArea,
	}
	w.internalGraph = initializeGraph(warehouseGraph{w.getEdgeDistances(), w.getNeighborNodes()})
	w.productLocs = productLocations(init)
	w.aisleLocs = aisleLocations(init)
	fmt.Println("Warehouse Initialized")
	fmt.Println(w.internalGraph)
	fmt.Println()
	return w
}

// Returns the edge distances for the warehouse graph
func (w warehouse) getEdgeDistances() []int64 {
	edgeDistances := make([]int64, len(w.junctions))

	for i := range w.junctions {
		edgeDistances[i] = w.junctions[i].distanceToNextJunction
	}
	//Remove Last Empty Vertex from Slice
	edgeDistances = edgeDistances[:len(edgeDistances)-1]
	return edgeDistances
}

// Returns the neighbor nodes for the warehouse graph
func (w warehouse) getNeighborNodes() []neighborNode {
	id := len(w.junctions)

	var neighborNodes []neighborNode
	for _, junctionNode := range sortMap(w.junctions) {
		for i := 0; i < len(w.junctions[junctionNode].areas); i++ {
			neighborNodes = append(neighborNodes, neighborNode{
				junctionNode: junctionNode,
				id:           id,
			})
			id++
		}
	}
	return neighborNodes
}
