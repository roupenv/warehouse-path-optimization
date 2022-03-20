package main

type workArea struct {
	numberItems int
}

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

type warehouse struct {
	junctions        map[junctionNode]junctionData
	workers          []worker
	workAreaLocation junctionNode
	productLocs      map[string]int
	aisleLocs        map[int]area
}

func (w *warehouse) initialize(init warehouse) {
	w.junctions = init.junctions
	w.workers = init.workers
	w.workAreaLocation = init.workAreaLocation
	w.productLocations(init)
	w.aisleLocations(init)
}

func (w warehouse) getEdgeDistances() []int64 {
	edgeDistances := make([]int64, len(w.junctions))

	for i := range w.junctions {
		edgeDistances[i] = w.junctions[i].distanceToNextJunction
	}
	//Remove Last Empty Vertex from Slice
	edgeDistances = edgeDistances[:len(edgeDistances)-1]
	return edgeDistances
}

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

type warehouseNodes struct {
	edges     []int64
	neighbors []neighborNode
}

func (w warehouse) getAllEdges() warehouseNodes {
	edgeDistance := w.getEdgeDistances()
	neighborNodes := w.getNeighborNodes()

	return warehouseNodes{edges: edgeDistance, neighbors: neighborNodes}
}