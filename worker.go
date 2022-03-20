package main

import (
	"fmt"

	"github.com/yourbasic/graph"
)

type dolly struct {
	capacity int
}

type workerStatus int

const (
	idle workerStatus = iota
	moving
	picking
)

type worker struct {
	id                    int
	dolly                 dolly
	status                workerStatus
	currentLocation       junctionNode
	totalDistanceTraveled int64
	totalTrips            int
	currentItems          int
}

func (wo *worker) move(newLocation int, g *graph.Immutable, warehouse *warehouse) {
	path, dist := graph.ShortestPath(g, int(wo.currentLocation), newLocation)
	endOfPath := path[len(path)-1]
	for i := 0; i < len(path); i++ {
		// fmt.Println(path[i])
	}
	fmt.Println("Distance Traveled", dist)
	warehouse.unoccupiedAisle(int(wo.currentLocation), wo.id)
	wo.currentLocation = junctionNode(endOfPath)
	warehouse.occupyAisle(endOfPath, wo.id)
	wo.totalDistanceTraveled += dist
	fmt.Println("Total Distance Traveled", wo.totalDistanceTraveled)

}

func (wo *worker) goHome(g *graph.Immutable, warehouse *warehouse) {
	path, dist := graph.ShortestPath(g, int(wo.currentLocation), int(warehouse.workAreaLocation))
	endOfPath := path[len(path)-1]
	for i := 0; i < len(path); i++ {
		// fmt.Println(path[i])
	}
	fmt.Println("Going Back to Work Area! Distance:", dist)
	warehouse.unoccupiedAisle(int(wo.currentLocation), wo.id)
	wo.currentLocation = junctionNode(endOfPath)
	warehouse.occupyAisle(endOfPath, wo.id)

	if wo.currentItems > 0 {
		fmt.Println("Emptying Items")
		wo.currentItems = 0
	}
	wo.totalTrips++

	wo.totalDistanceTraveled += dist
	fmt.Println("Total Distance Traveled", wo.totalDistanceTraveled)
}

func (wo *worker) getProduct(sku string, quantity int, g *graph.Immutable, warehouse *warehouse) {
	aisleLoc := warehouse.productLocs[sku]

	wo.move(aisleLoc, g, warehouse)
	if wo.currentItems <= wo.dolly.capacity {
		wo.currentItems = quantity
	}
	fmt.Println("Current Items on Dolly", wo.currentItems)

	wo.goHome(g, warehouse)
}
