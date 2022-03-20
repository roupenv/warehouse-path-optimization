package main

import (
	"fmt"
)

type aisle struct {
	occupied bool
	products []product
}

type aisleLocation map[int]area

func (w *warehouse) aisleLocations(init warehouse) {
	newAisleLocations := make(aisleLocation)
	id := len(init.junctions)

	for _, junctionNode := range sortMap(init.junctions) {
		for area := 0; area < len(init.junctions[junctionNode].areas); area++ {
			newAisleLocations[id] = init.junctions[junctionNode].areas[area]
			id++
		}
	}
	w.aisleLocs = newAisleLocations
}

func (w *warehouse) occupyAisle(aisleId int, workerId int) {
	aisleToUpdate := w.aisleLocs[aisleId]
	aisleToUpdate.aisle.occupied = true
	fmt.Printf("Aisle %v Occupied By Worker %v \n", aisleId, workerId)
}

func (w *warehouse) unoccupiedAisle(aisleId int, workerId int) {
	aisleToUpdate := w.aisleLocs[aisleId]
	aisleToUpdate.aisle.occupied = false
	fmt.Printf("Aisle %v No Longer Occupied by Worker %v \n", aisleId, workerId)

}
