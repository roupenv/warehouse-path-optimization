package main

var warehouseConfig = warehouse{
	junctions: map[junctionNode]junctionData{
		0: {
			areas: []area{
				{
					aisle: aisle{
						products: []product{
							{sku: "A"},
							{sku: "B"},
							{sku: "C"},
						},
					},
				},
				{
					aisle: aisle{
						products: []product{
							{sku: "D"},
							{sku: "E"},
						},
					},
				},
			},
			nextJunction:           1,
			distanceToNextJunction: 5,
		},
		1: {
			areas: []area{
				{
					aisle: aisle{
						products: []product{
							{sku: "F"},
							{sku: "G"},
							{sku: "H"},
						},
					},
				},
				{
					aisle: aisle{
						products: []product{
							{sku: "I"},
							{sku: "J"},
						},
					},
				},
			},
			nextJunction:           2,
			distanceToNextJunction: 5,
		},
		2: {
			areas: []area{
				{
					aisle: aisle{
						products: []product{
							{sku: "K"},
							{sku: "L"},
						},
					},
				},
				{
					aisle: aisle{
						products: []product{
							{sku: "M"},
							{sku: "N"},
							{sku: "O"},
						},
					},
				},
			},
			nextJunction:           3,
			distanceToNextJunction: 20,
		},
		3: {
			areas: []area{
				{
					aisle: aisle{
						products: []product{
							{sku: "P"},
							{sku: "Q"},
							{sku: "R"},
						},
					},
				},
				{
					aisle: aisle{
						products: []product{
							{sku: "S"},
							{sku: "T"},
						},
					},
				},
				{
					workArea: workArea{
						numberItems: 0,
					},
				},
			},
			nextJunction:           4,
			distanceToNextJunction: 5,
		},
		4: {
			areas: []area{
				{
					aisle: aisle{
						products: []product{
							{sku: "U"},
							{sku: "V"},
						},
					},
				},
				{
					aisle: aisle{
						products: []product{
							{sku: "W"},
							{sku: "X"},
							{sku: "Y"},
						},
					},
				},
			},
			nextJunction:           1,
			distanceToNextJunction: 5,
		},
		5: {
			areas: []area{
				{
					aisle: aisle{
						products: []product{
							{sku: "Z"},
							{sku: "AA"},
							{sku: "AB"},
						},
					},
				},
				{
					aisle: aisle{
						products: []product{
							{sku: "AC"},
							{sku: "AD"},
						},
					},
				},
			},
		},
	},
	workers: []worker{
		{
			id: 0,
			dolly: dolly{
				capacity: 12,
			},
			status:          idle,
			currentLocation: 14,
		},
	},
	workAreaLocation: 14,
}
