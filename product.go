package main

type product struct {
	sku string
}

type productLocation map[string]int

func (w *warehouse) productLocations(init warehouse) {
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
	w.productLocs = newProductLocations
}
