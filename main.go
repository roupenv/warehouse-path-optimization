package main

// import "fmt"

func main() {
	var warehouse warehouse
	warehouse.initialize(warehouseConfig)
	worker := warehouse.workers[0]
	worker.move(&warehouse, 0)
	worker.move(&warehouse, 1)

	// worker.getProduct(&warehouse, "AD", 10)
	// worker.getProduct(&warehouse,"A", 10)
	// worker.getCurrentLocation(&warehouse)

	// newOrder := order{
	// 	lineItem{ "A", 20},
	// 	lineItem{ "B", 20},
	// 	lineItem{ "C", 20},
	// 	lineItem{ "D", 20},
	// 	lineItem{ "E", 20},
	// 	lineItem{ "F", 20},
	// 	lineItem{ "G", 20},
	// }

	// var sim sim
	// sim.startWork(worker, newOrder)
	// sim.incrementClock()

	// fmt.Println(sim.clock)
}

