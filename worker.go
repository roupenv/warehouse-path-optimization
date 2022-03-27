package main

import (
	"fmt"

	"github.com/yourbasic/graph"
)

type dolly struct {
	capacity int
}

type workerStats struct {
	timeWaiting           float64
	timeBusy              float64
	timeIdle              float64
	totalTrips            int
	totalDistanceTraveled int64
	totalUnitsPicked      int
}

type workerStatus int

const (
	idle workerStatus = iota
	waiting
	busy
	picking
)

type worker struct {
	id              int
	currentJob      job
	dolly           dolly
	status          workerStatus
	currentLocation junctionNode
	currentItems    int
	workerStats
	warehouse *warehouse
}

func newWorker(id int, capacity int, wh *warehouse) *worker {
	return &worker{
		id: id,
		dolly: dolly{
			capacity: capacity,
		},
		status:          idle,
		currentLocation: warehouseConfig.workArea.location,
		warehouse:       wh,
	}
}

func (wo worker) getStatus() string {
	return [...]string{"idle", "waiting", "busy", "picking"}[wo.status]
}

func (wo *worker) assignJob(j job, wP *workerPool) {
	wo.currentJob = j
	wo.status = busy
	fmt.Println("Assigning Job", wo.currentJob.id, "to worker", wo.id)

	wo.executeTask()

	//Finished Job
	wo.currentJob.status = complete
	wo.finishJob(wP)
}

func (wo *worker) executeTask() {
	//Get Product from Aisle
	wo.getProduct(wo.currentJob.sku, wo.currentJob.quantity)

	sleepRandomTime(100)

	//Then Go Home
	wo.goHome()

}

func (wo *worker) finishJob(wP *workerPool) {
	fmt.Println("Job Completed")

	wo.status = idle
	wo.currentJob = job{}

	wP.addWorkerBackToPool(wo)
}

func (wo *worker) move(newLocation int) {
	path, dist := graph.ShortestPath(wo.warehouse.internalGraph, int(wo.currentLocation), newLocation)
	endOfPath := path[len(path)-1]

	// for i := 0; i < len(path); i++ {
	// 	// fmt.Println(path[i])
	// }

	fmt.Println("Distance Traveled", dist)

	//Unoccupy Current Location
	wo.warehouse.unoccupiedAisle(int(wo.currentLocation), wo.id)

	//Set Workers Current location to end of path
	wo.currentLocation = junctionNode(endOfPath)

	//Occupy Current Location
	wo.warehouse.occupyAisle(endOfPath, wo.id)

	//Update Workers Total Distance Traveled
	wo.totalDistanceTraveled += dist
	wo.timeBusy += float64(dist) / float64(walkingPace)
	// fmt.Println("Total Distance Traveled", wo.totalDistanceTraveled)
}

func (wo *worker) pickItems(warehouse *warehouse, a chan aisle) {
	aisle := <-a
	if !aisle.occupied {
		fmt.Println("Picking Items")
		wo.status = picking
	} else if aisle.occupied {
		fmt.Println("Aisle is Occupied")
		wo.status = waiting
	}

}

func (wo *worker) goHome() {
	path, dist := graph.ShortestPath(wo.warehouse.internalGraph, int(wo.currentLocation), int(wo.warehouse.workArea.location))
	endOfPath := path[len(path)-1]

	// for i := 0; i < len(path); i++ {
	// 	// fmt.Println(path[i])
	// }

	fmt.Println("Going Back to Work Area! Distance:", dist)

	//Set Workers Current location to end of path
	wo.warehouse.unoccupiedAisle(int(wo.currentLocation), wo.id)

	//Set Workers Current location to end of path, which is Home
	wo.currentLocation = junctionNode(endOfPath)

	//Occupy Current Location
	wo.warehouse.occupyAisle(endOfPath, wo.id)

	if wo.currentItems > 0 {
		fmt.Println("Emptying Items")
		wo.currentItems = 0
	}

	//Update Workers Total Distance Traveled
	wo.totalDistanceTraveled += dist
	fmt.Println("Total Distance Traveled", wo.totalDistanceTraveled)

	//Update Worker Time Busy
	wo.timeBusy += float64(dist) / float64(walkingPace)

	//Increment Total Trips
	wo.totalTrips++

	//Increment work Area number units delivered
}

func (wo *worker) getProduct(item sku, quantity int) {
	aisleLoc := wo.warehouse.productLocs[item]

	wo.move(aisleLoc)

	//Update Workers Current Items After Move
	if wo.currentItems <= wo.dolly.capacity {
		wo.currentItems = quantity
	}
	fmt.Println("Current Items on Dolly", wo.currentItems)

	//Increment Total Units Picked
	wo.totalUnitsPicked += wo.currentItems

}

func (wo *worker) getCurrentLocation() {
	if wo.currentLocation == wo.warehouse.workArea.location {
		fmt.Println("Current Location is:", "Home")
	} else {
		fmt.Println("Current Location is:", wo.currentLocation)
	}
}
