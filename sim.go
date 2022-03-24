package main

import (
	"fmt"
	"github.com/yourbasic/graph"
)

type sim struct {
	clock int
}

func generatePath(warehouse *warehouse, currentLocation int, finalLocation int) []int {
	path, _ := graph.ShortestPath(warehouse.wGraph, currentLocation, finalLocation)
	return path
}

func (s *sim) incrementClock() {
	s.clock++
}

func (s *sim) startWork(wo worker, order order) {
	if wo.status == idle {
		fmt.Println(order)
	}

}


type event struct {
}


type process struct {
	events []event
	worker    *worker
	warehouse *warehouse
}

func (p *process) init(warehouse *warehouse, worker *worker) {
	p.worker = worker
	p.warehouse = warehouse
}

func (p *process) start() {
	p.worker.status = moving

}

func (p *process) next(nextLoc int) {
	p.worker.move(p.warehouse, nextLoc)
}

func (p *process) skip() {
}

func (p *process) finished() {
	p.worker.status = idle

}
