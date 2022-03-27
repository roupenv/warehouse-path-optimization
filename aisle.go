package main

import (
	"fmt"
)

type aisle struct {
	occupied bool
	products []product
}

type aisleLocation map[int]area


func aisleLocations(init warehouse) aisleLocation{
	newAisleLocations := make(aisleLocation)
	id := len(init.junctions)

	for _, junctionNode := range sortMap(init.junctions) {
		for area := 0; area < len(init.junctions[junctionNode].areas); area++ {
			newAisleLocations[id] = init.junctions[junctionNode].areas[area]
			id++
		}
	}
	return newAisleLocations
}

func (w *warehouse) occupyAisle(aisleId int, workerId int) {
	aisleToUpdate := w.aisleLocs[aisleId]
	aisleToUpdate.aisle.occupied = true
	if aisleId == int(w.workArea.location) {
		fmt.Printf("Home Occupied By Worker %v \n", workerId)
	} else {
		fmt.Printf("Aisle %v Occupied By Worker %v \n", aisleId, workerId)
	}
}

func (w *warehouse) unoccupiedAisle(aisleId int, workerId int) {
	aisleToUpdate := w.aisleLocs[aisleId]
	aisleToUpdate.aisle.occupied = false
	if aisleId == int(w.workArea.location) {
		fmt.Printf("Worker %v Left Home  \n", workerId)
	} else {
		fmt.Printf("Aisle %v No Longer Occupied by Worker %v \n", aisleId, workerId)
	}

}
