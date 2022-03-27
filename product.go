package main

type sku string
type product struct {
	sku
}

type productLocation map[sku]int


func productLocations(init warehouse) productLocation {
	newProductLocations := make(productLocation)
	id := len(init.junctions)

	for _, junctionNode := range sortMap(init.junctions) {
		for area := 0; area < len(init.junctions[junctionNode].areas); area++ {
			for product := 0; product < len(init.junctions[junctionNode].areas[area].aisle.products); product++ {
				sku := init.junctions[junctionNode].areas[area].aisle.products[product].sku
				newProductLocations[sku] = id
			}
			id++
		}
	}
	return newProductLocations
}
